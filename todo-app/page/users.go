package page

import (
	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo-examples/todo-app/util"

	"github.com/zrcoder/amisgo/comp"
)

func Login() any {
	demoInput := ""
	if util.ReadOnly() {
		demoInput = "amisgo"
	}
	return page(
		"",
		comp.Flex().ClassName("pt-20 bg-red").Items(
			comp.Wrapper().ClassName("w-96").Body(
				comp.Form().Title("").Api(api.Login).AutoFocus(true).Redirect("/todos").Body(
					comp.InputText().Name("name").Label("Name").Value(demoInput),
					comp.InputPassword().Name("password").Label("Password").Value(demoInput),
				).Actions(
					comp.Button().Label("sign up").ActionType("link").Link("/register"),
					comp.Button().Primary(true).Label("login").ActionType("submit"),
				),
			),
		),
	)
}

func Register() any {
	return page(
		"",
		comp.Flex().ClassName("pt-20").Items(
			comp.Wrapper().ClassName("w-96").Body(
				comp.Form().Title("").Api(api.Register).AutoFocus(true).Redirect("/login").Body(
					comp.InputText().Name("name").Label("Name"),
					comp.InputPassword().Name("password").Label("Password"),
				).Actions(
					comp.Button().ActionType("submit").Label("sign up"),
				),
			),
		),
	)
}
