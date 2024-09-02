package comp

import "github.com/zrcoder/amisgo/comp"

func DualEditor(left, right EditorCfg, title string, action, reverseAction func(input any) (output any, err error)) any {
	left.Name = "input"
	right.Name = "output"
	left.ReadOnly = true
	right.ReadOnly = true
	actions := make([]any, 0, 2)
	if action != nil {
		left.ReadOnly = false
		actions = append(actions,
			comp.Action().Icon("fa fa-arrow-right").Transform(left.Name, right.Name, "Done", action))
	}
	if reverseAction != nil {
		right.ReadOnly = false
		actions = append(actions,
			comp.Action().Icon("fa fa-arrow-left").Transform(right.Name, left.Name, "Done", reverseAction))
	}
	return comp.Form().Title(title).ColumnCount(3).AutoFocus(true).WrapWithPanel(false).Body(
		Editor(left),
		comp.ButtonGroup().Vertical(true).Buttons(
			actions...,
		),
		Editor(right),
	).Actions()
}
