package page

import (
	"amisgo-examples/todo-app/db"
	"log"
	"strconv"

	"github.com/zrcoder/amisgo/comp"
)

func List() any {
	return comp.Page().Body(
		comp.Button().Icon("fa fa-plus").Level("primary").ActionType("drawer").Drawer(editDrawer("", false)),
		comp.Service().Body(
			comp.Crud().Columns(
				comp.Column().Name("title").Label("Title"),
				comp.Column().Name("is_complated").Label("Complated").Body(comp.Checkbox().OptionType("button").Name("is_complated")),
				comp.Column().Name("created_at").Label("Create Time"),
				comp.Column().Name("update_at").Label("Update Time"),
				comp.Column().Type("operation").Buttons(
					comp.Button().Level("danger").Label("Delete").ActionType("dialog").Dialog(comp.Dialog().Title("Delete this task?")),
					comp.Button().Label("Detail").Icon("fa fa-detail").ActionType("drawer").Drawer(editDrawer("${ID}", true)),
				),
			),
		).GetData(func() (any, error) {
			todos, err := db.ListTodos()
			if err != nil {
				return nil, err
			}
			return comp.Response{Data: comp.Data{"items": todos, "total": len(todos)}}, nil
		}),
	)
}

func editDrawer(todoID string, readOnly bool) any {
	res := comp.Drawer().
		Size("xl").
		Position("top").
		ShowCloseButton(false).
		Confirm(!readOnly).
		Body(
			comp.Form().
				Static(readOnly).
				Body(
					comp.InputText().Name("title").Placeholder("Title"),
					comp.InputRichText().Name("detail").Size("lg"),
				).
				Rules(
					comp.Rule().Rule("data.title && data.detail").Message("Both title and content can't be empty"),
				).
				Go(func(d comp.Data) error {
					_, err := db.AddTodo(d.Get("title").(string), d.Get("detail").(string))
					return err
				}),
		)
	if todoID != "" {
		res.GetData(todoID, func(id any) comp.Data {
			log.Println("todo id:", id)
			istr := id.(string)
			iint, _ := strconv.ParseInt(istr, 10, 64)
			todo, _ := db.GetTodoFull(iint)
			res, _ := comp.ParseData(todo)
			return res
		})
	}
	return res
}
