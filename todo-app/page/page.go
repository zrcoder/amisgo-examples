package page

import (
	"errors"

	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo-examples/todo-app/db"
	"github.com/zrcoder/amisgo-examples/todo-app/model"
	"github.com/zrcoder/amisgo-examples/todo-app/util"

	"github.com/zrcoder/amisgo/comp"
	am "github.com/zrcoder/amisgo/model"
)

func Index() any {
	return comp.Page().ClassName("p-8").Title(comp.Tpl().Tpl("TODOs").ClassName("font-bold")).Body(
		comp.Crud().Name("todos").Api(api.Todos).SyncLocation(false).
			Columns(
				comp.Column().Name("is_completed").Label("Done").Type("status"),
				comp.Column().Name("title").Label("Title"),
				comp.Column().Name("due_date").Label("Due Date").Type("date").Sortable(true),
			).
			FilterDefaultVisible(false).FilterTogglable(true).
			Filter(
				comp.Form().Title("").Body(
					comp.Switch().Name("is_completed").Label("Done"),
					comp.InputText().Name("title").Label("Keywords"),
					comp.Button().Icon("fa fa-search").Label("Search").Primary(true).ActionType("submit"),
					comp.Button().Icon("fa fa-refresh").Label("Reset").ActionType("reset"),
				).Actions()).
			OnEvent(
				am.Schema{
					"rowClick": am.Schema{
						"actions": []comp.MEventAction{
							comp.EventAction().ActionType("drawer").Drawer(
								detail(api.Todo+"?id=${event.data.item.id}", "patch:"+api.Todo+"?id=${event.data.item.id}"),
							),
						},
					},
				},
			).
			HeaderToolbar(
				"filter-toggler",
				"bulkActions",
				"pagination",
			).
			FooterToolbar().
			BulkActions(
				comp.Button().Icon("fa fa-trash").Level("danger").Label("Delete").ActionType("ajax").ConfirmText("Delete the tasks?").Api("delete:"+api.Todo+"?ids=${ids}").ReloadWindow(),
			),
		comp.Button().Icon("fa fa-plus").Primary(true).ClassName("w-full").Label("Add").ActionType("drawer").Drawer(detail("", "")),
	)
}

func detail(getApi, editApi string) any {
	isCreate := getApi == ""

	form := comp.Form().Mode("normal").AutoFocus(true).WrapWithPanel(false).Body(
		comp.Group().Body(
			comp.InputText().Name("title").Label("Title"),
			comp.InputDatetime().Name("due_date").Label("Due Date").Value("+1days").DisplayFormat("YYYY-MM-DD").ValueFormat("YYYY-MM-DDTHH:mm:ssZ"),
		),
		comp.Switch().Name("is_completed").Option("Done").Disabled(isCreate),
		comp.Markdown().Options(am.Schema{"html": true}).Name("detail"),
		comp.Editor().Name("detail").Language("markdown").Size("xl").Value("${detail}").AllowFullscreen(false).Options(am.Schema{
			"overviewRulerBorder": false,
		}),
	).Rules(
		comp.Rule().Rule("data.title && data.detail").Message("Both title and content can't be empty"),
	).Reload("todos")

	if isCreate {
		form.SubmitTo(&model.Todo{}, func(a any) error {
			if util.ReadOnly() {
				return errors.New(api.ReadonlyMsg)
			}
			return db.AddTodo(a.(*model.Todo))
		})
	} else {
		form.InitApi(getApi)
	}
	if editApi != "" {
		form.Api(editApi)
	}

	return comp.Drawer().Size("lg").ShowCloseButton(false).CloseOnOutside(true).Body(form)
}
