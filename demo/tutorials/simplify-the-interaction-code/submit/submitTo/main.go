package main

import (
	"fmt"
	"net/http"

	sdk "gitee.com/rdor/amis-sdk/v6"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

type Info struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	app := amisgo.New(conf.WithLocalSdk(http.FS(sdk.FS)))

	info := &Info{}
	index := app.Page().Body(
		app.Form().Api("/user").Body(
			app.InputText().Label("姓名").Name("name"),
			app.InputEmail().Label("邮箱").Name("email"),
		).SubmitTo(info, func() error {
			fmt.Println(info)
			return nil
		}),
	)
	app.Mount("/", index)
	app.Run(":8888")
}
