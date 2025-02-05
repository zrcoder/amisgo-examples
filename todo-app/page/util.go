package page

import (
	"github.com/zrcoder/amisgo"
)

func page(app *amisgo.App, toolbar any, body ...any) any {
	toolbar = app.InputGroup().Body(app.ThemeButtonGroupSelect(), app.Wrapper(), toolbar)
	return app.Page().UseMobileUI(false).ClassName("p-8").Title(
		app.Button().Label("TODOs").ActionType("link").Link("/").ClassName("font-bold bg-none border-none"),
	).Toolbar(toolbar).Body(body...)
}
