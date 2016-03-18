CREATE TABLE categories(
  id SERIAL,
  job_id int,
  skill_id int,
  FOREIGN KEY (job_id) REFERENCES jobs (id),
  FOREIGN KEY (skill_id) REFERENCES skills (id),
  PRIMARY KEY (id)
);
