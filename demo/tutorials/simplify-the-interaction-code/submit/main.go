package main

import (
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/schema"
)

func main() {
	app := amisgo.New()
	index := app.Page().Body(
		app.Form().Api("/user").Body(
			app.InputText().Label("姓名").Name("name"),
			app.InputEmail().Label("邮箱").Name("email"),
		).Submit(
			func(s schema.Schema) error {
				name := s.Get("name").(string)
				email := s.Get("email").(string)
				fmt.Println(name, email)
				// 将用户信息存入数据库 ...
				return nil
			},
		),
	)
	app.Mount("/", index)
	app.Run(":8888")
}
