package model

import (
	"strconv"
	"time"
)

type Todo struct {
	ID          int64     `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Priority    int64     `json:"priority" db:"priority"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
	IsCompleted bool      `json:"is_completed" db:"is_completed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Detail      string    `json:"detail" db:"detail"`
}

type ListRequest struct {
	TitleKeywords string `form:"title"`
	IsCompleted   string `form:"is_completed"`
	OrderBy       string `form:"orderBy"`
	OrderDir      string `form:"orderDir"`
	Page          string `form:"page"`
	PerPage       string `form:"perPage"`
	Limit         int    `form:"-"`
	Offset        int    `form:"-"`
}

func (r *ListRequest) Regular() {
	r.Limit, r.Offset = 10, 0
	if v, err := strconv.Atoi(r.PerPage); err == nil {
		r.Limit = int(v)
	}
	if v, err := strconv.Atoi(r.Page); err == nil {
		r.Offset = (int(v) - 1) * r.Limit
	}
}
