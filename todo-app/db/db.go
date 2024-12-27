package db

import (
	_ "embed"
	"fmt"
	"log/slog"

	"github.com/zrcoder/amisgo-examples/todo-app/model"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

//go:embed schema.sql
var createTablesSql string

var db *sqlx.DB

func Init() error {
	var err error
	db, err = sqlx.Open("sqlite", "todo.db?_pragma=foreign_keys(1)")
	if err != nil {
		return err
	}

	// create tables
	_, err = db.Exec(createTablesSql)
	return err
}

func Close() {
	err := db.Close()
	slog.Info("db closed", "error", err)
}

const addTodo = `
INSERT INTO todos (title, priority, due_date, is_completed, detail, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
`

func AddTodo(todoInput *model.Todo) error {
	_, err := db.Exec(addTodo,
		todoInput.Title, todoInput.Priority, todoInput.DueDate, todoInput.IsCompleted, todoInput.Detail)
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
WHERE id = ?
`

func UpdateTodo(todo *model.Todo) error {
	_, err := db.Exec(updateTodo, todo.Title, todo.Priority, todo.DueDate, todo.IsCompleted, todo.Detail, todo.ID)
	return err
}

const getTodo = `
SELECT todos.id, todos.title, todos.priority, todos.due_date, todos.is_completed, todos.created_at, todos.updated_at, todos.detail
FROM todos
WHERE todos.id = ? LIMIT 1
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

const (
	getTotal  = `SELECT COUNT(*) FROM todos`
	listTodos = `
SELECT id, title, priority, due_date, is_completed, created_at, updated_at
FROM todos
LIMIT ? OFFSET ?
`
)

func ListTodos(limit, offset int) ([]model.Todo, int, error) {
	total := 0
	err := db.QueryRow(getTotal).Scan(&total)
	if err != nil {
		return nil, 0, err
	}
	rows, err := db.Query(listTodos, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	todos := make([]model.Todo, 0, limit)
	for rows.Next() {
		var todo model.Todo
		if err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Priority,
			&todo.DueDate,
			&todo.IsCompleted,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		); err != nil {
			rows.Close()
			return nil, 0, err
		}
		todos = append(todos, todo)
	}
	if err = rows.Close(); err != nil {
		return nil, 0, err
	}
	if err := rows.Err(); err != nil {
		return nil, 0, err
	}
	return todos, total, nil
}
