package ui

func (u *UI) page(toolbar any, body ...any) any {
	toolbar = u.InputGroup().Body(u.ThemeButtonGroupSelect(), u.Wrapper(), toolbar)
	return u.Page().UseMobileUI(false).ClassName("p-8").Title(
		u.Button().Label("TODOs").ActionType("link").Link("/").ClassName("font-bold bg-none border-none"),
	).Toolbar(toolbar).Body(body...)
}
