package main

import (
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
	"github.com/zrcoder/amisgo/schema"
)

var app *amisgo.App

func main() {
	app = amisgo.New()
	index := app.App().Pages(
		app.PageItem().Url("/").Redirect("/base"),
		app.PageItem().Children(
			app.PageItem().Label("Base").Url("base").Schema(base()),
			app.PageItem().Label("Cell Style").Url("cell").Schema(cell()),
			app.PageItem().Label("Col Style").Url("col").Schema(col()),
			app.PageItem().Label("Caption").Url("caption").Schema(caption()),
			app.PageItem().Label("Condition").Url("condition").Schema(condition()),
			app.PageItem().Label("Visible on").Url("visibleon").Schema(visibleOn()),
			app.PageItem().Label("Layout").Url("layout").Schema(layout()),
		),
	)
	app.Mount("/", index)
	fmt.Println("Serving on http://localhost:8080")
	panic(app.Run(":8080"))
}

func base() comp.TableView {
	return app.TableView().Trs(
		app.Tr().Background("#F7F7F7").Tds(
			app.Td().Body(app.Tpl().Tpl("地区")),
			app.Td().Body(app.Tpl().Tpl("城市")),
			app.Td().Body(app.Tpl().Tpl("销量")),
		),
		app.Tr().Tds(
			app.Td().Body(app.Tpl().Tpl("华北")).Rowspan(2),
			app.Td().Body(app.Tpl().Tpl("北京")),
			app.Td().Body(app.Tpl().Tpl("20")),
		),
		app.Tr().Tds(
			app.Td().Body(app.Tpl().Tpl("天津")),
			app.Td().Body(app.Tpl().Tpl("10")),
		),
	)
}

func cell() comp.TableView {
	return app.TableView().Trs(
		app.Tr().Background("#F7F7F7").Tds(
			app.Td().Body(app.Tpl().Tpl("地区")),
			app.Td().Body(app.Tpl().Tpl("城市")),
			app.Td().Body(app.Tpl().Tpl("销量")),
		),
		app.Tr().Tds(
			app.Td().Body(app.Tpl().Tpl("华北")).Style(schema.Schema{"borderBottomWidth": 0, "borderLeftWidth": 0}),
			app.Td().Body(app.Tpl().Tpl("北京")),
			app.Td().Body(app.Tpl().Tpl("20")).Style((schema.Schema{"borderBottomWidth": 0, "borderRightWidth": 0})),
		),
	)
}

func col() comp.TableView {
	return app.TableView().
		Cols(
			app.Tcol().Span(2),
			app.Tcol().Style(schema.Schema{"background": "#00F7F7"}),
		).
		Trs(
			app.Tr().Background("#F7F7F7").Tds(
				app.Td().Body(app.Tpl().Tpl("地区")),
				app.Td().Body(app.Tpl().Tpl("城市")),
				app.Td().Body(app.Tpl().Tpl("销量")),
			),
			app.Tr().Tds(
				app.Td().Body(app.Tpl().Tpl("华北")).Rowspan(2),
				app.Td().Body(app.Tpl().Tpl("北京")),
				app.Td().Body(app.Tpl().Tpl("20")),
			),
			app.Tr().Tds(
				app.Td().Body(app.Tpl().Tpl("天津")),
				app.Td().Body(app.Tpl().Tpl("10")),
			),
		)
}

func caption() comp.TableView {
	return app.TableView().
		Caption("标题").
		// CaptionSide("bottom").
		Cols(
			app.Tcol().Span(2),
			app.Tcol().Style(schema.Schema{"background": "#00F7F7"}),
		).
		Trs(
			app.Tr().Background("#F7F7F7").Tds(
				app.Td().Body(app.Tpl().Tpl("地区")),
				app.Td().Body(app.Tpl().Tpl("城市")),
				app.Td().Body(app.Tpl().Tpl("销量")),
			),
			app.Tr().Tds(
				app.Td().Body(app.Tpl().Tpl("华北")).Rowspan(2),
				app.Td().Body(app.Tpl().Tpl("北京")),
				app.Td().Body(app.Tpl().Tpl("20")),
			),
			app.Tr().Tds(
				app.Td().Body(app.Tpl().Tpl("天津")),
				app.Td().Body(app.Tpl().Tpl("10")),
			),
		)
}

func condition() comp.Service {
	return app.Service().
		Data(
			schema.Schema{"score": 40},
		).
		Body(
			app.TableView().Trs(
				app.Tr().Tds(
					app.Td().Background("${score > 50 ? 'green' : 'red'}").Body("score ${score} > 50 ?"),
					app.Td().Background("${score<100 ? 'green': 'red'}").Body("score ${score} < 100?"),
				),
			),
		)
}

func visibleOn() comp.Page {
	return app.Page().Body(
		app.Switch().Label("显示第一行").Name("row1").Value(true),
		app.Switch().Label("显示北京单元格").Name("beijing").Value(true),
		app.TableView().Trs(
			app.Tr().Background("#f7f7f7").VisibleOn("row1").
				Tds(app.Td().Body("地区"), app.Td().Body("城市"), app.Td().Body("销量")),
			app.Tr().Tds(app.Td().Body("华北"), app.Td().VisibleOn("beijing").Body("北京")),
		),
	)
}

func layout() comp.TableView {
	return app.TableView().Border(true).Trs(
		app.Tr().Background("orange").Tds(app.Td().Colspan(4).Align("center").Body("Header")),
		app.Tr().Tds(
			app.Td().Background("gray").Width(200).Body("Side").Rowspan(2),
			app.Td().Background("pink").Align("center").Body("Top Right").Colspan(3),
		),
		app.Tr().Background("lightgreen").Height(200).Tds(
			app.Td().Align("center").Body("C1"),
			app.Td().Align("center").Body("C2"),
			app.Td().Align("center").Body("C3"),
		),
		app.Tr().Background("lightblue").Tds(app.Td().Colspan(4).Align("center").Body("Footer")),
	)
}
