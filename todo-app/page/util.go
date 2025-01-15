package page

import "github.com/zrcoder/amisgo/comp"

func page(toolbar any, body ...any) any {
	toolbar = comp.InputGroup().Body(comp.ThemeButtonGroupSelect(), comp.Wrapper(), toolbar)
	return comp.Page().UseMobileUI(false).ClassName("p-8").Title(
		comp.Button().Label("TODOs").ActionType("link").Link("/").ClassName("font-bold bg-none border-none"),
	).Toolbar(toolbar).Body(body...)
}
