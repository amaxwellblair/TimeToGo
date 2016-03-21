package models_test

import (
	"testing"

	"github.com/amaxwellblair/TimeToGo/models"
	"github.com/amaxwellblair/TimeToGo/testHelpers"
)

func TestStore_Open(t *testing.T) {
	s := testhelper.NewTestStore()
	s.MustOpen()
	defer s.Close()
}

func TestStore_CreateJob(t *testing.T) {
	s := testhelper.NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()

	if err := s.CreateJob(&models.Job{
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
	s := testhelper.NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()

	if err := s.CreateSkill(&models.Skill{
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
	s := testhelper.NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()

	s.CreateIntegration(t)

	jobs, err := s.JobsBySkill("ruby")
	if err != nil {
		t.Fatal(err)
	}
	if len(jobs) != 2 {
		t.Fatalf("unexpected length: %#v", len(jobs))
	}
	for i := 0; i < len(jobs); i++ {
		if jobs[i].Name != "Hacker" && jobs[i].Name != "QA" {
			t.Fatalf("unexpected name: %s", jobs[i].Name)
		}
	}
}

func TestStore_CreateCategory_SkillsByJob(t *testing.T) {
	s := testhelper.NewTestStore()
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
	s := testhelper.NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()
	actualJobs, _ := s.CreateIntegration(t)

	expectJobs, err := s.Jobs()
	if err != nil {
		t.Fatal(err)
	}

	compare := func(j models.Job, jobs []*models.Job) bool {
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
	s := testhelper.NewTestStore()
	s.MustOpen()
	defer s.Close()
	defer s.CleanDatabase()
	_, actualSkills := s.CreateIntegration(t)

	expectSkills, err := s.Skills()
	if err != nil {
		t.Fatal(err)
	}

	compare := func(s models.Skill, skills []*models.Skill) bool {
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
