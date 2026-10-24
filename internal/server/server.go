package server

import (
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gopath/internal/content"
	"gopath/internal/runner"
	"gopath/internal/store"
)

type Server struct {
	Catalog         *content.Catalog
	Store           *store.Store
	Runner          *runner.Runner
	Web             fs.FS
	Version         string
	GoVersion       string
	GoModuleVersion string
}

func jsonReply(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
func fail(w http.ResponseWriter, code int, msg string) {
	jsonReply(w, code, map[string]string{"error": msg})
}
func decode(r *http.Request, v any) error {
	return json.NewDecoder(io.LimitReader(r.Body, 2<<20)).Decode(v)
}
func (s *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		jsonReply(w, 200, map[string]any{"version": s.Version, "goVersion": s.GoVersion, "goAvailable": s.GoVersion != ""})
	})
	m.HandleFunc("GET /api/course", s.course)
	m.HandleFunc("GET /api/search", s.search)
	m.HandleFunc("GET /api/concepts/{id}", s.concept)
	m.HandleFunc("POST /api/quiz/{conceptId}/{questionId}/check", s.checkQuiz)
	m.HandleFunc("GET /api/exercises/{id}", s.getExercise)
	m.HandleFunc("PUT /api/exercises/{id}/files", s.saveFiles)
	m.HandleFunc("POST /api/exercises/{id}/run", s.execute(false))
	m.HandleFunc("POST /api/exercises/{id}/test", s.execute(true))
	m.HandleFunc("POST /api/exercises/{id}/stop", s.stop)
	m.HandleFunc("POST /api/exercises/{id}/hint", s.hint)
	m.HandleFunc("POST /api/exercises/{id}/solution", s.solution)
	m.HandleFunc("POST /api/exercises/{id}/open", s.openEditor)
	m.HandleFunc("POST /api/exercises/{id}/sync", s.syncDisk)
	m.HandleFunc("GET /api/progress/export", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", "attachment; filename=gopath-progress.json")
		jsonReply(w, 200, s.Store.Snapshot())
	})
	m.HandleFunc("POST /api/progress/import", func(w http.ResponseWriter, r *http.Request) {
		b, e := io.ReadAll(io.LimitReader(r.Body, 10<<20))
		if e != nil || s.Store.Import(b) != nil {
			fail(w, 400, "invalid progress JSON")
			return
		}
		jsonReply(w, 200, map[string]bool{"ok": true})
	})
	m.HandleFunc("PUT /api/settings", func(w http.ResponseWriter, r *http.Request) {
		var v struct {
			UnlockAll bool `json:"unlockAll"`
		}
		if decode(r, &v) != nil {
			fail(w, 400, "invalid JSON")
			return
		}
		_ = s.Store.Update(func(st *store.State) { st.UnlockAll = v.UnlockAll })
		jsonReply(w, 200, v)
	})
	m.HandleFunc("GET /api/projects/{id}", s.project)
	m.HandleFunc("GET /api/milestones/{id}", s.getExercise)
	m.HandleFunc("PUT /api/milestones/{id}/files", s.saveFiles)
	m.HandleFunc("POST /api/milestones/{id}/run", s.execute(false))
	m.HandleFunc("POST /api/milestones/{id}/test", s.execute(true))
	m.HandleFunc("POST /api/milestones/{id}/stop", s.stop)
	m.HandleFunc("POST /api/milestones/{id}/hint", s.hint)
	m.HandleFunc("POST /api/milestones/{id}/solution", s.solution)
	m.HandleFunc("POST /api/milestones/{id}/open", s.openEditor)
	m.HandleFunc("POST /api/milestones/{id}/sync", s.syncDisk)
	if s.Web != nil {
		m.Handle("/", http.FileServer(http.FS(s.Web)))
	} else {
		m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			_, _ = io.WriteString(w, "<h1>Gopath frontend is not built</h1><p>Run <code>cd web && npm install && npm run build</code>.</p>")
		})
	}
	return m
}
func (s *Server) unlocked(id string, st store.State) bool {
	if st.UnlockAll {
		return true
	}
	for i, v := range s.Catalog.Order {
		if v == id {
			return i == 0 || s.conceptDone(s.Catalog.Order[i-1], st)
		}
	}
	return true
}
func (s *Server) conceptDone(id string, st store.State) bool {
	co := s.Catalog.Concepts[id]
	if co == nil {
		return false
	}
	for _, q := range co.Quiz {
		if !st.Quiz[id][q.ID] {
			return false
		}
	}
	for _, ex := range co.Exercises {
		if !st.Passed[ex.ID] {
			return false
		}
	}
	return true
}
func (s *Server) projectUnlocked(id string, st store.State) bool {
	if st.UnlockAll {
		return true
	}
	p := s.Catalog.Projects[id]
	if p == nil {
		return false
	}
	for _, ch := range s.Catalog.Course.Chapters {
		if ch.ID == p.After {
			for _, co := range ch.Concepts {
				if !s.conceptDone(co.ID, st) {
					return false
				}
			}
			return true
		}
	}
	return false
}
func (s *Server) course(w http.ResponseWriter, r *http.Request) {
	st := s.Store.Snapshot()
	type item struct {
		content.ConceptRef
		State string `json:"state"`
	}
	type chapter struct {
		ID       string `json:"id"`
		Title    string `json:"title"`
		Concepts []item `json:"concepts"`
		Done     int    `json:"done"`
	}
	out := []chapter{}
	for _, ch := range s.Catalog.Course.Chapters {
		c := chapter{ID: ch.ID, Title: ch.Title, Concepts: []item{}}
		for _, ref := range ch.Concepts {
			state := "locked"
			if s.conceptDone(ref.ID, st) {
				state = "done"
				c.Done++
			} else if s.unlocked(ref.ID, st) {
				state = "not-started"
				if ref.ID == st.LastConcept || st.LastConcept == "" && len(out) == 0 && len(c.Concepts) == 0 {
					state = "current"
				}
			}
			c.Concepts = append(c.Concepts, item{ref, state})
		}
		out = append(out, c)
	}
	projectRefs := []map[string]any{}
	for _, p := range s.Catalog.Course.Projects {
		state := "locked"
		if s.projectUnlocked(p.ID, st) {
			state = "current"
			all := true
			for _, m := range s.Catalog.Projects[p.ID].Milestones {
				if !st.Passed[m.ID] {
					all = false
				}
			}
			if all {
				state = "done"
			}
		}
		projectRefs = append(projectRefs, map[string]any{"id": p.ID, "title": p.Title, "after": p.After, "state": state})
	}
	jsonReply(w, 200, map[string]any{"chapters": out, "projects": projectRefs, "progress": st})
}
func (s *Server) concept(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	co := s.Catalog.Concepts[id]
	if co == nil {
		fail(w, 404, "concept not found")
		return
	}
	st := s.Store.Snapshot()
	if !s.unlocked(id, st) {
		fail(w, 403, "concept locked")
		return
	}
	copyCo := *co
	copyCo.Quiz = append([]content.Question(nil), co.Quiz...)
	for i := range copyCo.Quiz {
		copyCo.Quiz[i].Answer = ""
		copyCo.Quiz[i].Explanation = ""
	}
	copyCo.Exercises = nil
	jsonReply(w, 200, map[string]any{"concept": copyCo, "exercises": exerciseList(co.Exercises, st), "progress": st})
	_ = s.Store.Update(func(st *store.State) { st.LastConcept = id })
}
func exerciseList(xs []*content.Exercise, st store.State) []map[string]any {
	out := []map[string]any{}
	for _, x := range xs {
		out = append(out, map[string]any{"id": x.ID, "title": x.Title, "difficulty": x.Difficulty, "passed": st.Passed[x.ID]})
	}
	return out
}
func (s *Server) checkQuiz(w http.ResponseWriter, r *http.Request) {
	id, qid := r.PathValue("conceptId"), r.PathValue("questionId")
	co := s.Catalog.Concepts[id]
	if co == nil {
		fail(w, 404, "concept not found")
		return
	}
	var body struct {
		Answer string `json:"answer"`
	}
	if decode(r, &body) != nil {
		fail(w, 400, "invalid JSON")
		return
	}
	for _, q := range co.Quiz {
		if q.ID == qid {
			valid := false
			for _, o := range q.Options {
				if o.ID == body.Answer {
					valid = true
				}
			}
			if !valid {
				fail(w, 400, "unknown option")
				return
			}
			correct := body.Answer == q.Answer
			_ = s.Store.Update(func(st *store.State) {
				if st.Quiz[id] == nil {
					st.Quiz[id] = map[string]bool{}
				}
				if correct {
					st.Quiz[id][qid] = true
				} else {
					st.Wrong[id+"/"+qid] = time.Now().Format("2006-01-02")
				}
			})
			jsonReply(w, 200, map[string]any{"correct": correct, "explanation": q.Explanation, "passed": s.Store.Snapshot().Quiz[id][qid]})
			return
		}
	}
	fail(w, 404, "question not found")
}
func (s *Server) find(id string) (content.ExerciseMeta, string, map[string]string, map[string]string, string, string, bool) {
	if e := s.Catalog.Exercises[id]; e != nil {
		return e.ExerciseMeta, e.README, e.Starter, e.Solution, e.VisibleTest, e.HiddenTest, true
	}
	if m := s.Catalog.Milestones[id]; m != nil && s.projectUnlocked(m.ProjectID, s.Store.Snapshot()) {
		return m.ExerciseMeta, m.README, m.Starter, m.Solution, m.VisibleTest, m.HiddenTest, true
	}
	return content.ExerciseMeta{}, "", nil, nil, "", "", false
}
func (s *Server) getExercise(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	meta, readme, starter, _, visible, _, ok := s.find(id)
	if !ok {
		fail(w, 404, "exercise not found")
		return
	}
	st := s.Store.Snapshot()
	files := st.Files[id]
	if len(files) == 0 {
		files = starter
	}
	jsonReply(w, 200, map[string]any{"meta": meta, "readme": readme, "files": files, "starter": starter, "visibleTests": visible, "hintsShown": st.Hints[id], "attempts": st.Attempts[id], "passed": st.Passed[id], "solutionViewed": st.Solutions[id]})
}
func (s *Server) saveFiles(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, _, _, _, _, _, ok := s.find(id)
	if !ok {
		fail(w, 404, "exercise not found")
		return
	}
	var v struct {
		Files map[string]string `json:"files"`
	}
	if decode(r, &v) != nil || len(v.Files) == 0 {
		fail(w, 400, "invalid files")
		return
	}
	for name, src := range v.Files {
		if strings.Contains(name, "..") || filepath.IsAbs(name) || !strings.HasSuffix(name, ".go") || len(src) > 256<<10 {
			fail(w, 400, "invalid file")
			return
		}
	}
	_ = s.Store.Update(func(st *store.State) { st.Files[id] = v.Files })
	jsonReply(w, 200, map[string]bool{"ok": true})
}
func (s *Server) execute(test bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		meta, _, starter, _, visible, hidden, ok := s.find(id)
		if !ok {
			fail(w, 404, "exercise not found")
			return
		}
		if s.GoVersion == "" {
			fail(w, 503, "Go is not installed. Install Go from go.dev/dl and restart Gopath.")
			return
		}
		st := s.Store.Snapshot()
		files := st.Files[id]
		if len(files) == 0 {
			files = starter
		}
		timeout := 10 * time.Second
		if _, project := s.Catalog.Milestones[id]; project {
			timeout = 30 * time.Second
		}
		if meta.Timeout > 0 {
			timeout = time.Duration(meta.Timeout) * time.Second
		}
		result, e := s.Runner.Execute(r.Context(), id, runner.Request{Files: files, VisibleTest: visible, HiddenTest: hidden, Version: s.GoModuleVersion, Timeout: timeout, Test: test})
		if e != nil {
			if strings.Contains(e.Error(), "already running") {
				fail(w, 409, e.Error())
			} else {
				fail(w, 500, e.Error())
			}
			return
		}
		if test {
			_ = s.Store.Update(func(st *store.State) { st.Attempts[id]++ })
			if result.Passed {
				_ = s.Store.RecordPass(id)
			}
		}
		jsonReply(w, 200, result)
	}
}
func (s *Server) stop(w http.ResponseWriter, r *http.Request) {
	jsonReply(w, 200, map[string]bool{"stopped": s.Runner.Stop(r.PathValue("id"))})
}
func (s *Server) hint(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	meta, _, _, _, _, _, ok := s.find(id)
	if !ok {
		fail(w, 404, "exercise not found")
		return
	}
	st := s.Store.Snapshot()
	n := st.Hints[id]
	if n < len(meta.Hints) {
		n++
		_ = s.Store.Update(func(st *store.State) { st.Hints[id] = n })
	}
	jsonReply(w, 200, map[string]any{"hint": func() string {
		if n == 0 {
			return ""
		}
		return meta.Hints[n-1]
	}(), "shown": n, "total": len(meta.Hints)})
}
func (s *Server) solution(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, _, _, solution, _, _, ok := s.find(id)
	if !ok {
		fail(w, 404, "exercise not found")
		return
	}
	_ = s.Store.Update(func(st *store.State) { st.Solutions[id] = true })
	jsonReply(w, 200, map[string]any{"files": solution})
}
func (s *Server) project(w http.ResponseWriter, r *http.Request) {
	p := s.Catalog.Projects[r.PathValue("id")]
	if p == nil {
		fail(w, 404, "project not found")
		return
	}
	st := s.Store.Snapshot()
	if !s.projectUnlocked(p.ID, st) {
		fail(w, 403, "project locked")
		return
	}
	out := []map[string]any{}
	for _, m := range p.Milestones {
		out = append(out, map[string]any{"id": m.ID, "title": m.Title, "passed": st.Passed[m.ID]})
	}
	jsonReply(w, 200, map[string]any{"project": p.ProjectRef, "readme": p.README, "milestones": out})
}
func (s *Server) diskPath(id string) (string, error) {
	home, e := os.UserHomeDir()
	if e != nil {
		return "", e
	}
	if ex := s.Catalog.Exercises[id]; ex != nil {
		return filepath.Join(home, "gopath-workspace", ex.ConceptID, id), nil
	}
	if m := s.Catalog.Milestones[id]; m != nil {
		return filepath.Join(home, "gopath-workspace", m.ProjectID, id), nil
	}
	return "", errors.New("not found")
}
func (s *Server) openEditor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, _, starter, _, visible, _, ok := s.find(id)
	if !ok {
		fail(w, 404, "exercise not found")
		return
	}
	p, e := s.diskPath(id)
	if e != nil {
		fail(w, 500, e.Error())
		return
	}
	if e = os.MkdirAll(p, 0700); e != nil {
		fail(w, 500, e.Error())
		return
	}
	files := s.Store.Snapshot().Files[id]
	if len(files) == 0 {
		files = starter
	}
	for name, src := range files {
		dest := filepath.Join(p, name)
		if e = os.MkdirAll(filepath.Dir(dest), 0700); e != nil {
			fail(w, 500, e.Error())
			return
		}
		if e = os.WriteFile(dest, []byte(src), 0600); e != nil {
			fail(w, 500, e.Error())
			return
		}
	}
	_ = os.WriteFile(filepath.Join(p, "visible_test.go"), []byte(visible), 0600)
	_ = os.WriteFile(filepath.Join(p, "go.mod"), []byte("module exercise\n\ngo "+s.GoModuleVersion+"\n"), 0600)
	jsonReply(w, 200, map[string]string{"path": p})
}
func (s *Server) syncDisk(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, _, starter, _, _, _, ok := s.find(id)
	if !ok {
		fail(w, 404, "exercise not found")
		return
	}
	p, e := s.diskPath(id)
	if e != nil {
		fail(w, 500, e.Error())
		return
	}
	files := map[string]string{}
	for name := range starter {
		b, e := os.ReadFile(filepath.Join(p, name))
		if e != nil {
			fail(w, 400, "file missing: "+name)
			return
		}
		files[name] = string(b)
	}
	_ = s.Store.Update(func(st *store.State) { st.Files[id] = files })
	jsonReply(w, 200, map[string]any{"files": files})
}
func atoi(v string) int { n, _ := strconv.Atoi(v); return n }

func (s *Server) search(w http.ResponseWriter, r *http.Request) {
	q := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))
	out := []map[string]string{}
	if q == "" {
		jsonReply(w, 200, out)
		return
	}
	for _, ch := range s.Catalog.Course.Chapters {
		for _, ref := range ch.Concepts {
			co := s.Catalog.Concepts[ref.ID]
			if strings.Contains(strings.ToLower(ref.Title+" "+strings.Join(ref.Keywords, " ")), q) {
				out = append(out, map[string]string{"id": ref.ID, "title": ref.Title, "kind": "concept", "conceptId": ref.ID})
			}
			for _, card := range co.Cards {
				if strings.Contains(strings.ToLower(card.Title+" "+card.Body), q) {
					out = append(out, map[string]string{"id": ref.ID + "/card/" + card.Title, "title": card.Title, "kind": "card", "conceptId": ref.ID})
				}
			}
			for _, ex := range co.Exercises {
				if strings.Contains(strings.ToLower(ex.Title+" "+ex.README), q) {
					out = append(out, map[string]string{"id": ex.ID, "title": ex.Title, "kind": "exercise", "conceptId": ref.ID})
				}
			}
		}
	}
	if len(out) > 30 {
		out = out[:30]
	}
	jsonReply(w, 200, out)
}
