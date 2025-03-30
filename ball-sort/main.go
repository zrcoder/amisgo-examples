package main

import (
	_ "embed"
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/conf"
)

var (
	app  *amisgo.App
	game *Game
)

//go:embed main.css
var customCSS string

func main() {
	app = amisgo.New(
		conf.WithThemes(
			conf.Theme{Value: conf.ThemeCxd, Label: "○"},
			conf.Theme{Value: conf.ThemeDark, Label: "☽"},
		),
		conf.WithCustomCSS(customCSS),
	)
	game = NewGame(app)

	index := app.Page().
		Title(app.Tpl().Tpl("Ball Sort Puzzle").ClassName("text-2xl font-bold")).
		Toolbar(app.ThemeButtonGroupSelect()).
		Body(
			app.Service().Name("game").GetData(func() (any, error) {
				ui, err := game.UI()
				if err != nil {
					return nil, err
				}
				return map[string]any{"game-ui": ui}, nil
			}).Body(
				app.Flex().Items(app.Amis().Name("game-ui")),
			),
		)
	app.Mount("/", index)

	fmt.Println("Server started at http://localhost:3000")
	panic(app.Run(":3000"))
}
