package comp

import "github.com/zrcoder/amisgo/comp"

type EditorCfg struct {
	Name     string
	Lang     string
	Label    string
	Value    string
	ReadOnly bool
}

func Editor(e EditorCfg) any {
	if e.Lang == "" {
		e.Lang = "text"
	}
	return comp.Editor().
		Name(e.Name).
		Language(e.Lang).
		Label(e.Label).
		Value(e.Value).
		Disabled(e.ReadOnly).
		Options(comp.Schema{"fontSize": 14}).
		Size("xxl").
		AllowFullscreen(false)
}
