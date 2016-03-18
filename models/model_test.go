package models

import "testing"

// DATASOURCETEST is the db name and authentication for testing
const DATASOURCETEST = "user=maxwell dbname=careerCrawl_test sslmode=disable"

func TestStore_Open(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.Close()
}

func TestStore_CreateJob(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.CleanDatabase()

	if err := s.CreateJob(&Job{
		Name: "Hacker",
	}); err != nil {
		t.Fatal(err)
	}

	if job, err := s.Job("Hacker"); err != nil {
		t.Fatal(err)
	} else if job.Name != "Hacker" {
		t.Fatalf("unexpected name: %s", job.Name)
	}
}

func TestStore_CreateSkill(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()

	if err := s.CreateSkill(&Skill{
		Name: "ruby",
	}); err != nil {
		t.Fatal(err)
	}

	if skill, err := s.Skill("ruby"); err != nil {
		t.Fatal(err)
	} else if skill.Name != "ruby" {
		t.Fatalf("unexpected name: %s", skill.Name)
	}
}

func TestStore_CreateCategory_JobsBySkill(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()

	s.CreateIntegration(t)

	jobs, err := s.JobsBySkill("ruby")
	if err != nil {
		t.Fatal(err)
	}
	if len(jobs) != 2 {
		t.Fatalf("unexpected length: %s", len(jobs))
	}
	for i := 0; i < len(jobs); i++ {
		if jobs[i].Name != "Hacker" && jobs[i].Name != "QA" {
			t.Fatalf("unexpected name: %s", jobs[i].Name)
		}
	}
}

func TestStore_CreateCategory_SkillsByJob(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()

	s.CreateIntegration(t)

	skills, err := s.SkillsByJob("Coder")
	if err != nil {
		t.Fatal(err)
	}
	if len(skills) != 2 {
		t.Fatalf("unexpected length: %#v", skills)
	}
	for i := 0; i < len(skills); i++ {
		if skills[i].Name != "go" && skills[i].Name != "python" {
			t.Fatalf("unexpected skill: %s", skills[i].Name)
		}
	}
}

func TestStore_Jobs(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()
	actualJobs, _ := s.CreateIntegration(t)

	expectJobs, err := s.Jobs()
	if err != nil {
		t.Fatal(err)
	}

	compare := func(j Job, jobs []*Job) bool {
		for _, v := range jobs {
			if j.Name == v.Name {
				return false
			}
		}
		return true
	}
	for _, v := range expectJobs {
		if compare(v, actualJobs) {
			t.Fatalf("unexpected jobs: %s", v.Name)
		}
	}
}

func TestStore_Skills(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()
	_, actualSkills := s.CreateIntegration(t)

	expectSkills, err := s.Skills()
	if err != nil {
		t.Fatal(err)
	}

	compare := func(s Skill, skills []*Skill) bool {
		for _, v := range skills {
			if s.Name == v.Name {
				return false
			}
		}
		return true
	}
	for _, v := range expectSkills {
		if compare(v, actualSkills) {
			t.Fatalf("unexpected skills: %s", v.Name)
		}
	}
}

type TestStore struct {
	Store
}

func NewTestStore() *TestStore {
	return &TestStore{}
}

func (t *TestStore) MustOpen() {
	if err := t.Open(DATASOURCETEST); err != nil {
		panic(err)
	}
	if err := t.DB.Ping(); err != nil {
		panic(err)
	}
}

func (t *TestStore) CreateIntegration(test *testing.T) ([]*Job, []*Skill) {
	if err := t.CreateSkill(&Skill{
		Name: "ruby",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateSkill(&Skill{
		Name: "go",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateSkill(&Skill{
		Name: "python",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateJob(&Job{
		Name: "Hacker",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateJob(&Job{
		Name: "Coder",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateJob(&Job{
		Name: "QA",
	}); err != nil {
		test.Fatal(err)
	}

	hacker, err := t.Job("Hacker")
	if err != nil {
		test.Fatal(err)
	}
	qa, err := t.Job("QA")
	if err != nil {
		test.Fatal(err)
	}
	coder, err := t.Job("Coder")
	if err != nil {
		test.Fatal(err)
	}

	ruby, err := t.Skill("ruby")
	if err != nil {
		test.Fatal(err)
	}
	golang, err := t.Skill("go")
	if err != nil {
		test.Fatal(err)
	}
	python, err := t.Skill("python")
	if err != nil {
		test.Fatal(err)
	}

	if err := t.CreateCategory(&Category{
		JobID:   hacker.ID,
		SkillID: ruby.ID,
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateCategory(&Category{
		JobID:   qa.ID,
		SkillID: ruby.ID,
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateCategory(&Category{
		JobID:   coder.ID,
		SkillID: golang.ID,
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateCategory(&Category{
		JobID:   coder.ID,
		SkillID: python.ID,
	}); err != nil {
		test.Fatal(err)
	}
	return []*Job{hacker, qa, coder}, []*Skill{ruby, golang, python}
}

func (t *TestStore) CleanDatabase() {
	table := []string{"categories", "jobs", "skills"}
	for i := 0; i < len(table); i++ {
		sess := t.DB.NewSession(nil)
		sess.DeleteFrom(table[i]).Exec()
	}
}
