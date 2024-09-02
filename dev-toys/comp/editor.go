package comp

import "github.com/zrcoder/amisgo/comp"

type EditorCfg struct {
	Name     string
	Lang     string
	Label    string
	ReadOnly bool
}

func Editor(e EditorCfg) any {
	if e.Lang == "" {
		e.Lang = "text"
	}
	return comp.Editor().
		AllowFullscreen(false).
		Name(e.Name).
		Language(e.Lang).
		Disabled(e.ReadOnly).
		Size("xxl").
		Label(e.Label)
}
