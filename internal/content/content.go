package content

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"
)

type Course struct {
	Chapters []Chapter    `json:"chapters"`
	Projects []ProjectRef `json:"projects"`
}
type Chapter struct {
	ID       string       `json:"id"`
	Title    string       `json:"title"`
	Concepts []ConceptRef `json:"concepts"`
}
type ConceptRef struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Dir      string   `json:"dir"`
	Keywords []string `json:"keywords"`
}
type ProjectRef struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	After string `json:"after"`
	Dir   string `json:"dir"`
}
type Question struct {
	ID          string   `json:"id"`
	Question    string   `json:"question"`
	Code        string   `json:"code,omitempty"`
	Options     []Option `json:"options"`
	Answer      string   `json:"answer,omitempty"`
	Explanation string   `json:"explanation,omitempty"`
}
type Option struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Code bool   `json:"code,omitempty"`
}
type Card struct {
	Title   string `json:"title"`
	Kind    string `json:"kind"`
	Body    string `json:"body"`
	Example string `json:"example"`
}
type ExerciseMeta struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Difficulty string   `json:"difficulty"`
	Hints      []string `json:"hints"`
	Timeout    int      `json:"timeout,omitempty"`
}
type Exercise struct {
	ExerciseMeta
	Dir         string            `json:"-"`
	ConceptID   string            `json:"conceptId"`
	README      string            `json:"readme"`
	Starter     map[string]string `json:"starter"`
	Solution    map[string]string `json:"-"`
	VisibleTest string            `json:"visibleTests"`
	HiddenTest  string            `json:"-"`
}
type Concept struct {
	ConceptRef
	ChapterID string      `json:"chapterId"`
	Lesson    string      `json:"lesson"`
	Quiz      []Question  `json:"quiz"`
	Cards     []Card      `json:"cards"`
	Exercises []*Exercise `json:"exercises"`
}
type Milestone struct {
	ExerciseMeta
	Dir         string            `json:"-"`
	README      string            `json:"readme"`
	Starter     map[string]string `json:"starter"`
	Solution    map[string]string `json:"-"`
	VisibleTest string            `json:"visibleTests"`
	HiddenTest  string            `json:"-"`
	ProjectID   string            `json:"projectId"`
}
type Project struct {
	ProjectRef
	README     string       `json:"readme"`
	Milestones []*Milestone `json:"milestones"`
}
type Catalog struct {
	Course     Course
	Concepts   map[string]*Concept
	Exercises  map[string]*Exercise
	Projects   map[string]*Project
	Milestones map[string]*Milestone
	Order      []string
}

func readJSON[T any](fsys fs.FS, name string, out *T) error {
	b, e := fs.ReadFile(fsys, name)
	if e != nil {
		return e
	}
	if e = json.Unmarshal(b, out); e != nil {
		return fmt.Errorf("%s: %w", name, e)
	}
	return nil
}
func readText(fsys fs.FS, name string) (string, error) {
	b, e := fs.ReadFile(fsys, name)
	return string(b), e
}
func goFiles(fsys fs.FS, dir string) (map[string]string, error) {
	out := map[string]string{}
	err := fs.WalkDir(fsys, dir, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if !d.IsDir() && strings.HasSuffix(p, ".go") {
			s, e := readText(fsys, p)
			if e != nil {
				return e
			}
			out[strings.TrimPrefix(p, dir+"/")] = s
		}
		return nil
	})
	return out, err
}
func Load(fsys fs.FS) (*Catalog, error) {
	c := &Catalog{Concepts: map[string]*Concept{}, Exercises: map[string]*Exercise{}, Projects: map[string]*Project{}, Milestones: map[string]*Milestone{}}
	if e := readJSON(fsys, "course.yaml", &c.Course); e != nil {
		return nil, e
	}
	for _, ch := range c.Course.Chapters {
		for _, ref := range ch.Concepts {
			if _, ok := c.Concepts[ref.ID]; ok {
				return nil, fmt.Errorf("duplicate concept %s", ref.ID)
			}
			co := &Concept{ConceptRef: ref, ChapterID: ch.ID}
			var e error
			co.Lesson, e = readText(fsys, path.Join(ref.Dir, "lesson.md"))
			if e != nil {
				return nil, e
			}
			if e = readJSON(fsys, path.Join(ref.Dir, "quiz.yaml"), &co.Quiz); e != nil {
				return nil, e
			}
			if e = readJSON(fsys, path.Join(ref.Dir, "cards.yaml"), &co.Cards); e != nil {
				return nil, e
			}
			entries, e := fs.ReadDir(fsys, path.Join(ref.Dir, "exercises"))
			if e != nil {
				return nil, e
			}
			for _, entry := range entries {
				if !entry.IsDir() {
					continue
				}
				dir := path.Join(ref.Dir, "exercises", entry.Name())
				ex := &Exercise{Dir: dir, ConceptID: ref.ID}
				if e = readJSON(fsys, path.Join(dir, "exercise.yaml"), &ex.ExerciseMeta); e != nil {
					return nil, e
				}
				if ex.ID == "" {
					return nil, fmt.Errorf("missing exercise id at %s", dir)
				}
				ex.README, e = readText(fsys, path.Join(dir, "README.md"))
				if e != nil {
					return nil, e
				}
				ex.Starter, e = goFiles(fsys, path.Join(dir, "starter"))
				if e != nil {
					return nil, e
				}
				ex.Solution, e = goFiles(fsys, path.Join(dir, "solution"))
				if e != nil {
					return nil, e
				}
				ex.VisibleTest, e = readText(fsys, path.Join(dir, "visible_test.go"))
				if e != nil {
					return nil, e
				}
				ex.HiddenTest, e = readText(fsys, path.Join(dir, "hidden_test.go"))
				if e != nil {
					return nil, e
				}
				co.Exercises = append(co.Exercises, ex)
				if _, ok := c.Exercises[ex.ID]; ok {
					return nil, fmt.Errorf("duplicate exercise %s", ex.ID)
				}
				c.Exercises[ex.ID] = ex
			}
			sort.Slice(co.Exercises, func(i, j int) bool { return co.Exercises[i].ID < co.Exercises[j].ID })
			c.Concepts[ref.ID] = co
			c.Order = append(c.Order, ref.ID)
		}
	}
	for _, ref := range c.Course.Projects {
		p := &Project{ProjectRef: ref}
		var e error
		p.README, e = readText(fsys, path.Join(ref.Dir, "README.md"))
		if e != nil {
			return nil, e
		}
		entries, e := fs.ReadDir(fsys, path.Join(ref.Dir, "milestones"))
		if e != nil {
			return nil, e
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			dir := path.Join(ref.Dir, "milestones", entry.Name())
			m := &Milestone{Dir: dir, ProjectID: ref.ID}
			if e = readJSON(fsys, path.Join(dir, "milestone.yaml"), &m.ExerciseMeta); e != nil {
				return nil, e
			}
			m.README, e = readText(fsys, path.Join(dir, "README.md"))
			if e != nil {
				return nil, e
			}
			m.Starter, e = goFiles(fsys, path.Join(dir, "starter"))
			if e != nil {
				return nil, e
			}
			m.Solution, e = goFiles(fsys, path.Join(dir, "solution"))
			if e != nil {
				return nil, e
			}
			m.VisibleTest, e = readText(fsys, path.Join(dir, "visible_test.go"))
			if e != nil {
				return nil, e
			}
			m.HiddenTest, e = readText(fsys, path.Join(dir, "hidden_test.go"))
			if e != nil {
				return nil, e
			}
			p.Milestones = append(p.Milestones, m)
			c.Milestones[m.ID] = m
		}
		sort.Slice(p.Milestones, func(i, j int) bool { return p.Milestones[i].ID < p.Milestones[j].ID })
		c.Projects[p.ID] = p
	}
	return c, nil
}
func (c *Catalog) Validate() error {
	if len(c.Concepts) != 34 {
		return fmt.Errorf("expected 34 concepts, got %d", len(c.Concepts))
	}
	for id, co := range c.Concepts {
		if len(co.Quiz) < 4 || len(co.Quiz) > 8 {
			return fmt.Errorf("%s: need 4–8 questions", id)
		}
		if len(co.Exercises) < 2 || len(co.Exercises) > 4 {
			return fmt.Errorf("%s: need 2–4 exercises", id)
		}
		if len(co.Cards) < 1 || len(co.Cards) > 4 {
			return fmt.Errorf("%s: need 1–4 cards", id)
		}
		for _, q := range co.Quiz {
			found := false
			for _, o := range q.Options {
				if o.ID == q.Answer {
					found = true
				}
			}
			if !found {
				return fmt.Errorf("%s/%s: invalid answer", id, q.ID)
			}
		}
	}
	if len(c.Projects) != 7 {
		return fmt.Errorf("expected 7 projects, got %d", len(c.Projects))
	}
	return nil
}
