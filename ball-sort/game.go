package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

const (
	LevelEasy   = "easy"
	LevelMedium = "medium"
	LevelHard   = "hard"

	BottleBallCount   = 4
	EmptyBottles      = 2
	BallSize          = 40
	bottleClass       = "relative w-18 h-auto mx-2"
	bottleButtonClass = "absolute inset-0 h-full rounded-xl bottle-button"
)

var (
	colors = []string{
		"red", "green", "blue", "yellow", "brown", "pink", "purple",
	}
	succeedMsgs = []string{
		"Wanderful!", "Brilliant!", "Excellent!", "Fantastic!", "Awesome!",
	}
	LevelBottlesDict = map[string]int{
		LevelEasy:   5,
		LevelMedium: 6,
		LevelHard:   7,
	}
)

type Game struct {
	rd *rand.Rand
	*amisgo.App
	Level            string
	Bottles          []*Bottle
	ShiftBall        *Ball
	DoneBottlesCount int
}

type Bottle struct {
	*Game
	Index int
	Balls []*Ball
}

type Ball struct {
	*Bottle
	Type        int
	Placeholder bool
}

func NewGame(app *amisgo.App) *Game {
	res := &Game{
		Level: LevelEasy,
		App:   app,
		rd:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	res.Reset()
	return res
}

func (g *Game) Reset() {
	g.rd.Shuffle(len(colors), func(i, j int) {
		colors[i], colors[j] = colors[j], colors[i]
	})
	n := LevelBottlesDict[g.Level]
	balls := make([]*Ball, 0, n*BottleBallCount)
	for i := range n {
		for range BottleBallCount {
			balls = append(balls, &Ball{Type: i})
		}
	}
	g.rd.Shuffle(len(balls), func(i, j int) {
		balls[i], balls[j] = balls[j], balls[i]
	})
	g.Bottles = make([]*Bottle, 0, n+EmptyBottles)
	i := 0
	for range n {
		bottle := &Bottle{Game: g}
		bottle.Balls = balls[i : i+BottleBallCount]
		for _, ball := range bottle.Balls {
			ball.Bottle = bottle
		}
		i += BottleBallCount
		g.Bottles = append(g.Bottles, bottle)
	}
	for range EmptyBottles {
		bottle := &Bottle{
			Game:  g,
			Balls: make([]*Ball, 0, BottleBallCount),
		}
		g.Bottles = append(g.Bottles, bottle)
	}
	g.rd.Shuffle(len(g.Bottles), func(i, j int) {
		g.Bottles[i], g.Bottles[j] = g.Bottles[j], g.Bottles[i]
	})
	for i, bottle := range g.Bottles {
		bottle.Index = i
	}
	g.ShiftBall = nil
	g.DoneBottlesCount = 0
}

func (g *Game) SetLevel(level string) {
	g.Level = level
	g.Reset()
}

func (g *Game) SelectBottle(i int) {
	if i < 0 || i >= len(g.Bottles) {
		return
	}
	bottle := g.Bottles[i]
	if bottle.IsDone() {
		return
	}
	if bottle.IsEmpty() {
		if g.ShiftBall == nil {
			return
		}
		bottle.Push(g.ShiftBall)
		g.ShiftBall = nil
		return
	}
	if g.ShiftBall == nil {
		g.ShiftBall = bottle.Pop()
		return
	}
	if g.ShiftBall.Bottle.Index == i {
		bottle.Push(g.ShiftBall)
		g.ShiftBall = nil
		return
	}
	if g.ShiftBall.Type != bottle.Top().Type || bottle.IsFull() {
		g.ShiftBall.Bottle.Push(g.ShiftBall)
		g.ShiftBall = bottle.Pop()
		return
	}
	bottle.Push(g.ShiftBall)
	g.ShiftBall = nil
	g.DoneBottlesCount += bottle.checkDone()
}

func (g *Game) IsDone() bool {
	return g.DoneBottlesCount == LevelBottlesDict[g.Level]
}

func (g *Game) StateInfo() (info, infoClass string) {
	if game.IsDone() {
		info = succeedMsgs[game.rd.Intn(len(succeedMsgs))]
		infoClass = "text-2xl font-bold text-success"
	} else {
		info = fmt.Sprintf("Done: %d/%d", game.DoneBottlesCount, LevelBottlesDict[g.Level])
		infoClass = "text-xl font-bold text-info"
	}
	return
}

func (b *Bottle) Pop() *Ball {
	n := len(b.Balls)
	if n == 0 {
		return nil
	}
	res := b.Balls[n-1]
	b.Balls = b.Balls[:n-1]
	return res
}

func (b *Bottle) Push(ball *Ball) {
	b.Balls = append(b.Balls, ball)
	ball.Bottle = b
}

func (b *Bottle) Top() *Ball {
	if len(b.Balls) == 0 {
		return nil
	}
	return b.Balls[len(b.Balls)-1]
}

func (b *Bottle) IsEmpty() bool {
	return len(b.Balls) == 0
}

func (b *Bottle) IsFull() bool {
	return len(b.Balls) == BottleBallCount
}

func (b *Bottle) checkDone() int {
	if !b.IsFull() {
		return 0
	}
	for i := 1; i < len(b.Balls); i++ {
		if b.Balls[i].Type != b.Balls[0].Type {
			return 0
		}
	}
	return 1
}

func (b *Bottle) IsDone() bool {
	return b.checkDone() == 1
}

func (b *Bottle) IsShiftBall() bool {
	return b.Game.ShiftBall != nil && b.Game.ShiftBall.Bottle == b
}

func (g *Game) UI() (any, error) {
	total := len(g.Bottles)
	bottles := make([]any, 0, total)
	for _, bottle := range g.Bottles {
		bottles = append(bottles, bottle.UI())
	}
	ui := g.App.Flex().Items(
		bottles...,
	)
	return ui, nil
}

func (b *Bottle) UI() any {
	done := b.IsDone()
	var top any
	switch {
	case done:
		top = starUI(b.App)
	case b.IsShiftBall():
		top = b.Game.ShiftBall.UI()
	default:
		top = placeholderBallUI(b.App)
	}

	items := make([]any, BottleBallCount)
	for i, ball := range b.Balls {
		items[BottleBallCount-i-1] = ball.UI()
	}
	for i := BottleBallCount - len(b.Balls) - 1; i >= 0; i-- {
		items[i] = placeholderBallUI(b.App)
	}
	key := string(rune('A' + b.Index))
	return b.App.Form().WrapWithPanel(false).Submit(func(s schema.Schema) error {
		b.Game.SelectBottle(b.Index)
		return nil
	}).Body(
		b.App.Wrapper().ClassName(bottleClass).Body(
			top,
		),
		b.App.Wrapper().ClassName(bottleClass).
			Body(
				items,
				b.Button().HotKey(key).
					ActionType("submit").Reload("game").
					ClassName(bottleButtonClass).Disabled(done),
			),
		b.Flex().Items(
			b.Tpl().Tpl(key),
		),
	)
}

func (b *Ball) Color() string {
	return colors[b.Type]
}

func (b *Ball) UI() comp.Shape {
	return ballUI(b.App, b.Color())
}

func placeholderBallUI(app *amisgo.App) comp.Shape {
	return ballUI(app, "transparent")
}

func ballUI(app *amisgo.App, color string) comp.Shape {
	return app.Shape().ShapeType("circle").Width(BallSize).Height(BallSize).Color(color)
}

func starUI(app *amisgo.App) comp.Shape {
	return app.Shape().ShapeType("star").Width(BallSize).Height(BallSize).Color("orange")
}
