package game

import (
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

func (g *Game) makeBottleForms() {
	g.bottleForms = make([]comp.Form, len(g.colors)+EmptyBottles)
	for i := range g.bottleForms {
		g.bottleForms[i] = g.App.Form().WrapWithPanel(false).Submit(
			func(s schema.Schema) error {
				g.SelectBottle(i)
				return nil
			},
		)
	}
}

func (g *Game) makeLevelUI() {
	g.levelUI = g.App.Flex().Items(
		g.levelForm(-1), g.levelForm(1), g.App.Wrapper(), g.levelForm(0),
	)
}

func (g *Game) levelForm(delta int) comp.Form {
	var label, icon, hotkey string
	var action func()
	var primary bool
	switch delta {
	case 1:
		label = "<Ctrl-N>"
		icon = "fa fa-arrow-right"
		hotkey = "ctrl+n"
		action = g.NextLevel
	case -1:
		label = "<Ctrl-P>"
		icon = "fa fa-arrow-left"
		hotkey = "ctrl+p"
		action = g.PrevLevel
	default:
		label = "<Ctrl-R>"
		icon = "fa fa-refresh"
		hotkey = "ctrl+r"
		action = g.Reset
		primary = true
	}
	return g.App.Form().Mode("inline").WrapWithPanel(false).Submit(
		func(s schema.Schema) error {
			action()
			return nil
		}).
		Body(
			g.App.Button().Label(label).Icon(icon).Primary(primary).
				ActionType("submit").Reload(gameNameUI).HotKey(hotkey),
		)
}
