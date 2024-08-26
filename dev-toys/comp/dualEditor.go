package comp

import "github.com/zrcoder/amisgo/comp"

func DualEditor(left, right *EditorCfg, action func(any) any) any {
	right.readOnly = true
	return comp.Form().Title("").ColumnCount(2).Body(Editor(left), Editor(right)).Actions(
		comp.Action().Label("Go").Icon("fa fa-arrow-right").Transform(left.Name, right.Name, "done", action),
	)
}
