package store

import "testing"

func TestPersistsAndImports(t *testing.T) {
	dir := t.TempDir()
	s, e := New(dir)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.RecordPass("exercise"); e != nil {
		t.Fatal(e)
	}
	again, e := New(dir)
	if e != nil {
		t.Fatal(e)
	}
	if !again.Snapshot().Passed["exercise"] {
		t.Fatal("pass not persisted")
	}
	if e = again.Import([]byte(`{"passed":{"other":true}}`)); e != nil {
		t.Fatal(e)
	}
	if !again.Snapshot().Passed["other"] {
		t.Fatal("import failed")
	}
}
