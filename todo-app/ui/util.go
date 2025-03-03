package ui

import "github.com/zrcoder/amisgo/comp"

func (u *UI) page(body ...any) comp.Page {
	return u.Page().UseMobileUI(false).ClassName("p-8").Title(
		u.Button().Icon("fa fa-tasks").Label("${i18n.name}").ActionType("link").Link("/").ClassName("font-bold bg-none border-none"),
	).Toolbar(u.DrawerAction().Label("${i18n.settings}").Icon("fa fa-cog").Drawer(u.setting())).Body(body...)
}
