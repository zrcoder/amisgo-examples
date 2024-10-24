package model

import (
	"time"
)

type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Priority    int64     `json:"priority"`
	DueDate     time.Time `json:"due_date"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoDetail struct {
	TodoID int64  `json:"todo_id"`
	Detail string `json:"detail"`
}

type TodoFull struct {
	Todo
	TodoDetail
}
