package ui

import (
	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo/comp"
)

func (u UI) setting() comp.Drawer {
	return u.Drawer().Actions().Body(
		u.Form().Mode("inline").WrapWithPanel(false).Api(api.Login).Body(
			u.Group().Body(
				u.ThemeButtonGroupSelect().Label("${i18n.theme}"),
				u.LocaleButtonGroupSelect().Label("${i18n.lang}"),
			),
			u.SubmitAction().Label("${i18n.logout}").Icon("fa fa-sign-out").Redirect("/login").Api(api.Logout),
		),
	)
}
