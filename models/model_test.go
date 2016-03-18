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

func TestStore_CreateCategory(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()

	if err := s.CreateSkill(&Skill{
		Name: "ruby",
	}); err != nil {
		t.Fatal(err)
	} else if err := s.CreateJob(&Job{
		Name: "Hacker",
	}); err != nil {
		t.Fatal(err)
	}

	if j, err := s.Job("Hacker"); err != nil {
		t.Fatal(err)
	} else if skill, err := s.Skill("ruby"); err != nil {
		t.Fatal(err)
	} else if err := s.CreateCategory(&Category{
		JobID:   j.ID,
		SkillID: skill.ID,
	}); err != nil {
		t.Fatal(err)
	} else if jobs, err := s.JobsBySkill("ruby"); err != nil {
		t.Fatal(err)
	} else if jobs[0].Name != "Hacker" {
		t.Fatalf("unexpected name: %s", jobs[0].Name)
	}
}

func TestStore_CreateCategory_MultipleSkills(t *testing.T) {
	s := NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()

	if err := s.CreateSkill(&Skill{
		Name: "ruby",
	}); err != nil {
		t.Fatal(err)
	} else if err := s.CreateSkill(&Skill{
		Name: "go",
	}); err != nil {
		t.Fatal(err)
	} else if err := s.CreateJob(&Job{
		Name: "Hacker",
	}); err != nil {
		t.Fatal(err)
	} else if err := s.CreateJob(&Job{
		Name: "Coder",
	}); err != nil {
		t.Fatal(err)
	} else if err := s.CreateJob(&Job{
		Name: "QA",
	}); err != nil {
		t.Fatal(err)
	}

	if j, err := s.Job("Hacker"); err != nil {
		t.Fatal(err)
	} else if skill, err := s.Skill("ruby"); err != nil {
		t.Fatal(err)
	} else if err := s.CreateCategory(&Category{
		JobID:   j.ID,
		SkillID: skill.ID,
	}); err != nil {
		t.Fatal(err)
	} else if jobs, err := s.JobsBySkill("ruby"); err != nil {
		t.Fatal(err)
	} else if jobs[0].Name != "Hacker" {
		t.Fatalf("unexpected name: %s", jobs[0].Name)
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

func (t *TestStore) CleanDatabase() {
	table := []string{"categories", "jobs", "skills"}
	for i := 0; i < len(table); i++ {
		sess := t.DB.NewSession(nil)
		sess.DeleteFrom(table[i]).Exec()
	}
}
