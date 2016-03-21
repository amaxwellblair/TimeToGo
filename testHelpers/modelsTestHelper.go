package testhelper

import (
	"testing"

	"github.com/amaxwellblair/TimeToGo/models"
)

// DATASOURCETEST is the db name and authentication for testing
const DATASOURCETEST = "user=maxwell dbname=careerCrawl_test sslmode=disable"

// TestStore will serve as a wrapper for our Store
type TestStore struct {
	models.Store
}

// NewTestStore creates a new TestStore
func NewTestStore() *TestStore {
	return &TestStore{}
}

// MustOpen opens a new store or panics
func (t *TestStore) MustOpen() {
	if err := t.Open(DATASOURCETEST); err != nil {
		panic(err)
	}
	if err := t.DB.Ping(); err != nil {
		panic(err)
	}
}

// CreateIntegration creates new data for testing
func (t *TestStore) CreateIntegration(test *testing.T) ([]*models.Job, []*models.Skill) {
	if err := t.CreateSkill(&models.Skill{
		Name: "ruby",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateSkill(&models.Skill{
		Name: "go",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateSkill(&models.Skill{
		Name: "python",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateJob(&models.Job{
		Name: "Hacker",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateJob(&models.Job{
		Name: "Coder",
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateJob(&models.Job{
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

	if err := t.CreateCategory(&models.Category{
		JobID:   hacker.ID,
		SkillID: ruby.ID,
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateCategory(&models.Category{
		JobID:   qa.ID,
		SkillID: ruby.ID,
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateCategory(&models.Category{
		JobID:   coder.ID,
		SkillID: golang.ID,
	}); err != nil {
		test.Fatal(err)
	} else if err := t.CreateCategory(&models.Category{
		JobID:   coder.ID,
		SkillID: python.ID,
	}); err != nil {
		test.Fatal(err)
	}
	return []*models.Job{hacker, qa, coder}, []*models.Skill{ruby, golang, python}
}

// CleanDatabase deletes all of the data from the Database
func (t *TestStore) CleanDatabase() {
	table := []string{"categories", "jobs", "skills"}
	for i := 0; i < len(table); i++ {
		sess := t.DB.NewSession(nil)
		sess.DeleteFrom(table[i]).Exec()
	}
}
