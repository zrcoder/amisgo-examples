package divider

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
)

func Demos(a *amisgo.App) []app.Item {
	return []app.Item{
		{Name: "Base", View: a.Divider()},
		{Name: "Rotate", View: a.Wrapper().Body(
			a.Tpl().Tpl("A"),
			a.Divider().Rotate(45),
			a.Tpl().Tpl("B"),
		)},
		{Name: "Style", View: a.Wrapper().Body(
			a.Divider().LineStyle("solid").Color("orange"),
			a.Divider().LineStyle("dashed").Color("green"),
		)},
		{Name: "With Title", View: a.Wrapper().Body(
			a.Divider().Title("title left").TitlePosition("left"),
			a.Divider().Title("title center").TitlePosition("center"),
			a.Divider().Title("title right").TitlePosition("right"),
		)},
	}
}
