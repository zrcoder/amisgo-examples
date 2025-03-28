package main

import (
	_ "embed"
	"fmt"

	"github.com/zrcoder/amisgo"
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
			conf.Theme{Value: conf.ThemeDark, Label: "☾"},
			conf.Theme{Value: conf.ThemeCxd, Label: "○"},
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
					"level":     game.Level,
					"info":      info,
					"infoClass": infoClass,
				}, nil
			}).Body(
				app.Flex().Items(app.Tpl().Tpl("${info}").ClassName("${infoClass}")),
				app.Wrapper(),
				app.Flex().Items(
					app.Amis().Name("game-ui"),
				),
				app.Wrapper(),
				app.Flex().Items(),
				app.Flex().Items(
					app.Tpl().Tpl("Click any bottle or press the bottle key to select a bottle.").ClassName("text-xl text-gray-500"),
				),
				app.Divider(),
				app.Flex().Items(
					app.Form().Mode("inline").WrapWithPanel(false).Submit(
						func(s schema.Schema) error {
							game.SetLevel(s.Get("level").(string))
							return nil
						},
					).Body(
						app.ButtonGroupSelect().Label("LEVEL").Name("level").Options(
							app.Option().Label("Easy").Value("easy"),
							app.Option().Label("Medium").Value("medium"),
							app.Option().Label("Hard").Value("hard"),
						),
						app.Button().Label("SET <Ctrl+R>").Primary(true).Icon("fa fa-refresh").ActionType("submit").Reload("game").HotKey("ctrl+r"),
					),
				),
			),
		)
	app.Mount("/", index)

	fmt.Println("Server started at http://localhost:3000")
	panic(app.Run(":3000"))
}
