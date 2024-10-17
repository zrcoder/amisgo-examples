package db

import (
	"context"
	"database/sql"
	_ "embed"
	"log"

	"amisgo-examples/todo-app/db/data"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

var Query *data.Queries

func init() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./db/todo.db")
	if err != nil {
		log.Panic(err)
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Panic(err)
	}

	Query = data.New(db)
}
