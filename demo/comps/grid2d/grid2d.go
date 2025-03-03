package grid2d

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
)

func Demos(a *amisgo.App) []app.Demo {
	return []app.Demo{
		{Name: "Base", View: a.Grid2D().Grids(
			a.GridItem(a.Tpl().Tpl("[grid-1] x:1 y:1 h:1 w:6")).X(1).Y(1).H(1).W(6).GridClassName("bg-green-300"),
			a.GridItem(a.Tpl().Tpl("[grid-2] x:7 y:1 h:1 w:6")).X(7).Y(1).H(1).W(6).GridClassName("bg-blue-300"),
			a.GridItem(a.Tpl().Tpl("[grid-3] x:1 y:2 h:2 w:4")).X(1).Y(2).H(2).W(4).GridClassName("bg-red-300"),
			a.GridItem(a.Tpl().Tpl("[grid-4] x:5 y:2 h:1 w:8")).X(5).Y(2).H(1).W(8).GridClassName("bg-purple-300"),
		)},
		{Name: "Size", View: a.Grid2D().Cols(3).Grids(
			a.GridItem(a.Tpl().Tpl("1")).X(1).Y(1).H(1).W(1).GridClassName("bg-green-300").Width(100),
			a.GridItem(a.Tpl().Tpl("2")).X(2).Y(1).H(1).W(1).GridClassName("bg-blue-300").Height(100),
			a.GridItem(a.Tpl().Tpl("3")).X(3).Y(1).H(1).W(1).GridClassName("bg-red-300").Width(100),
			a.GridItem(a.Tpl().Tpl("4")).X(2).Y(2).H(1).W(1).GridClassName("bg-purple-300"),
		)},
		{Name: "Auto Height", View: a.Grid2D().Grids(
			a.GridItem(a.Tpl().Tpl("1</br>2</br>3</br>4</br>5</br>6</br>")).X(2).Y(1).H(1).W(1).GridClassName("bg-blue-300").Height("auto"),
		)},
		{Name: "Gap", View: a.Grid2D().Gap(10).GapRow(5).Grids(
			a.GridItem(a.Tpl().Tpl("[grid-1] x:1 y:1 h:1 w:6")).X(1).Y(1).H(1).W(6).GridClassName("bg-green-300").Width(100),
			a.GridItem(a.Tpl().Tpl("[grid-2] x:7 y:1 h:1 w:6")).X(7).Y(1).H(1).W(6).GridClassName("bg-blue-300").Height(100),
			a.GridItem(a.Tpl().Tpl("[grid-3] x:1 y:2 h:2 w:4")).X(1).Y(2).H(2).W(4).GridClassName("bg-red-300").Width(100),
			a.GridItem(a.Tpl().Tpl("[grid-4] x:5 y:2 h:1 w:8")).X(5).Y(2).H(1).W(8).GridClassName("bg-purple-300"),
		)},
		{Name: "Align", View: a.Grid2D().Cols(3).Grids(
			a.GridItem(a.Tpl().Tpl("[grid-1] x:1 y:1 h:1 w:1")).X(1).Y(1).H(1).W(1).GridClassName("bg-green-300"),
			a.GridItem(a.Tpl().Tpl("[grid-3] x:2 y:1 h:1 w:1 align: center, valign: middle")).X(2).Y(1).H(1).W(1).GridClassName("bg-red-300").Align("center").Valign("middle"),
			a.GridItem(a.Tpl().Tpl("[grid-4] x:3 y:1 h:1 w:1")).X(3).Y(1).H(1).W(1).GridClassName("bg-purple-300"),
		)},
	}
}
