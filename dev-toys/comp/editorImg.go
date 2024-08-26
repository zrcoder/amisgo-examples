package comp

import "github.com/zrcoder/amisgo/comp"

func EditorImg(lang string, transfor func(any) any) any {
	return comp.Form().Title("").ColumnCount(2).Body(
		Editor(&EditorCfg{Lang: lang}),
		comp.Wrapper().Style(comp.Schema{"width": "50%"}).Body(
			comp.Image().Static(true).ImageMode("original"),
		),
	).Actions(
		comp.Action().Level("primary").Icon("fa fa-arrow-right").Transform("editor", "img.src", "Done", transfor),
	)
}
