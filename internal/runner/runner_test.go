package runner

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestParseTests(t *testing.T) {
	input := `{"Action":"run","Test":"TestA"}
{"Action":"output","Test":"TestA","Output":"hello\\n"}
{"Action":"pass","Test":"TestA","Elapsed":0.01}
{"Action":"run","Test":"TestHidden1"}
{"Action":"output","Test":"TestHidden1","Output":"hint: boundary\\n"}
{"Action":"fail","Test":"TestHidden1","Elapsed":0.02}`
	got := parseTests(input)
	if len(got) != 2 || got[0].Status != "pass" || got[1].Status != "fail" || !got[1].Hidden {
		t.Fatalf("unexpected results: %#v", got)
	}
	if strings.Contains(plainOutput(input), "Action") {
		t.Fatal("JSON events leaked")
	}
}
func TestTimeoutKillsLoop(t *testing.T) {
	if _, _, e := Version(); e != nil {
		t.Skip("Go toolchain missing")
	}
	r := New(t.TempDir())
	res, e := r.Execute(context.Background(), "loop", Request{Files: map[string]string{"main.go": "package main\nfunc main(){for {}}\n"}, Version: "1.22", Timeout: 10 * time.Second})
	if e != nil {
		t.Fatal(e)
	}
	if !res.TimedOut {
		t.Fatalf("wanted timeout, got %#v", res)
	}
}
