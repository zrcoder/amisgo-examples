package main

import (
	"github.com/zrcoder/amisgo/comp"
)

var page = comp.Page().
	Toolbar(
		comp.Form().StaticLabelClassName("").
			PanelClassName("mb-0").
			Body(
				comp.Select().
					Label("区域").
					Name("businessLineId").
					SelectFirst(true).
					Mode("inline").
					Options(
						"北京",
						"上海",
					).
					CheckAll(false),
				comp.InputDateRange().
					Label("时间范围").
					Name("dateRange").
					Inline(true).
					Value("-1month,+0month").
					InputFormat("YYYY-MM-DD").
					CloseOnSelect(true).
					Clearable(false),
			).
			Mode("inline").
			Target("mainPage").
			SubmitOnChange(true).
			SubmitOnInit(true),
	).
	Body(
		comp.Grid().
			Columns(
				comp.Panel().
					ClassName("h-full").
					Body(
						comp.Tabs().
							Tabs(
								comp.Tab().
									Title("消费趋势").
									Tab(
										comp.Chart().
											Config(trendChartCOnfig),
									),
								comp.Tab().
									Title("账户余额").
									Tab("0.00"),
							),
					),
				comp.Panel().
					ClassName("h-full").
					Body("Test"),
			),
		comp.CrudTable().
			ClassName("m-t-sm").
			SyncLocation(false).
			Api(itemsRouter). // This demonstrates the use of Api;
			// You can also use the FetchData method, which is much simpler than Api and does not require ServeApi, as shown below:
			// FetchData(func() (any, error) {
			// 	return items, nil
			// }).
			Columns(
				comp.Column().Name("id").Label("ID"),
				comp.Column().Name("engine").Label("Rendering engine"),
				comp.Column().Name("browser").Label("Browser"),
				comp.Column().Name("platform").Label("Platform(s)"),
				comp.Column().Name("version").Label("Engine version"),
				comp.Column().Name("grade").Label("CSS grade"),
			),
	)
