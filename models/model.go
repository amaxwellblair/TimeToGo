package models

import (
	"github.com/gocraft/dbr"

	// Postgres driver
	_ "github.com/lib/pq"
)

// Store holds the database
type Store struct {
	DB *dbr.Connection
}

// NewStore returns a new Store
func NewStore() *Store {
	return &Store{}
}

// Open a new database
func (s *Store) Open(dataSource string) error {
	var err error
	s.DB, err = dbr.Open("postgres", dataSource, nil)
	return err
}

// Close disconnects from the database
func (s *Store) Close() error {
	return s.DB.Close()
}

// Job represents a job posting
type Job struct {
	Name       string
	Categories []string
}

// CreateJob creates a new job
func (s *Store) CreateJob(j *Job) error {
	sess := s.DB.NewSession(nil)
	if _, err := sess.InsertInto("jobs").Columns("name").Record(j).Exec(); err != nil {
		return err
	}
	return nil
}

// Job retrieves a specific job by name
func (s *Store) Job(name string) (*Job, error) {
	sess := s.DB.NewSession(nil)
	var j Job
	if err := sess.Select("name").From("jobs").Where("name= ?", name).LoadStruct(&j); err != nil {
		return nil, err
	}

	return &j, nil
}

// Skill represents a skill posted in a job
type Skill struct {
	Name string
}

// CreateSkill creates a new skill
func (s *Store) CreateSkill(skill *Skill) (*int64, error) {
	return nil, nil
}
