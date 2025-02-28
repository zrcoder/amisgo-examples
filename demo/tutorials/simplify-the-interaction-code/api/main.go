package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zrcoder/amisgo"
)

func main() {
	app := amisgo.New()
	index := app.Page().Body(
		app.Form().Api("/user").Body(
			app.InputText().Label("姓名").Name("name"),
			app.InputEmail().Label("邮箱").Name("email"),
		),
	)
	app.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		input, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		m := map[string]string{}
		json.Unmarshal(input, &m)

		name := m["name"]
		email := m["email"]
		fmt.Println(name, email)
		// 将用户信息存入数据库 ...
	})
	app.Mount("/", index)
	app.Run(":8888")
}
