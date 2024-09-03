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
