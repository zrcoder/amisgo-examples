package db

import (
	_ "embed"
	"log"

	"github.com/zrcoder/amisgo-examples/todo-app/util"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

//go:embed schema.sql
var prepareSQL string

var db *sqlx.DB

func init() {
	dbName := "todo.db"
	if util.ReadOnly() {
		dbName = "todo-sample.db"
	}
	var err error
	db, err = sqlx.Open("sqlite", dbName+"?_pragma=foreign_keys(1)")
	if err != nil {
		log.Fatal("open db error:", err)
	}

	// create tables
	_, err = db.Exec(prepareSQL)
	if err != nil {
		log.Fatal("create tables error:", err)
	}
}

func Close() error {
	return db.Close()
}
