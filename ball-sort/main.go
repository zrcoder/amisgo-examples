package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
	"github.com/zrcoder/amisgo/schema"
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
				info, infoClass := game.StateInfo()
				return map[string]any{
					"game-ui":   ui,
					"level":     strings.ToTitle(game.CurrentLevel().Name),
					"info":      info,
					"infoClass": infoClass,
				}, nil
			}).Body(
				app.Flex().Items(
					app.Tpl().Tpl("LEVEL: [ $level ]").ClassName("text-xl"),
					app.Wrapper(),
					app.Tpl().Tpl("${info}").ClassName("${infoClass}"),
				),
				app.Wrapper(),
				app.Flex().Items(app.Amis().Name("game-ui")),
				app.Wrapper(),
				app.Flex().Items(
					app.Tpl().Tpl("Click any bottle or press the bottle key to select a bottle.").ClassName("text-xl text-gray-500"),
				),
				app.Divider(),
				app.Flex().Items(levelForm(-1), levelForm(1), levelForm(0)),
			),
		)
	app.Mount("/", index)

	fmt.Println("Server started at http://localhost:3000")
	panic(app.Run(":3000"))
}

func levelForm(delta int) comp.Form {
	var label, icon, hotkey string
	var action func()
	var primary bool
	switch delta {
	case 1:
		label = "<Ctrl-N>"
		icon = "fa fa-arrow-right"
		hotkey = "ctrl+n"
		action = game.NextLevel
	case -1:
		label = "<Ctrl-P>"
		icon = "fa fa-arrow-left"
		hotkey = "ctrl+p"
		action = game.PrevLevel
	default:
		label = "<Ctrl-R>"
		icon = "fa fa-refresh"
		hotkey = "ctrl+r"
		action = game.Reset
		primary = true
	}
	return app.Form().Mode("inline").WrapWithPanel(false).Submit(
		func(s schema.Schema) error {
			action()
			return nil
		}).
		Body(
			app.Button().Label(label).Icon(icon).Primary(primary).
				ActionType("submit").Reload("game").HotKey(hotkey),
		)
}
