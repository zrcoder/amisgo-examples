package db

import (
	_ "embed"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/zrcoder/amisgo-examples/todo-app/model"
	"github.com/zrcoder/amisgo-examples/todo-app/util"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

//go:embed schema.sql
var createTablesSql string

var db *sqlx.DB

func Init() error {
	dbName := "todo.db"
	if util.ReadOnly() {
		dbName = "todo-sample.db"
	}
	var err error
	db, err = sqlx.Open("sqlite", dbName+"?_pragma=foreign_keys(1)")
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

func ListTodos(params *model.ListRequest) ([]model.Todo, int, error) {
	conditions := []string{}
	args := map[string]any{}

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

	orderClause := ""
	if params.OrderBy != "" {
		validOrderFields := map[string]bool{
			"created_at": true,
			"due_date":   true,
			"title":      true,
			"priority":   true,
		}

		if validOrderFields[params.OrderBy] {
			orderDir := "ASC"
			if params.OrderDir == "desc" {
				orderDir = "DESC"
			}
			orderClause = fmt.Sprintf("ORDER BY %s %s", params.OrderBy, orderDir)
		}
	}

	args["limit"] = params.Limit
	args["offset"] = params.Offset

	countQuery := "SELECT COUNT(*) FROM todos " + whereClause
	countQuery, countArgs, err := sqlx.Named(countQuery, args)
	if err != nil {
		return nil, 0, err
	}
	countQuery = db.Rebind(countQuery)

	var total int
	err = db.Get(&total, countQuery, countArgs...)
	if err != nil {
		return nil, 0, err
	}

	listQuery := fmt.Sprintf(`
		SELECT id, title, priority, due_date, is_completed, created_at, updated_at
		FROM todos
		%s
		%s
		LIMIT :limit OFFSET :offset
	`, whereClause, orderClause)

	listQuery, listArgs, err := sqlx.Named(listQuery, args)
	if err != nil {
		return nil, 0, err
	}
	listQuery = db.Rebind(listQuery)
	todos := []model.Todo{}
	err = db.Select(&todos, listQuery, listArgs...)
	if err != nil {
		return nil, 0, err
	}
	return todos, total, nil
}
