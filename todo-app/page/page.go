package page

import (
	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo-examples/todo-app/db"
	"github.com/zrcoder/amisgo-examples/todo-app/model"

	"github.com/zrcoder/amisgo/comp"
)

func List() any {
	return comp.Page().Body(
		comp.Button().Icon("fa fa-plus").Level("primary").ActionType("drawer").Drawer(detail("", "")),
		comp.Crud().Name("todos").Api(api.Todos).SyncLocation(false).Columns(
			comp.Column().Name("title").Label("Title"),
			comp.Column().Name("created_at").Label("Create Time").Type("datetime"),
			comp.Column().Name("updated_at").Label("Update Time").Type("datetime"),
			comp.Column().Name("is_completed").Label("Done").Type("status"),
			comp.Column().Type("operation").Buttons(
				// comp.Button().Icon("fa fa-eye").Label("Detail").ActionType("drawer").Drawer(detail(api.Todo+"?id=${id}", "")),
				comp.Button().Icon("fa fa-edit").Label("Edit").ActionType("drawer").Drawer(detail(api.Todo+"?id=${id}", "patch:"+api.Todo+"?id=${id}")),
				comp.Button().Icon("fa fa-trash").Level("danger").Label("Delete").ActionType("ajax").ConfirmText("Delete this task?").Api("delete:"+api.Todo+"?id=${id}").ReloadWindow(),
			),
		),
	)
}

func detail(getApi, editApi string) any {
	readOnly := getApi != "" && editApi == ""
	isCreate := getApi == ""
	var content any = comp.Markdown().Options(comp.Schema{"html": true}).Name("detail")
	if !readOnly {
		content = comp.Group().Body(
			content,
			comp.Editor().Name("detail").Language("markdown").Size("xxl").Value("${detail}"),
		)
	}
	var state any = comp.Switch().Name("is_completed").Label("Done")
	if isCreate {
		state = comp.Checkbox().Name("is_completed").Label("Done").Disabled(true)
	}
	form := comp.Form().
		Static(readOnly).
		AutoFocus(true).
		WrapWithPanel(false).
		Body(
			comp.Group().Body(
				state,
				comp.InputText().Name("title").Placeholder("Title"),
			),
			content,
		).
		Rules(
			comp.Rule().Rule("data.title && data.detail").Message("Both title and content can't be empty"),
		)
	if !readOnly {
		form.Reload("todos")
	}
	if isCreate {
		form.Go(func(d comp.Data) error {
			todo := &model.TodoFull{}
			todo.Title = d.Get("title").(string)
			todo.Detail = d.Get("detail").(string)
			_, err := db.AddTodo(todo)
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
