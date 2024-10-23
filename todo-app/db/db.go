package db

import (
	"database/sql"
	_ "embed"
	"time"

	"amisgo-examples/todo-app/model"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var createTablesSql string

var db *sql.DB

func Init() error {
	var err error
	db, err = sql.Open("sqlite3", "./db/todo.db")
	if err != nil {
		return err
	}

	// create tables
	_, err = db.Exec(createTablesSql)
	return err
}

const addTodo = `
INSERT INTO todos (title, priority, due_date, is_completed, created_at, updated_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, title, priority, due_date, is_completed, created_at, updated_at
`

const addTodoDetail = `
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
	err = row.Scan(&detail.DetailID, &detail.TodoID, &detail.Detail)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	return todo, err
}

const deleteTodo = `DELETE FROM todos WHERE id = ?`

func DeleteTodo(id int64) error {
	_, err := db.Exec(deleteTodo, id)
	return err
}

const updateTodo = `
UPDATE todos
SET title = ?, priority = ?, due_date = ?, is_completed = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`
const uodateTodoDetail = `
UPDATE todo_details
SET detail = ?
WHERE todo_id = ?
`

func UpdateTodo(todoID int64, title, content string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.Exec(updateTodo, title, 0, time.Time{}, false, todoID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(uodateTodoDetail, content, todoID)
	if err != nil {
		return err
	}
	return tx.Commit()
}

const getTodoFull = `
SELECT todos.id, todos.title, todos.priority, todos.due_date, todos.is_completed, todos.created_at, todos.updated_at, todo_details.detail
FROM todos
LEFT JOIN todo_details ON todos.id = todo_details.todo_id
WHERE todos.id = ? LIMIT 1
`

func GetTodoFull(id int64) (*model.TodoFull, error) {
	row := db.QueryRow(getTodoFull, id)
	todo := &model.TodoFull{}
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

const getTotal = `SELECT COUNT(*) FROM todos`
const listTodos = `
SELECT id, title, priority, due_date, is_completed, created_at, updated_at
FROM todos
LIMIT ? OFFSET ?
`

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
