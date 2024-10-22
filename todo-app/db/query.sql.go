package db

/*
func (q *Queries) UodateTodoDetail(ctx context.Context, todoID int64, content string) error {
	const uodateTodoDetail = `-- name: UodateTodoDetail :exec
UPDATE todo_details
SET detail = ?, updated_at = CURRENT_TIMESTAMP
WHERE todo_id = ?
`
	_, err := q.db.ExecContext(ctx, uodateTodoDetail, content, todoID)
	return err
}

func (q *Queries) UpdateTodo(ctx context.Context, arg Todo) error {
	const updateTodo = `-- name: UpdateTodo :exec
UPDATE todos
SET title = ?, priority = ?, due_date = ?, is_completed = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`
	_, err := q.db.ExecContext(ctx, updateTodo,
		arg.Title,
		arg.Priority,
		arg.DueDate,
		arg.IsCompleted,
		arg.ID,
	)
	return err
}
*/
