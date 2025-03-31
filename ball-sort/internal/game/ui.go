package game

import (
	"fmt"

	"github.com/zrcoder/amisgo/comp"
)

const (
	gameNameUI = "game"
	ballSize   = 40
)

func (g *Game) Index() comp.Page {
	return g.App.Page().
		Title(g.App.Tpl().Tpl("Ball Sort Puzzle").ClassName("text-2xl font-bold")).
		Toolbar(g.App.ThemeButtonGroupSelect()).
		Body(
			g.App.Service().Name(gameNameUI).GetData(func() (any, error) {
				ui, err := g.Main()
				if err != nil {
					return nil, err
				}
				return map[string]any{"index": ui}, nil
			}).Body(
				g.App.Amis().Name("index"),
			),
		)
}

func (g *Game) Main() (any, error) {
	var state []any
	if g.IsDone() {
		info := succeedMsgs[g.rd.Intn(len(succeedMsgs))]
		infoClass := "text-2xl font-bold text-success"
		state = []any{g.App.Tpl().Tpl(info).ClassName(infoClass)}
	} else {
		levelInfo := fmt.Sprintf("LEVEL: [ %s ]", g.CurrentLevel().Name)
		info := fmt.Sprintf("Done: %d/%d", g.DoneBottlesCount, g.CurrentLevel().Bottles)
		infoClass := "text-xl font-bold text-info"
		state = []any{
			g.App.Tpl().Tpl(levelInfo).ClassName("text-xl"),
			g.App.Wrapper(),
			g.App.Tpl().Tpl(info).ClassName(infoClass),
		}
	}
	bottles := make([]any, len(g.Bottles))
	for i, bottle := range g.Bottles {
		bottles[i] = bottle.UI()
	}
	return g.Wrapper().Body(
		g.App.Flex().Items(state...),
		g.App.Wrapper(),
		g.App.Flex().Items(bottles...),
		g.App.Wrapper(),
		g.App.Flex().Items(
			g.App.Tpl().
				Tpl("Click any bottle or press the bottle key to select a bottle.").
				ClassName("text-xl text-gray-500"),
		),
		g.App.Divider(),
		g.levelUI,
	), nil
}

func (b *Bottle) UI() any {
	done := b.IsDone()
	var top any
	switch {
	case done:
		top = b.Game.starUI()
	case b.IsShiftBall():
		top = b.Game.ShiftBall.UI()
	default:
		top = b.Game.placeholderBallUI()
	}

	items := make([]any, BottleBallCount)
	for i, ball := range b.Balls {
		items[BottleBallCount-i-1] = ball.UI()
	}
	for i := BottleBallCount - len(b.Balls) - 1; i >= 0; i-- {
		items[i] = b.Game.placeholderBallUI()
	}

	key := string(rune('A' + b.Index))
	return b.Game.bottleForms[b.Index].Body(
		b.App.Wrapper().ClassName("mx-2").Body(top),
		b.App.Wrapper().ClassName("relative w-18 h-auto mx-2").Body(
			items,
			b.Button().HotKey(key).ActionType("submit").Reload(gameNameUI).
				ClassName("absolute inset-0 h-full rounded-xl bottle-button").Disabled(done),
		),
		b.Flex().Items(b.Tpl().Tpl(key)),
	)
}

func (b *Ball) UI() comp.Shape {
	return b.Game.shape("circle", b.Game.colors[b.Type])
}

func (g *Game) placeholderBallUI() comp.Shape {
	return g.shape("circle", "transparent")
}

func (g *Game) starUI() comp.Shape {
	return g.shape("star", "orange")
}

func (g *Game) shape(shape, color string) comp.Shape {
	return g.App.Shape().ShapeType(shape).Width(ballSize).Height(ballSize).Color(color)
}
