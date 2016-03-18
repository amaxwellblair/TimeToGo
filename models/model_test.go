package models

import "testing"

// DATASOURCETEST is the db name and authentication for testing
const DATASOURCETEST = "user=maxwell dbname=careerCrawl_test sslmode=disable"

func TestStore_Open(t *testing.T) {
	s := NewStore()
	if err := s.Open(DATASOURCETEST); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	defer s.Close()
	if err := s.DB.Ping(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestStore_CreateJob(t *testing.T) {
	s := NewStore()
	if err := s.Open(DATASOURCETEST); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	defer s.Close()

	if err := s.CreateJob(&Job{
		Name:       "Hacker",
		Categories: []string{"python", "ruby", "c"},
	}); err != nil {
		t.Fatal(err)
	}
	sess := s.DB.NewSession(nil)
	defer sess.DeleteFrom("jobs").Where("name = ?", "Hacker").Exec()

	if job, err := s.Job("Hacker"); err != nil {
		t.Fatal(err)
	} else if job.Name != "Hacker" {
		t.Fatalf("unexpected name: %s", job.Name)
	}
}
