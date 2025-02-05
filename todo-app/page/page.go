package page

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/todo-app/api"

	am "github.com/zrcoder/amisgo/model"
)

func Index(app *amisgo.App) any {
	return page(
		app,
		app.Form().WrapWithPanel(false).Api(api.Logout).Mode("inline").Body(
			app.Button().Label("Logout").ActionType("submit").Redirect("/login"),
		),
		app.Crud().Name("todos").Api(api.Todos).SyncLocation(false).
			Columns(
				app.Column().Name("is_completed").Label("Done").Type("status"),
				app.Column().Name("title").Label("Title"),
				app.Column().Name("due_date").Label("Due Date").Type("date").Sortable(true),
			).
			FilterDefaultVisible(false).FilterTogglable(true).
			Filter(
				app.Form().Title("").Body(
					app.Switch().Name("is_completed").Label("Done"),
					app.InputText().Name("title").Label("Keywords"),
					app.Button().Icon("fa fa-search").Label("Search").Primary(true).ActionType("submit"),
					app.Button().Icon("fa fa-refresh").Label("Reset").ActionType("reset"),
				).Actions()).
			OnEvent(
				am.Schema{
					"rowClick": am.Schema{
						"actions": []am.EventAction{
							app.EventAction().ActionType("drawer").Drawer(
								detail(app, api.Todo+"?id=${event.data.item.id}", "patch:"+api.Todo+"?id=${event.data.item.id}"),
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
				app.Button().Icon("fa fa-trash").Level("danger").Label("Delete").ActionType("ajax").ConfirmText("Delete the tasks?").Api("delete:"+api.Todo+"?ids=${ids}").ReloadWindow(),
			),
		app.Button().Icon("fa fa-plus").Primary(true).ClassName("w-full").Label("Add").ActionType("drawer").Drawer(detail(app, "", "")),
	)
}

func detail(app *amisgo.App, getApi, editApi string) any {
	isCreate := getApi == ""

	form := app.Form().Mode("normal").AutoFocus(true).WrapWithPanel(false).Body(
		app.Group().Body(
			app.InputText().Name("title").Label("Title"),
			app.InputDatetime().Name("due_date").Label("Due Date").Value("+1days").DisplayFormat("YYYY-MM-DD").ValueFormat("YYYY-MM-DDTHH:mm:ssZ"),
		),
		app.Switch().Name("is_completed").Option("Done").Disabled(isCreate),
		app.Markdown().Options(am.Schema{"html": true}).Name("detail"),
		app.Editor().Name("detail").Language("markdown").Size("xl").Value("${detail}").AllowFullscreen(false).Options(am.Schema{
			"overviewRulerBorder": false,
		}),
	).Rules(
		app.Rule().Rule("data.title && data.detail").Message("Both title and content can't be empty"),
	).Reload("todos")

	if isCreate {
		form.Api(api.Todo)
	} else {
		form.InitApi(getApi)
	}
	if editApi != "" {
		form.Api(editApi)
	}

	return app.Drawer().Size("lg").ShowCloseButton(false).CloseOnOutside(true).Body(form)
}
