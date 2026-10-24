package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopath"
	"gopath/internal/content"
	"gopath/internal/runner"
)

func main() {
	root, e := fs.Sub(gopath.FS(), "content")
	if e != nil {
		log.Fatal(e)
	}
	catalog, e := content.Load(root)
	if e != nil {
		log.Fatal(e)
	}
	if e = catalog.Validate(); e != nil {
		log.Fatal(e)
	}
	_, version, e := runner.Version()
	if e != nil {
		log.Fatal("Go is required to verify content: ", e)
	}
	cache := os.Getenv("GOCACHE")
	if cache == "" {
		cacheDir, cacheErr := os.UserCacheDir()
		if cacheErr != nil {
			log.Fatal(cacheErr)
		}
		cache = filepath.Join(cacheDir, "go-build")
	}
	if e := os.MkdirAll(cache, 0700); e != nil {
		log.Fatal(e)
	}
	run := runner.New(cache)
	total := 0
	verify := func(id string, starter, solution map[string]string, visible, hidden string, timeout time.Duration) {
		if len(starter) == 0 || len(solution) == 0 {
			log.Fatalf("%s: missing code", id)
		}
		if strings.Count(visible, "func Test") < 3 || strings.Count(hidden, "func Test") < 3 {
			log.Fatalf("%s: need at least 3 visible and hidden tests", id)
		}
		req := runner.Request{Files: starter, Version: version, Timeout: timeout}
		built, e := run.Execute(nilContext{}, id+"-build", req)
		if e != nil || built.ExitCode != 0 {
			log.Fatalf("%s: starter does not run: %v %s", id, e, built.Stderr)
		}
		req.Test = true
		req.VisibleTest = visible
		req.HiddenTest = hidden
		failed, e := run.Execute(nilContext{}, id+"-starter", req)
		if e != nil || failed.Passed {
			log.Fatalf("%s: starter must fail tests: %v", id, e)
		}
		req.Files = solution
		passed, e := run.Execute(nilContext{}, id+"-solution", req)
		if e != nil || !passed.Passed {
			log.Fatalf("%s: solution fails: %v %s %s", id, e, passed.Stdout, passed.Stderr)
		}
		total++
	}
	for _, ex := range catalog.Exercises {
		verify(ex.ID, ex.Starter, ex.Solution, ex.VisibleTest, ex.HiddenTest, 30*time.Second)
	}
	for _, m := range catalog.Milestones {
		verify(m.ID, m.Starter, m.Solution, m.VisibleTest, m.HiddenTest, 45*time.Second)
	}
	fmt.Printf("Verified %d exercises and milestones\n", total)
}

type nilContext struct{}

func (nilContext) Deadline() (time.Time, bool) { return time.Time{}, false }
func (nilContext) Done() <-chan struct{}       { return nil }
func (nilContext) Err() error                  { return nil }
func (nilContext) Value(any) any               { return nil }
