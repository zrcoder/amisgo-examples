package page

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo-examples/todo-app/util"
)

func Login(app *amisgo.App) any {
	demoInput := ""
	if util.ReadOnly() {
		demoInput = "amisgo"
	}
	return page(
		app,
		"",
		app.Flex().ClassName("pt-20 bg-red").Items(
			app.Wrapper().ClassName("w-96").Body(
				app.Form().Title("").Api(api.Login).AutoFocus(true).Redirect("/todos").Body(
					app.InputText().Name("name").Label("Name").Value(demoInput).Required(true),
					app.InputPassword().Name("password").Label("Password").Value(demoInput).Required(true),
				).Actions(
					app.Button().Label("sign up").ActionType("link").Link("/register"),
					app.Button().Primary(true).Label("login").ActionType("submit"),
				),
			),
		),
	)
}

func Register(app *amisgo.App) any {
	return page(
		app,
		"",
		app.Flex().ClassName("pt-20").Items(
			app.Wrapper().ClassName("w-96").Body(
				app.Form().Title("").Api(api.Register).AutoFocus(true).Redirect("/login").Body(
					app.InputText().Name("name").Label("Name").Required(true),
					app.InputPassword().Name("password").Label("Password").Required(true),
				).Actions(
					app.Button().ActionType("submit").Label("sign up"),
				),
			),
		),
	)
}
