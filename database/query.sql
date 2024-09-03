-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY due_date, priority DESC;

-- name: ListTodaysTodos :many
SELECT * FROM todos
WHERE date(due_date) = date('now')
ORDER BY due_date, priority DESC;

-- name: CreateTodo :one
INSERT INTO todos (
name, iscomplete, priority, due_date
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: CompleteTodo :exec
UPDATE todos
set iscomplete = ?
WHERE id = ?;

-- name: UpdateTodoDueDate :exec
UPDATE todos
set due_date = ?
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;
