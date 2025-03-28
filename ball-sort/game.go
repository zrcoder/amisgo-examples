package main

import (
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

	BucketBallCount   = 4
	EmptyBuckets      = 2
	BallSize          = 40
	bucketClass       = "relative w-18 h-auto mx-2"
	bucketButtonClass = "absolute inset-0 h-full rounded-xl bucket-button"
)

var (
	colors = []string{
		"red", "green", "blue", "yellow", "orange", "pink", "purple",
	}
	LevelBucketsDict = map[string]int{
		LevelEasy:   5,
		LevelMedium: 6,
		LevelHard:   7,
	}
)

type Game struct {
	rd *rand.Rand
	*amisgo.App
	Level           string
	Buckets         []*Bucket
	ShiftBall       *Ball
	DoneBucketCount int
}

type Bucket struct {
	*Game
	Index int
	Balls []*Ball
}

type Ball struct {
	*Bucket
	ID          int
	Color       string
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
	n := LevelBucketsDict[g.Level]
	balls := make([]*Ball, 0, n*BucketBallCount)
	for i := range n {
		for range BucketBallCount {
			ball := &Ball{ID: i}
			ball.Nomalize()
			balls = append(balls, ball)
		}
	}
	g.rd.Shuffle(len(balls), func(i, j int) {
		balls[i], balls[j] = balls[j], balls[i]
	})
	g.Buckets = make([]*Bucket, 0, n+EmptyBuckets)
	i := 0
	for range n {
		bucket := &Bucket{Game: g}
		bucket.Balls = balls[i : i+BucketBallCount]
		for _, ball := range bucket.Balls {
			ball.Bucket = bucket
		}
		i += BucketBallCount
		g.Buckets = append(g.Buckets, bucket)
	}
	for range EmptyBuckets {
		bucket := &Bucket{
			Game:  g,
			Balls: make([]*Ball, 0, BucketBallCount),
		}
		g.Buckets = append(g.Buckets, bucket)
	}
	g.rd.Shuffle(len(g.Buckets), func(i, j int) {
		g.Buckets[i], g.Buckets[j] = g.Buckets[j], g.Buckets[i]
	})
	for i, bucket := range g.Buckets {
		bucket.Index = i
	}
	g.ShiftBall = nil
	g.DoneBucketCount = 0
}

func (g *Game) SetLevel(level string) {
	g.Level = level
	g.Reset()
}

func (g *Game) SelectBucket(i int) {
	if i < 0 || i >= len(g.Buckets) {
		return
	}
	bucket := g.Buckets[i]
	if bucket.IsDone() {
		return
	}
	if bucket.IsEmpty() {
		if g.ShiftBall == nil {
			return
		}
		bucket.Push(g.ShiftBall)
		g.ShiftBall = nil
		return
	}
	if g.ShiftBall == nil {
		g.ShiftBall = bucket.Pop()
		return
	}
	if g.ShiftBall.Bucket.Index == i {
		bucket.Push(g.ShiftBall)
		g.ShiftBall = nil
		return
	}
	if g.ShiftBall.ID != bucket.Top().ID {
		g.ShiftBall.Bucket.Push(g.ShiftBall)
		g.ShiftBall = bucket.Pop()
		return
	}
	if bucket.IsFull() {
		return
	}
	bucket.Push(g.ShiftBall)
	g.ShiftBall = nil
	g.DoneBucketCount += bucket.checkDone()
}

func (g *Game) IsDone() bool {
	return g.DoneBucketCount == LevelBucketsDict[g.Level]
}

func (b *Bucket) Pop() *Ball {
	n := len(b.Balls)
	if n == 0 {
		return nil
	}
	res := b.Balls[n-1]
	b.Balls = b.Balls[:n-1]
	return res
}

func (b *Bucket) Push(ball *Ball) {
	b.Balls = append(b.Balls, ball)
	ball.Bucket = b
}

func (b *Bucket) Top() *Ball {
	if len(b.Balls) == 0 {
		return nil
	}
	return b.Balls[len(b.Balls)-1]
}

func (b *Bucket) IsEmpty() bool {
	return len(b.Balls) == 0
}

func (b *Bucket) IsFull() bool {
	return len(b.Balls) == BucketBallCount
}

func (b *Bucket) checkDone() int {
	if !b.IsFull() {
		return 0
	}
	for i := 1; i < len(b.Balls); i++ {
		if b.Balls[i].ID != b.Balls[0].ID {
			return 0
		}
	}
	return 1
}

func (b *Bucket) IsDone() bool {
	return b.checkDone() == 1
}

func (b *Bucket) IsShiftBall() bool {
	return b.Game.ShiftBall != nil && b.Game.ShiftBall.Bucket == b
}

func (g *Game) UI() (any, error) {
	total := len(g.Buckets)
	buckets := make([]any, 0, total)
	for _, bucket := range g.Buckets {
		buckets = append(buckets, bucket.UI())
	}
	half := total / 2
	ui := g.App.Wrapper().Body(
		g.App.Flex().Justify("center").Items(
			buckets[:half]...,
		),
		g.App.Flex().Justify("center").Items(
			buckets[half:]...,
		),
	)
	return ui, nil
}

func (b *Bucket) UI() any {
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

	items := make([]any, BucketBallCount)
	for i, ball := range b.Balls {
		items[BucketBallCount-i-1] = ball.UI()
	}
	for i := BucketBallCount - len(b.Balls) - 1; i >= 0; i-- {
		items[i] = placeholderBallUI(b.App)
	}
	return b.App.Form().WrapWithPanel(false).Submit(func(s schema.Schema) error {
		b.Game.SelectBucket(b.Index)
		return nil
	}).Body(
		b.App.Wrapper().ClassName(bucketClass).Body(
			top,
		),
		b.App.Wrapper().ClassName(bucketClass).
			Body(
				items,
				b.Button().ActionType("submit").Reload("game").ClassName(bucketButtonClass).Disabled(done),
			),
	)
}

func (b *Ball) Nomalize() {
	b.Color = colors[b.ID]
}

func (b *Ball) UI() comp.Shape {
	return ballUI(b.App, b.Color)
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
