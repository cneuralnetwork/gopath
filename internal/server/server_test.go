package server

import (
	"encoding/json"
	"io/fs"
	"net/http/httptest"
	"testing"

	"gopath"
	"gopath/internal/content"
	"gopath/internal/runner"
	"gopath/internal/store"
)

func TestHealthAndQuiz(t *testing.T) {
	root, e := fs.Sub(gopath.FS(), "content")
	if e != nil {
		t.Fatal(e)
	}
	cat, e := content.Load(root)
	if e != nil {
		t.Fatal(e)
	}
	st, e := store.New(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	s := &Server{Catalog: cat, Store: st, Runner: runner.New(t.TempDir()), Version: "test"}
	h := s.Handler()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/api/health", nil))
	if rec.Code != 200 {
		t.Fatalf("health status %d", rec.Code)
	}
	var health map[string]any
	if e = json.Unmarshal(rec.Body.Bytes(), &health); e != nil {
		t.Fatal(e)
	}
	if health["version"] != "test" {
		t.Fatalf("unexpected health: %v", health)
	}
}
