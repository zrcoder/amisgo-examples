package main

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

var (
	//go:embed i18n/zh-CN.json
	zhCN json.RawMessage
	//go:embed i18n/en-US.json
	enUS json.RawMessage
)

func main() {
	app := amisgo.New(
		conf.WithLocales(
			conf.Locale{Value: conf.LocaleZhCN, Label: "汉", Dict: zhCN},
			conf.Locale{Value: conf.LocaleEnUS, Label: "En", Dict: enUS},
		),
	)
	index := app.Page().Title("amisgo").Body(
		app.LocaleButtonGroupSelect(),
		app.Form().Body(
			app.InputText().Label("${i18n.index.name}").Name("name"),
			app.InputEmail().Label("${i18n.index.email}").Name("email"),
		),
	)
	app.Mount("/", index)

	fmt.Println("Please visit http://localhost:8080")
	panic(app.Run(":8080"))
}
