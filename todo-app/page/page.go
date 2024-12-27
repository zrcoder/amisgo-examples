package page

import (
	"encoding/json"
	"fmt"

	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo-examples/todo-app/db"
	"github.com/zrcoder/amisgo-examples/todo-app/model"

	"github.com/zrcoder/amisgo/comp"
)

func List() any {
	return comp.Flex().Items(
		comp.Page().ClassName("py-20 w-3/4").Title("My Todos").Body(
			comp.Crud().Name("todos").Api(api.Todos).SyncLocation(false).Columns(
				comp.Column().Name("title").Label("Title"),
				comp.Column().Name("due_date").Label("Due Date").Type("date"),
				comp.Column().Name("is_completed").Label("Done").Type("status"),
				comp.Column().Name("created_at").Label("Create Time").Type("datetime"),
			).OnEvent(
				comp.Schema{
					"rowClick": comp.Schema{
						"actions": []comp.MEventAction{
							comp.EventAction().ActionType("drawer").Drawer(
								detail(api.Todo+"?id=${event.data.item.id}", "patch:"+api.Todo+"?id=${event.data.item.id}"),
							),
						},
					},
				},
			).
				HeaderToolbar(
					comp.Button().Icon("fa fa-plus").Label("Add").ActionType("drawer").Drawer(detail("", "")),
				).
				FooterToolbar(
					"bulkActions",
					"pagination",
				).
				BulkActions(
					comp.Button().Icon("fa fa-trash").Level("danger").Label("Delete").ActionType("ajax").ConfirmText("Delete the tasks?").Api("delete:" + api.Todo + "?ids=${ids}").ReloadWindow(),
				),
		),
	)
}

func detail(getApi, editApi string) any {
	isCreate := getApi == ""

	form := comp.Form().
		Mode("normal").
		AutoFocus(true).
		WrapWithPanel(false).
		Body(
			comp.Group().Body(
				comp.InputText().Name("title").Placeholder("Title").Label("Title"),
				comp.InputDatetime().Name("due_date").Label("Due Date").Value("+1days").ValueFormat("YYYY-MM-DDTHH:mm:ssZ"),
			),
			comp.Switch().Name("is_completed").Option("Done").Disabled(isCreate),
			comp.Group().Label("Detail").Body(
				comp.Editor().Name("detail").Language("markdown").Size("xxl").Value("${detail}").AllowFullscreen(false).Options(
					comp.Schema{
						"overviewRulerBorder": false,
					},
				),
				comp.Markdown().Options(comp.Schema{"html": true}).Name("detail"),
			),
		).
		Rules(
			comp.Rule().Rule("data.title && data.detail").Message("Both title and content can't be empty"),
		).
		Reload("todos")

	if isCreate {
		form.Go(func(d comp.Data) error {
			todo := &model.Todo{}
			fmt.Println("---", string(d.Json()))
			err := json.Unmarshal(d.Json(), todo)
			if err != nil {
				return err
			}
			return db.AddTodo(todo)
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
		Body(form)
}
