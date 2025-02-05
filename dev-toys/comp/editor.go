package comp

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/model"
)

type EditorCfg struct {
	Name     string
	Lang     string
	Label    string
	Value    string
	ReadOnly bool
}

func Editor(app *amisgo.App, e EditorCfg) any {
	if e.Lang == "" {
		e.Lang = "text"
	}
	return app.Editor().
		Name(e.Name).
		Language(e.Lang).
		Label(e.Label).
		Value(e.Value).
		Disabled(e.ReadOnly).
		Options(model.Schema{"fontSize": 14}).
		Size("xxl").
		AllowFullscreen(false)
}
