package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/zrcoder/amisgo-examples/todo-app/model"
)

const addTodo = `
INSERT INTO todos (title, priority, due_date, is_completed, detail, created_at, updated_at, user_id)
VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, ?)
`

func AddTodo(todoInput *model.Todo) error {
	_, err := db.Exec(addTodo,
		todoInput.Title, todoInput.Priority, todoInput.DueDate, todoInput.IsCompleted, todoInput.Detail, todoInput.UserID)
	return err
}

const deleteTodo = `DELETE FROM todos WHERE id IN (?)`

func DeleteTodos(ids []int64) error {
	if len(ids) == 0 {
		return nil
	}

	query, args, err := sqlx.In(deleteTodo, ids)
	if err != nil {
		return fmt.Errorf("prepare query error: %w", err)
	}

	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error deleting todos: %w", err)
	}
	return nil
}

const updateTodo = `
UPDATE todos
SET title = ?, priority = ?, due_date = ?, is_completed = ?, detail = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ? and user_id = ?
`

func UpdateTodo(todo *model.Todo) error {
	_, err := db.Exec(updateTodo, todo.Title, todo.Priority, todo.DueDate, todo.IsCompleted, todo.Detail, todo.ID, todo.UserID)
	return err
}

const getTodo = `
SELECT id, title, priority, due_date, is_completed, created_at, updated_at, detail
FROM todos
WHERE id = ? LIMIT 1
`

func GetTodo(id int64) (*model.Todo, error) {
	row := db.QueryRow(getTodo, id)
	todo := &model.Todo{}
	err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Priority,
		&todo.DueDate,
		&todo.IsCompleted,
		&todo.CreatedAt,
		&todo.UpdatedAt,
		&todo.Detail,
	)
	return todo, err
}

func ListTodos(params *model.ListRequest) ([]model.Todo, int, error) {
	whereClause, args := buildWhereClause(params)
	orderClause := buildOrderClause(params)

	args["limit"] = params.Limit
	args["offset"] = params.Offset

	total, err := getTotalCount(whereClause, args)
	if err != nil {
		return nil, 0, err
	}

	todos, err := getTodos(whereClause, orderClause, args)
	if err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}

func buildWhereClause(params *model.ListRequest) (string, map[string]any) {
	conditions := []string{"user_id = :user_id"}
	args := map[string]any{"user_id": params.UserID}

	if params.TitleKeywords != "" {
		conditions = append(conditions, "title LIKE :title")
		args["title"] = "%" + params.TitleKeywords + "%"
	}
	if params.IsCompleted != "" {
		if completed, err := strconv.ParseBool(params.IsCompleted); err == nil {
			conditions = append(conditions, "is_completed = :is_completed")
			args["is_completed"] = completed
		}
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	return whereClause, args
}

func buildOrderClause(params *model.ListRequest) string {
	if params.OrderBy == "" {
		return ""
	}

	validOrderFields := map[string]bool{
		"created_at": true,
		"due_date":   true,
		"title":      true,
		"priority":   true,
	}

	if !validOrderFields[params.OrderBy] {
		return ""
	}

	orderDir := "ASC"
	if params.OrderDir == "desc" {
		orderDir = "DESC"
	}

	return fmt.Sprintf("ORDER BY %s %s", params.OrderBy, orderDir)
}

func getTotalCount(whereClause string, args map[string]any) (int, error) {
	countQuery := "SELECT COUNT(*) FROM todos " + whereClause
	countQuery, countArgs, err := sqlx.Named(countQuery, args)
	if err != nil {
		return 0, err
	}
	countQuery = db.Rebind(countQuery)

	var total int
	err = db.Get(&total, countQuery, countArgs...)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func getTodos(whereClause, orderClause string, args map[string]any) ([]model.Todo, error) {
	listQuery := fmt.Sprintf(`
		SELECT id, title, priority, due_date, is_completed, created_at, updated_at, user_id
		FROM todos
		%s
		%s
		LIMIT :limit OFFSET :offset
		`, whereClause, orderClause)

	listQuery, listArgs, err := sqlx.Named(listQuery, args)
	if err != nil {
		return nil, err
	}
	listQuery = db.Rebind(listQuery)

	todos := []model.Todo{}
	err = db.Select(&todos, listQuery, listArgs...)
	if err != nil {
		return nil, err
	}

	return todos, nil
}
