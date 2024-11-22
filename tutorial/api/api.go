package main

type Item struct {
	ID       string `json:"id"`
	Engine   string `json:"engine"`
	Browser  string `json:"browser"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
	Grade    string `json:"grade"`
}

var items = []Item{
	{Engine: "e1", Browser: "chrome", Platform: "windows", Version: "1.0"},
	{Browser: "safri", Platform: "macOS", Version: "2.0"},
}

const (
	trendChartCOnfig = `{
	"title": {
	  "text": "消费趋势"
	},
	"tooltip": {},
	"xAxis": {
	  "type": "category",
	  "boundaryGap": false,
	  "data": [
		"一月",
		"二月",
		"三月",
		"四月",
		"五月",
		"六月"
	  ]
	},
	"yAxis": {},
	"series": [
	  {
		"name": "销量",
		"type": "line",
		"areaStyle": {
		  "color": {
			"type": "linear",
			"x": 0,
			"y": 0,
			"x2": 0,
			"y2": 1,
			"colorStops": [
			  {
				"offset": 0,
				"color": "rgba(84, 112, 197, 1)"
			  },
			  {
				"offset": 1,
				"color": "rgba(84, 112, 197, 0)"
			  }
			],
			"global": false
		  }
		},
		"data": [
		  5,
		  20,
		  36,
		  10,
		  10,
		  20
		]
	  }
	]
  }`
)
