// Package store saves one learner's state in a small atomic JSON file.
package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type State struct {
	Files       map[string]map[string]string `json:"files"`
	Quiz        map[string]map[string]bool   `json:"quiz"`
	Wrong       map[string]string            `json:"wrong"`
	Passed      map[string]bool              `json:"passed"`
	Attempts    map[string]int               `json:"attempts"`
	Hints       map[string]int               `json:"hints"`
	Solutions   map[string]bool              `json:"solutions"`
	UnlockAll   bool                         `json:"unlockAll"`
	LastConcept string                       `json:"lastConcept"`
	Activity    map[string]bool              `json:"activity"`
}
type Store struct {
	mu    sync.Mutex
	path  string
	State State
}

func New(dir string) (*Store, error) {
	if e := os.MkdirAll(dir, 0700); e != nil {
		return nil, e
	}
	s := &Store{path: filepath.Join(dir, "progress.json")}
	s.State = State{Files: map[string]map[string]string{}, Quiz: map[string]map[string]bool{}, Wrong: map[string]string{}, Passed: map[string]bool{}, Attempts: map[string]int{}, Hints: map[string]int{}, Solutions: map[string]bool{}, Activity: map[string]bool{}}
	if b, e := os.ReadFile(s.path); e == nil {
		if e = json.Unmarshal(b, &s.State); e != nil {
			return nil, e
		}
	} else if !os.IsNotExist(e) {
		return nil, e
	}
	s.normalize()
	return s, nil
}
func (s *Store) normalize() {
	if s.State.Files == nil {
		s.State.Files = map[string]map[string]string{}
	}
	if s.State.Quiz == nil {
		s.State.Quiz = map[string]map[string]bool{}
	}
	if s.State.Wrong == nil {
		s.State.Wrong = map[string]string{}
	}
	if s.State.Passed == nil {
		s.State.Passed = map[string]bool{}
	}
	if s.State.Attempts == nil {
		s.State.Attempts = map[string]int{}
	}
	if s.State.Hints == nil {
		s.State.Hints = map[string]int{}
	}
	if s.State.Solutions == nil {
		s.State.Solutions = map[string]bool{}
	}
	if s.State.Activity == nil {
		s.State.Activity = map[string]bool{}
	}
}
func (s *Store) save() error {
	b, e := json.MarshalIndent(s.State, "", "  ")
	if e != nil {
		return e
	}
	tmp := s.path + ".tmp"
	if e = os.WriteFile(tmp, b, 0600); e != nil {
		return e
	}
	return os.Rename(tmp, s.path)
}
func (s *Store) Update(fn func(*State)) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	fn(&s.State)
	s.normalize()
	return s.save()
}
func (s *Store) Snapshot() State {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, _ := json.Marshal(s.State)
	var out State
	_ = json.Unmarshal(b, &out)
	return out
}
func (s *Store) Import(data []byte) error {
	var next State
	if e := json.Unmarshal(data, &next); e != nil {
		return e
	}
	return s.Update(func(st *State) { *st = next })
}
func (s *Store) RecordPass(id string) error {
	return s.Update(func(st *State) { st.Passed[id] = true; st.Activity[time.Now().Format("2006-01-02")] = true })
}
