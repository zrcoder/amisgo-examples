package page

import (
	"amisgo-examples/todo-app/api"
	"amisgo-examples/todo-app/db"

	"github.com/zrcoder/amisgo/comp"
)

func List() any {
	return comp.Page().Body(
		comp.Button().Icon("fa fa-plus").Level("primary").ActionType("drawer").Drawer(detail("", "")),
		comp.Crud().Name("todos").Api(api.Todos).SyncLocation(false).Columns(
			comp.Column().Name("title").Label("Title"),
			comp.Column().Name("created_at").Label("Create Time"),
			comp.Column().Name("updated_at").Label("Update Time"),
			comp.Column().Type("operation").Buttons(
				comp.Button().Icon("fa fa-eye").Label("Detail").ActionType("drawer").Drawer(detail(api.Todo+"?id=${id}", "")),
				comp.Button().Icon("fa fa-edit").Label("Edit").ActionType("drawer").Drawer(detail(api.Todo+"?id=${id}", "patch:"+api.Todo+"?id=${id}")),
				comp.Button().Icon("fa fa-trash").Level("danger").Label("Delete").ActionType("ajax").ConfirmText("Delete this task?").Api("delete:"+api.Todo+"?id=${id}").ReloadWindow(),
			),
		),
	)
}

func detail(getApi, editApi string) any {
	readOnly := getApi != "" && editApi == ""
	form := comp.Form().
		Static(readOnly).
		WrapWithPanel(false).
		Body(
			comp.InputText().Name("title").Placeholder("Title"),
			comp.InputRichText().Name("detail").Options(comp.Schema{
				"height":  600,
				"plugins": "code", // don't delete this line, it effects the height of the component
				"menubar": "",
			}),
		).
		Rules(
			comp.Rule().Rule("data.title && data.detail").Message("Both title and content can't be empty"),
		)

	if !readOnly {
		form.Reload("todos")
	}

	if getApi == "" { // add new task
		form.Go(func(d comp.Data) error {
			_, err := db.AddTodo(d.Get("title").(string), d.Get("detail").(string))
			return err
		})
	} else {
		form.InitApi(getApi)
	}
	if editApi != "" {
		form.Api(editApi)
	}

	return comp.Drawer().
		Size("xl").
		ShowCloseButton(false).
		CloseOnOutside(true).
		Confirm(!readOnly).
		Body(form)

}
