package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	app := amisgo.New()
	index := app.Page().Title("amisgo").Body(
		app.Form().
			Api("https://xxx/api/saveForm").
			Body(
				app.InputText().Label("姓名").Name("name"),
				app.InputEmail().Label("邮箱").Name("email"),
			),
	)
	app.Mount("/", index)

	panic(app.Run(":8080"))
}
