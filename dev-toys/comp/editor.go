package comp

import "github.com/zrcoder/amisgo/comp"

type EditorCfg struct {
	Name     string
	Langs    []string
	Lang     string
	Label    string
	readOnly bool
}

func Editor(e *EditorCfg) any {
	editor := comp.Editor().Language(e.Lang).ReadOnly(e.readOnly).Size("xxl")
	if e.Lang == "" {
		e.Lang = "text"
	}
	if len(e.Langs) == 0 {
		e.Langs = []string{e.Lang}
	}
	opts := make([]any, len(e.Langs))
	for i, v := range e.Langs {
		opts[i] = comp.Option().Label(v).Value(v)
	}
	return comp.Wrapper().Style(comp.Schema{"width": "50%"}).Body(
		comp.ButtonGroupSelect().Name("abc").Options(opts...).OptionValue(e.Lang),
		// comp.Tabs().ActiveKey(e.Lang).Swipeable(true).TabsMode("line").Tabs(tabs...),
		editor,
	)
}
