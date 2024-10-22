package db

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"time"

	"amisgo-examples/todo-app/model"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var cinitTablesSql string

var db *sql.DB

func init() {
	ctx := context.Background()
	var err error
	db, err = sql.Open("sqlite3", "./db/todo.db")
	if err != nil {
		log.Panic(err)
	}

	// create tables
	if _, err := db.ExecContext(ctx, cinitTablesSql); err != nil {
		log.Panic(err)
	}
}

const addTodo = `-- name: AddTodo :one
INSERT INTO todos (title, priority, due_date, is_completed, created_at, updated_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, title, priority, due_date, is_completed, created_at, updated_at
`

const addTodoDetail = `-- name: AddTodoDetail :one
INSERT INTO todo_details (todo_id, detail)
VALUES (?, ?)
RETURNING id, todo_id, detail
`

func AddTodo(title, content string) (*model.Todo, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	row := tx.QueryRow(addTodo, title, 1, time.Time{}, false)
	todo := &model.Todo{}
	err = row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Priority,
		&todo.DueDate,
		&todo.IsCompleted,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	detail := &model.TodoDetail{}
	row = tx.QueryRow(addTodoDetail, todo.ID, content)
	err = row.Scan(&detail.ID, &detail.TodoID, &detail.Detail)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	return todo, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = ?
`

func DeleteTodo(id int64) error {
	_, err := db.Exec(deleteTodo, id)
	return err
}

const getTodoFull = `-- name: GetTodo :one
SELECT todos.id, todos.title, todos.priority, todos.due_date, todos.is_completed, todos.created_at, todos.updated_at, todo_details.detail
FROM todos
LEFT JOIN todo_details ON todos.id = todo_details.todo_id
WHERE todos.id = ? LIMIT 1
`

func GetTodoFull(id int64) (*model.TodoFull, error) {
	row := db.QueryRow(getTodoFull, id)
	todo := &model.TodoFull{}
	err := row.Scan(
		&todo.Todo.ID,
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

const listTodos = `-- name: ListTodos :many
SELECT id, title, priority, due_date, is_completed, created_at, updated_at FROM todos
`

func ListTodos() ([]model.Todo, error) {
	rows, err := db.Query(listTodos)
	if err != nil {
		return nil, err
	}
	var todos []model.Todo
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
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
