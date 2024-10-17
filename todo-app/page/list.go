package page

import (
	"amisgo-examples/todo-app/db"
	"context"

	"github.com/zrcoder/amisgo/comp"
)

var List = comp.Page().Body(
	comp.Button().Label("New").ActionType("drawer").Drawer(
		comp.Drawer().Body(Add).Size("lg"),
	),
	comp.Service().GetData(func() (any, error) {
		todos, err := db.Query.ListTodos(context.Background())
		if err != nil {
			return nil, err
		}
		return comp.Response{Data: comp.Data{"items": todos, "total": len(todos)}}, nil
	}).Body(comp.Crud().Title("Todos").Columns(
		comp.Column().Name("Title"),
		comp.Column().Type("operation").Buttons(
			comp.Button().Level("danger").Label("Delete").ActionType("dialog").Dialog(comp.Dialog().Title("Delete this task?")),
		),
	)),
)
