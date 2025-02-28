package main

import (
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

func main() {
	app := amisgo.New(
		conf.WithLocales(
			conf.Locale{Value: conf.LocaleZhCN, Label: "汉"},
			conf.Locale{Value: conf.LocaleEnUS, Label: "En"},
		),
	)
	index := app.Page().Title("amisgo").Body(
		app.LocaleButtonGroupSelect(),
	)
	app.Mount("/", index)

	fmt.Println("Please visit http://localhost:8080")
	panic(app.Run(":8080"))
}
