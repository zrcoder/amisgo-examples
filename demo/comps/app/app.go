package app

import (
	"fmt"
	"net/http"

	sdk "gitee.com/rdor/amis-sdk/v6"
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/conf"
)

var (
	app  *amisgo.App
	tabs []any
)

func init() {
	app = amisgo.New(conf.WithLocalSdk(http.FS(sdk.FS)))
	app = amisgo.New()
}

func Register(title string, demo func(*amisgo.App) []Demo) {
	tabs = append(tabs, tab(title, demo(app)))
}

func tab(title string, items []Demo) comp.Tab {
	return app.Tab().Title(title).Tab(view(items))
}

func view(items []Demo) comp.Tabs {
	tabs := make([]any, 0, len(items))
	for _, item := range items {
		tabs = append(tabs, app.Tab().Title(item.Name).Tab(item.View))
	}
	return app.Tabs().TabsMode("line").Tabs(tabs...)
}

func Run() {
	app.Mount("/", app.Page().Body(app.Tabs().TabsMode("vertical").Tabs(tabs...)))
	fmt.Println("Serving on http://localhost:8080")
	panic(app.Run(":8080"))
}
