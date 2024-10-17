-- name: AddTodo :one
INSERT INTO todos (title, priority, due_date, is_completed, created_at, updated_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: AddTodoDetail :one
INSERT INTO todo_details (todo_id, detail, created_at, updated_at)
VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING *;

-- name: GetTodo :one
SELECT todos.*, todo_details.detail
FROM todos
LEFT JOIN todo_details ON todos.id = todo_details.todo_id
WHERE todos.id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY updated_at;


-- name: UpdateTodo :exec
UPDATE todos
SET title = ?, priority = ?, due_date = ?, is_completed = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UodateTodoDetail :exec
UPDATE todo_details
SET detail = ?, updated_at = CURRENT_TIMESTAMP
WHERE todo_id = ?;


-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?;
