package main

import (
	"net/http"
	"strconv"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/fairy/static"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
)

const imageCnt = 8

func main() {
	index := comp.Page().Body(
		comp.Tabs().TabsMode("sidebar").Swipeable(true).Tabs(
			comp.Tab().Title("Hello 1").ID("1").Tab(carousel(getOptions(1, 8))),
			comp.Tab().Title("Hello 2").ID("2").Tab(carousel(getOptions(2, 13))),
			comp.Tab().Title("Hello 3").ID("3").Tab(carousel(getOptions(3, 16))),
			comp.Tab().Title("Hello 4").ID("4").Tab(carousel(getOptions(4, 18))),
			comp.Tab().Title("Hello 5").ID("5").Tab(carousel(getOptions(5, 30))),
		),
	)
	app := amisgo.New(conf.WithTheme(conf.ThemeDark)).
		StaticFS("/static/", http.FS(static.FS)).
		Mount("/", index)

	panic(app.Run(":8080"))
}

func carousel(opts []any) any {
	return comp.Carousel().
		Height("800").
		ThumbMode("cover").
		Multiple(comp.Schema{"count": 2}).
		ControlsTheme("dark").
		Auto(false).
		AlwaysShowArrow(true).
		Animation("slide").
		Options(
			opts...,
		)
}

func getOptions(id, cnt int) []any {
	res := make([]any, cnt)
	prefix := "/static/" + strconv.Itoa(id) + "/"
	for i := range res {
		res[i] = comp.Schema{"image": prefix + strconv.Itoa(i+1) + ".webp"}
	}
	return res
}
