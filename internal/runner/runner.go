// Package runner isolates each learner invocation in a temporary Go module.
package runner

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

type TestResult struct {
	Name     string  `json:"name"`
	Status   string  `json:"status"`
	Duration float64 `json:"duration"`
	Output   string  `json:"output"`
	Hidden   bool    `json:"hidden"`
}
type Result struct {
	Stdout     string       `json:"stdout"`
	Stderr     string       `json:"stderr"`
	ExitCode   int          `json:"exitCode"`
	DurationMs int64        `json:"durationMs"`
	TimedOut   bool         `json:"timedOut"`
	Tests      []TestResult `json:"tests,omitempty"`
	Passed     bool         `json:"passed"`
}
type Request struct {
	Files       map[string]string
	VisibleTest string
	HiddenTest  string
	Version     string
	Timeout     time.Duration
	Test        bool
}
type Runner struct {
	mu       sync.Mutex
	active   map[string]context.CancelFunc
	CacheDir string
}

func New(cacheDir string) *Runner {
	return &Runner{active: map[string]context.CancelFunc{}, CacheDir: cacheDir}
}
func (r *Runner) Stop(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	cancel, ok := r.active[id]
	if ok {
		cancel()
	}
	return ok
}

var safeName = regexp.MustCompile(`^[a-zA-Z0-9_./-]+\.go$`)

func safePath(name string) bool {
	return safeName.MatchString(name) && !strings.HasPrefix(name, "/") && !strings.Contains(name, "..")
}

type limitedWriter struct {
	b bytes.Buffer
	n int
}

func (w *limitedWriter) Write(p []byte) (int, error) {
	count := len(p)
	if w.n < 65536 {
		take := 65536 - w.n
		if take > count {
			take = count
		}
		w.b.Write(p[:take])
		w.n += take
	}
	return count, nil
}
func (r *Runner) Execute(parent context.Context, id string, req Request) (Result, error) {
	r.mu.Lock()
	if _, ok := r.active[id]; ok {
		r.mu.Unlock()
		return Result{}, errors.New("this exercise is already running")
	}
	timeout := req.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}
	ctx, cancel := context.WithTimeout(parent, timeout)
	r.active[id] = cancel
	r.mu.Unlock()
	defer func() { cancel(); r.mu.Lock(); delete(r.active, id); r.mu.Unlock() }()
	if len(req.Files) == 0 {
		return Result{}, errors.New("no source files")
	}
	base := filepath.Join(os.TempDir(), "gopath-runs")
	if e := os.MkdirAll(base, 0700); e != nil {
		return Result{}, e
	}
	dir, e := os.MkdirTemp(base, "run-")
	if e != nil {
		return Result{}, e
	}
	defer os.RemoveAll(dir)
	for name, code := range req.Files {
		if !safePath(name) {
			return Result{}, fmt.Errorf("invalid source path %q", name)
		}
		dest := filepath.Join(dir, filepath.FromSlash(name))
		if e = os.MkdirAll(filepath.Dir(dest), 0700); e != nil {
			return Result{}, e
		}
		if e = os.WriteFile(dest, []byte(code), 0600); e != nil {
			return Result{}, e
		}
	}
	version := req.Version
	if version == "" {
		version = "1.22"
	}
	if e = os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module exercise\n\ngo "+version+"\n"), 0600); e != nil {
		return Result{}, e
	}
	if req.Test {
		if e = os.WriteFile(filepath.Join(dir, "visible_test.go"), []byte(req.VisibleTest), 0600); e != nil {
			return Result{}, e
		}
		if e = os.WriteFile(filepath.Join(dir, "hidden_test.go"), []byte(req.HiddenTest), 0600); e != nil {
			return Result{}, e
		}
	}
	args := []string{"run", "."}
	if req.Test {
		args = []string{"test", "-json", "-count=1", "./..."}
	}
	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOFLAGS=-mod=mod", "GOTOOLCHAIN=local", "CGO_ENABLED=0", "GOPROXY=off", "GOSUMDB=off", "GOCACHE="+r.CacheDir)
	if runtime.GOOS != "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}
	var out, errout limitedWriter
	cmd.Stdout = &out
	cmd.Stderr = &errout
	start := time.Now()
	e = cmd.Start()
	if e != nil {
		return Result{}, e
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case e = <-done:
	case <-ctx.Done():
		if cmd.Process != nil {
			if runtime.GOOS != "windows" {
				_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			} else {
				_ = cmd.Process.Kill()
			}
		}
		e = <-done
	}
	if ctx.Err() != nil && cmd.Process != nil && runtime.GOOS != "windows" {
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
	res := Result{Stdout: out.b.String(), Stderr: errout.b.String(), DurationMs: time.Since(start).Milliseconds(), TimedOut: ctx.Err() == context.DeadlineExceeded}
	if e != nil {
		var exit *exec.ExitError
		if errors.As(e, &exit) {
			res.ExitCode = exit.ExitCode()
		} else {
			res.ExitCode = -1
		}
	}
	if req.Test {
		res.Tests = parseTests(res.Stdout)
		res.Stdout = plainOutput(res.Stdout)
		res.Passed = res.ExitCode == 0 && !res.TimedOut
	} else {
		res.Passed = res.ExitCode == 0 && !res.TimedOut
	}
	return res, nil
}

type testEvent struct {
	Action  string  `json:"Action"`
	Test    string  `json:"Test"`
	Output  string  `json:"Output"`
	Elapsed float64 `json:"Elapsed"`
}

func parseTests(data string) []TestResult {
	state := map[string]*TestResult{}
	order := []string{}
	scan := bufio.NewScanner(strings.NewReader(data))
	scan.Buffer(make([]byte, 4096), 1024*1024)
	for scan.Scan() {
		var ev testEvent
		if json.Unmarshal(scan.Bytes(), &ev) != nil || ev.Test == "" {
			continue
		}
		t, ok := state[ev.Test]
		if !ok {
			t = &TestResult{Name: ev.Test, Hidden: strings.HasPrefix(ev.Test, "TestHidden")}
			state[ev.Test] = t
			order = append(order, ev.Test)
		}
		if ev.Action == "output" {
			if t.Hidden {
				if strings.Contains(ev.Output, "hint:") {
					t.Output += ev.Output
				}
			} else {
				t.Output += ev.Output
			}
		}
		if ev.Action == "pass" || ev.Action == "fail" || ev.Action == "skip" {
			t.Status = ev.Action
			t.Duration = ev.Elapsed
		}
	}
	results := make([]TestResult, 0, len(order))
	for _, name := range order {
		results = append(results, *state[name])
	}
	return results
}
func Version() (string, string, error) {
	p, e := exec.LookPath("go")
	if e != nil {
		return "", "", e
	}
	b, e := exec.Command(p, "version").Output()
	if e != nil {
		return "", "", e
	}
	re := regexp.MustCompile(`go([0-9]+\.[0-9]+)`)
	m := re.FindStringSubmatch(string(b))
	if len(m) < 2 {
		return strings.TrimSpace(string(b)), "1.22", nil
	}
	return strings.TrimSpace(string(b)), m[1], nil
}
func CopyTo(w io.Writer, r Result) {
	_, _ = io.WriteString(w, r.Stdout+r.Stderr+strconv.Itoa(r.ExitCode))
}

// plainOutput keeps compiler output while hiding the JSON event stream.
func plainOutput(data string) string {
	var b strings.Builder
	scan := bufio.NewScanner(strings.NewReader(data))
	scan.Buffer(make([]byte, 4096), 1024*1024)
	for scan.Scan() {
		var ev testEvent
		if json.Unmarshal(scan.Bytes(), &ev) == nil && ev.Action != "" {
			continue
		}
		b.WriteString(scan.Text())
		b.WriteByte('\n')
	}
	return b.String()
}
