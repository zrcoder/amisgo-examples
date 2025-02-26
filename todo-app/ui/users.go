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
		u.Flex().ClassName("pt-20 bg-red").Items(
			u.Wrapper().ClassName("w-96").Body(
				u.Form().Title("").Api(api.Login).AutoFocus(true).Redirect("/todos").Body(
					u.InputText().Name("name").Label("${i18n.user.name}").Value(demoInput).Required(true),
					u.InputPassword().Name("password").Label("${i18n.user.password}").Value(demoInput).Required(true),
				).Actions(
					u.Action().Label("${i18n.user.signUp}").ActionType("link").Link("/register"),
					u.SubmitAction().Primary(true).Label("${i18n.user.login}"),
				),
			),
		),
	)
}

func (u *UI) Register() comp.Page {
	return u.page(
		u.Flex().ClassName("pt-20").Items(
			u.Wrapper().ClassName("w-96").Body(
				u.Form().Title("").Api(api.Register).AutoFocus(true).Redirect("/login").Body(
					u.InputText().Name("name").Label("${i18n.user.name}").Required(true),
					u.InputPassword().Name("password").Label("${i18n.user.password}").Required(true),
				).Actions(
					u.SubmitAction().Label("${i18n.user.signUp}"),
				),
			),
		),
	)
}
