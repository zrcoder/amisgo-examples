package flex

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

func Demos(a *amisgo.App) []app.Demo {
	return []app.Demo{
		{Name: "Base", View: base(a)},
		{Name: "Justify", View: a.Wrapper().Body(
			"center",
			base(a).Justify("center"),
			"flex-start",
			base(a).Justify("flex-start"),
			"flex-end",
			base(a).Justify("flex-end"),
			"space-around",
			base(a).Justify("space-around"),
			base(a).Justify("space-between"),
			"space-between",
			base(a).Justify("space-evenly"),
			"space-evenly",
		)},
		{Name: "Align Items", View: a.Wrapper().Body(
			"center",
			base(a).ClassName("bg-black").Style(schema.Schema{"height": 100}).AlignItems("center"),
			"flex-start",
			base(a).ClassName("bg-black").Style(schema.Schema{"height": 100}).AlignItems("flex-start"),
			"flex-end",
			base(a).ClassName("bg-black").Style(schema.Schema{"height": 100}).AlignItems("flex-end"),
		)},
		{Name: "Direction", View: a.Wrapper().Body(
			"row",
			base(a).Direction("row"),
			"column",
			base(a).Direction("column"),
		)},
		{Name: "Mobie", View: base(a).Mobile(
			schema.Schema{
				"direction": "column",
			},
		)},
	}
}

func base(a *amisgo.App) comp.Flex {
	return a.Flex().Items(items(a)...)
}

func items(a *amisgo.App) []any {
	return []any{
		tpl(a, "red"),
		tpl(a, "green"),
		tpl(a, "blue"),
	}
}

func tpl(a *amisgo.App, color string) comp.Tpl {
	return a.Tpl().Style(schema.Schema{
		"backgroundColor": color,
		"width":           200,
		"height":          50,
		"margin":          5,
	})
}
