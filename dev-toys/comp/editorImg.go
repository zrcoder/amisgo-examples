package comp

import "github.com/zrcoder/amisgo/comp"

func EditorImg(lang string, transfor func(any) (any, error)) any {
	return comp.Form().Title("Json viewer").ColumnCount(3).WrapWithPanel(false).Body(
		comp.Wrapper().Style(comp.Schema{"width": "50%"}).Body(
			Editor(EditorCfg{Lang: lang, Name: "editor"}),
		),
		comp.ButtonGroup().Vertical(true).Buttons(
			comp.Action().Icon("fa fa-arrow-right").Transform("editor", "img", "Done", transfor),
		),
		comp.Flex().Style(comp.Schema{"width": "40%"}).AlignItems("center").Items(
			comp.Wrapper().Style(comp.Schema{"width": "90%"}).Body(
				comp.Image().Name("img").ImageMode("original").InnerClassName("no-border"),
			),
		),
	).Actions()
}
