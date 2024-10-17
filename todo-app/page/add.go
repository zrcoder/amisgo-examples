package page

import (
	"amisgo-examples/todo-app/db"
	"amisgo-examples/todo-app/db/data"
	"context"
	"database/sql"

	"github.com/zrcoder/amisgo/comp"
)

var Add = comp.Form().Title("Add Todo").Body(
	comp.InputText().Name("title").Placeholder("Title"),
	comp.InputRichText().Name("detail").Size("lg"),
).Go(func(d comp.Data) error {
	ctx := context.Background()
	todo, err := db.Query.AddTodo(ctx, data.AddTodoParams{
		Title: d.Get("title").(string),
	})
	if err != nil {
		return err
	}
	_, err = db.Query.AddTodoDetail(ctx, data.AddTodoDetailParams{
		TodoID: todo.ID,
		Detail: sql.NullString{Valid: true, String: d.Get("detail").(string)},
	})
	return err
})
