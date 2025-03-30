package main

import (
	"math/rand"
	"time"

	"github.com/zrcoder/amisgo"
)

const (
	LevelEasy   = "EASY"
	LevelMedium = "MEDIUM"
	LevelHard   = "HARD"
	LevelExpert = "EXPERT"

	BottleBallCount = 4
	EmptyBottles    = 2
)

var succeedMsgs = []string{
	"Wanderful!", "Brilliant!", "Excellent!", "Fantastic!", "Awesome!",
}

var levels = []Level{
	{Name: LevelEasy, Bottles: 5},
	{Name: LevelMedium, Bottles: 6},
	{Name: LevelHard, Bottles: 7},
	{Name: LevelExpert, Bottles: 8},
}

type Game struct {
	rd *rand.Rand
	*amisgo.App
	levelIndex       int
	Bottles          []*Bottle
	ShiftBall        *Ball
	DoneBottlesCount int
	colors           []string
}

type Level struct {
	Name    string
	Bottles int
}

type Bottle struct {
	*Game
	Index int
	Balls []*Ball
}

type Ball struct {
	*Bottle
	Type int
}

func NewGame(app *amisgo.App) *Game {
	res := &Game{
		App: app,
		rd:  rand.New(rand.NewSource(time.Now().UnixNano())),
		colors: []string{
			"red", "green", "blue", "yellow", "brown", "pink", "purple", "orange",
		},
	}
	res.Reset()
	return res
}

func (g *Game) Reset() {
	g.rd.Shuffle(len(g.colors), func(i, j int) {
		g.colors[i], g.colors[j] = g.colors[j], g.colors[i]
	})
	n := levels[g.levelIndex].Bottles
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
	for _, bottle := range g.Bottles {
		if bottle.IsDone() {
			g.DoneBottlesCount++
		}
	}
}

func (g *Game) PrevLevel() {
	if g.levelIndex > 0 {
		g.levelIndex--
		g.Reset()
	}
}

func (g *Game) NextLevel() {
	if g.levelIndex < len(levels)-1 {
		g.levelIndex++
		g.Reset()
	}
}

func (g *Game) CurrentLevel() Level {
	return levels[g.levelIndex]
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
	if g.ShiftBall.Bottle.Index != i &&
		(g.ShiftBall.Type != bottle.Top().Type || bottle.IsFull()) {
		g.ShiftBall.Bottle.Push(g.ShiftBall)
		g.ShiftBall = bottle.Pop()
		return
	}
	bottle.Push(g.ShiftBall)
	g.ShiftBall = nil
	g.DoneBottlesCount += bottle.checkDone()
}

func (g *Game) IsDone() bool {
	return g.DoneBottlesCount == levels[g.levelIndex].Bottles
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
