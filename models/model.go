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
	ID   int
	Name string
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
	if err := sess.Select("id", "name").From("jobs").Where("name= ?", name).LoadStruct(&j); err != nil {
		return nil, err
	}

	return &j, nil
}

// Skill represents a skill posted in a job
type Skill struct {
	ID   int
	Name string
}

// CreateSkill creates a new skill
func (s *Store) CreateSkill(skill *Skill) error {
	sess := s.DB.NewSession(nil)
	_, err := sess.InsertInto("skills").Columns("name").Record(skill).Exec()
	if err != nil {
		return err
	}
	return nil
}

// Skill returns a skill by name
func (s *Store) Skill(name string) (*Skill, error) {
	sess := s.DB.NewSession(nil)
	var skill Skill
	if err := sess.Select("id", "name").From("skills").Where("name = ?", name).LoadStruct(&skill); err != nil {
		return nil, err
	}

	return &skill, nil
}

// Category joins Jobs and Skills together
type Category struct {
	ID      int
	JobID   int
	SkillID int
}

// CreateCategory creates a new join between Jobs and Skills
func (s *Store) CreateCategory(c *Category) error {
	sess := s.DB.NewSession(nil)
	if _, err := sess.InsertInto("categories").Columns("job_id", "skill_id").Record(c).Exec(); err != nil {
		return err
	}
	return nil
}

// JobsBySkill returns jobs that have a particular skill
func (s *Store) JobsBySkill(skill string) ([]Job, error) {
	sess := s.DB.NewSession(nil)
	var jobs []Job
	_, err := sess.
		Select("*").
		From("skills").
		Where("skills.name = ?", skill).
		Join("categories", "categories.skill_id = skills.id").
		Join("jobs", "jobs.id = categories.job_id").
		Load(&jobs)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// SkillsByJob returns jobs that have a particular skill
func (s *Store) SkillsByJob(job string) ([]Skill, error) {
	sess := s.DB.NewSession(nil)
	var skills []Skill
	_, err := sess.
		Select("*").
		From("jobs").
		Where("jobs.name = ?", job).
		Join("categories", "categories.job_id = jobs.id").
		Join("skills", "skills.id = categories.skill_id").
		Load(&skills)
	if err != nil {
		return nil, err
	}
	return skills, nil
}

// Jobs returns all jobs
func (s *Store) Jobs() ([]Job, error) {
	sess := s.DB.NewSession(nil)
	var jobs []Job
	_, err := sess.
		Select("*").
		From("jobs").
		Load(&jobs)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

// Skills returns all jobs
func (s *Store) Skills() ([]Skill, error) {
	sess := s.DB.NewSession(nil)
	var skills []Skill
	_, err := sess.
		Select("*").
		From("skills").
		Load(&skills)
	if err != nil {
		return nil, err
	}
	return skills, nil
}
