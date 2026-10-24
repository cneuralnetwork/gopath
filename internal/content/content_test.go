package content

import (
	"io/fs"
	"testing"

	"gopath"
)

func TestLoadAndValidate(t *testing.T) {
	root, e := fs.Sub(gopath.FS(), "content")
	if e != nil {
		t.Fatal(e)
	}
	c, e := Load(root)
	if e != nil {
		t.Fatal(e)
	}
	if e = c.Validate(); e != nil {
		t.Fatal(e)
	}
	if got := len(c.Exercises); got != 68 {
		t.Fatalf("exercises: got %d, want 68", got)
	}
	if got := len(c.Milestones); got != 21 {
		t.Fatalf("milestones: got %d, want 21", got)
	}
}
