package ui

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/todo-app/api"

	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

type UI struct {
	*amisgo.App
}

func New(app *amisgo.App) *UI {
	return &UI{App: app}
}

func (u *UI) Index() comp.Page {
	return u.page(
		u.Form().WrapWithPanel(false).Api(api.Logout).Mode("inline").Body(
			u.SubmitAction().Label("Logout").Redirect("/login").Api(api.Logout),
		),
		u.Crud().Name("todos").Api(api.Todos).SyncLocation(false).
			Columns(
				u.Column().Name("is_completed").Label("Done").Type("status"),
				u.Column().Name("title").Label("Title"),
				u.Column().Name("due_date").Label("Due Date").Type("date").Sortable(true),
			).
			FilterDefaultVisible(false).FilterTogglable(true).
			Filter(
				u.Form().Title("").Body(
					u.Switch().Name("is_completed").Label("Done"),
					u.InputText().Name("title").Label("Keywords"),
					u.SubmitAction().Icon("fa fa-search").Label("Search").Primary(true),
					u.Action().Icon("fa fa-refresh").Label("Reset").ActionType("reset"),
				).Actions()).
			OnEvent(
				u.Event().RowClick(
					u.EventActions(
						u.EventActionDrawer(
							u.detail(api.Todo+"?id=${event.data.item.id}", "patch:"+api.Todo+"?id=${event.data.item.id}"),
						),
					),
				),
			).
			HeaderToolbar(
				"filter-toggler",
				"bulkActions",
				"pagination",
			).
			FooterToolbar().
			BulkActions(
				u.Action().Icon("fa fa-trash").Level("danger").Label("Delete").ActionType("ajax").ConfirmText("Delete the tasks?").Api("delete:"+api.Todo+"?ids=${ids}").ReloadWindow(),
			),
		u.DrawerAction().Icon("fa fa-plus").Primary(true).ClassName("w-full").Label("Add").Drawer(u.detail("", "")),
	)
}

func (u *UI) detail(getApi, editApi string) comp.Drawer {
	isCreate := getApi == ""

	form := u.Form().Mode("normal").AutoFocus(true).WrapWithPanel(false).Body(
		u.Group().Body(
			u.InputText().Name("title").Label("Title"),
			u.InputDatetime().Name("due_date").Label("Due Date").Value("+1days").DisplayFormat("YYYY-MM-DD").ValueFormat("YYYY-MM-DDTHH:mm:ssZ"),
		),
		u.Switch().Name("is_completed").Option("Done").Disabled(isCreate),
		u.Markdown().Options(schema.Schema{"html": true}).Name("detail"),
		u.Editor().Name("detail").Language("markdown").Size("xl").Value("${detail}").AllowFullscreen(false).Options(schema.Schema{
			"overviewRulerBorder": false,
		}),
	).Rules(
		u.Rule().Rule("data.title && data.detail").Message("Both title and content can't be empty"),
	).Reload("todos")

	if isCreate {
		form.Api(api.Todo)
	} else {
		form.InitApi(getApi)
	}
	if editApi != "" {
		form.Api(editApi)
	}

	return u.Drawer().Size("lg").ShowCloseButton(false).CloseOnOutside(true).Body(form)
}
