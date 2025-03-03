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
		u.Crud().Name("todos").Api(api.Todos).SyncLocation(false).
			Columns(
				u.Column(u.Status()).Name("is_completed").Label("${i18n.todos.done}"),
				u.Column().Name("title").Label("${i18n.todos.title}").Sortable(true),
				u.Column(u.Date()).Name("due_date").Label("${i18n.todos.dueDate}").Sortable(true),
			).
			FilterDefaultVisible(false).FilterTogglable(true).
			Filter(
				u.Form().Title("").Body(
					u.Switch().Name("is_completed").Label("${i18n.todos.done}"),
					u.InputText().Name("title").Label("${i18n.todos.keywords}"),
					u.SubmitAction().Icon("fa fa-search").Label("${i18n.todos.search}").Primary(true),
					u.Action().Icon("fa fa-refresh").Label("${i18n.todos.reset}").ActionType("reset"),
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
				u.Action().Icon("fa fa-trash").Level("danger").Label("${i18n.todos.delete}").ActionType("ajax").ConfirmText("Delete the tasks?").Api("delete:"+api.Todo+"?ids=${ids}").ReloadWindow(),
			),
		u.DrawerAction().Icon("fa fa-plus").Primary(true).ClassName("w-full").Label("${i18n.todos.add}").Drawer(u.detail("", "")),
	)
}

func (u *UI) detail(getApi, editApi string) comp.Drawer {
	isCreate := getApi == ""

	form := u.Form().Mode("normal").AutoFocus(true).WrapWithPanel(false).Body(
		u.Group().Body(
			u.InputText().Name("title").Label("${i18n.todos.title}").Required(true),
			u.InputDatetime().Name("due_date").Label("${i18n.todos.dueDate}").Value("+1days").DisplayFormat("YYYY-MM-DD").ValueFormat("YYYY-MM-DDTHH:mm:ssZ"),
		),
		u.Switch().Name("is_completed").Label("${i18n.todos.done}").Disabled(isCreate),
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
