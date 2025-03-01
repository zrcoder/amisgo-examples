package app

import (
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

var (
	app  *amisgo.App
	tabs []any
)

func init() {
	app = amisgo.New()
}

func Register(title string, demo func(*amisgo.App) map[string]any) {
	tabs = append(tabs, tab(title, demo(app)))
}

func tab(title string, dict map[string]any) comp.Tab {
	return app.Tab().Title(title).Tab(view(dict))
}

func view(tabDic map[string]any) comp.Tabs {
	tabs := make([]any, 0, len(tabDic))
	for title, schema := range tabDic {
		tabs = append(tabs, app.Tab().Title(title).Tab(schema))
	}
	return app.Tabs().Tabs(tabs...)
}

func Run() {
	app.Mount("/", app.Tabs().TabsMode("vertical").Tabs(
		tabs...,
	))
	fmt.Println("Serving on http://localhost:8080")
	panic(app.Run(":8080"))
}
