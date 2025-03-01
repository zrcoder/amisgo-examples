package collapse

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
	"github.com/zrcoder/amisgo/comp"
)

func Demos(a *amisgo.App) []app.Item {
	return []app.Item{
		{Name: "Base", View: base(a)},
		{Name: "Accordion", View: base(a).Accordion(true)},
		{Name: "Expand Icon", View: base(a).ExpandIcon(a.Icon().Icon("caret-right"))},
		{Name: "Expand Icon Position", View: base(a).ExpandIconPosition("right")},
		{Name: "Embed", View: embed(a)},
		{Name: "Disabled", View: disabled(a)},
		{Name: "Hide Arrow", View: hideArrow(a)},
		{Name: "FieldSet Style in Form", View: fieldSetStyle(a)},
	}
}

func base(app *amisgo.App) comp.CollapseGroup {
	return app.CollapseGroup().ActiveKey(1).Body(simpleBody(app)...)
}

func simpleBody(app *amisgo.App) []any {
	return []any{
		app.Collapse().Key("1").Header("Title 1").Body("Body 1"),
		app.Collapse().Key("2").Header("Title 2").Body("Body 2"),
		app.Collapse().Key("3").Header("Title 3").Body("Body 3"),
	}
}

func embed(app *amisgo.App) comp.CollapseGroup {
	return app.CollapseGroup().ActiveKey(1).Body(
		app.Collapse().Key("1").Header("Title 1").Body(
			"Body 1",
			app.CollapseGroup().ActiveKey(1).Body(
				app.Collapse().Header("Embed Title").Key("1").Body("Embed Content"),
			),
		),
		app.Collapse().Key("2").Header("Title 2").Body("Body 2"),
	)
}

func disabled(app *amisgo.App) comp.CollapseGroup {
	return app.CollapseGroup().ActiveKey(1).Body(
		app.Collapse().Key("1").Header("Title 1").Body("Body 1").Disabled(true),
		app.Collapse().Key("2").Header("Title 2").Body("Body 2"),
	)
}

func hideArrow(app *amisgo.App) comp.CollapseGroup {
	return app.CollapseGroup().ActiveKey(1).Body(
		app.Collapse().Key("1").Header("Title 1").Body("Body 1").ShowArrow(false),
		app.Collapse().Key("2").Header("Title 2").Body("Body 2"),
	)
}

func fieldSetStyle(app *amisgo.App) comp.Form {
	return app.Form().Body(
		app.Switch().Value(true).Name("fieldSetStyle").Label("FieldSet style"),
		base(app).EnableFieldSetStyle("${fieldSetStyle}"),
	)
}
