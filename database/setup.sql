DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS daily_todos;

CREATE TABLE todos (
  id   INTEGER PRIMARY KEY,
  name text    NOT NULL,
  iscomplete bool NOT NULL,
  priority int NOT NULL DEFAULT 1,
  date_created   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  due_date   TIMESTAMP NOT NULL
);

CREATE TABLE daily_todos (
  id   INTEGER PRIMARY KEY,
  name text    NOT NULL,
  iscomplete bool NOT NULL,
  priority int NOT NULL DEFAULT 1,
  date_created   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ADD SOME FAKE DATA
INSERT INTO todos (name, iscomplete, due_date) VALUES ('Learn SQL', false, '2024-09-01 14:04:30');
INSERT INTO daily_todos (name, iscomplete) VALUES ('Learn SQL', false);
SELECT * FROM todos;
SELECT * FROM daily_todos;
