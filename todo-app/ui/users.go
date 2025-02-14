package ui

import (
	"github.com/zrcoder/amisgo-examples/todo-app/api"
	"github.com/zrcoder/amisgo-examples/todo-app/util"
	"github.com/zrcoder/amisgo/comp"
)

func (u *UI) Login() comp.Page {
	demoInput := ""
	if util.IsDemo() {
		demoInput = "amisgo"
	}
	return u.page(
		"",
		u.Flex().ClassName("pt-20 bg-red").Items(
			u.Wrapper().ClassName("w-96").Body(
				u.Form().Title("").Api(api.Login).AutoFocus(true).Redirect("/todos").Body(
					u.InputText().Name("name").Label("Name").Value(demoInput).Required(true),
					u.InputPassword().Name("password").Label("Password").Value(demoInput).Required(true),
				).Actions(
					u.Action().Label("sign up").ActionType("link").Link("/register"),
					u.SubmitAction().Primary(true).Label("login"),
				),
			),
		),
	)
}

func (u *UI) Register() comp.Page {
	return u.page(
		"",
		u.Flex().ClassName("pt-20").Items(
			u.Wrapper().ClassName("w-96").Body(
				u.Form().Title("").Api(api.Register).AutoFocus(true).Redirect("/login").Body(
					u.InputText().Name("name").Label("Name").Required(true),
					u.InputPassword().Name("password").Label("Password").Required(true),
				).Actions(
					u.SubmitAction().Label("sign up"),
				),
			),
		),
	)
}
