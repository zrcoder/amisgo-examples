package selectcom

import (
	"fmt"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
	"github.com/zrcoder/amisgo/schema"
)

type Chapter struct {
	Label    string  `json:"label"`
	Children []Level `json:"children"`
	Data     any     `json:"-"`
}

type Level struct {
	Label string `json:"label"`
	Value string `json:"value"` // auto matained, dont't chage or use
	Data  any    `json:"-"`
}

func Demos(a *amisgo.App) []app.Demo {
	const levelSelectID = "levelSelect"
	return []app.Demo{
		{Name: "Select", View: a.Page().Body(
			a.Form().Mode("inline").WrapWithPanel(false).SubmitOnChange(true).Submit(
				func(s schema.Schema) error {
					value := s.Get(levelSelectID).(string)
					fmt.Println("select form submited:", value)
					return nil
				},
			).Body(
				a.Select().Name(levelSelectID).Label("LEVEL").SelectMode("chained").LabelClassName("text-xl font-bold").Value("test").Options(
					Chapter{
						Label: "1",
						Children: []Level{
							{Label: "1-1", Value: "1<::>1"},
							{Label: "1-2", Value: "1<::>2"},
						},
					},
					Chapter{
						Label: "2",
						Children: []Level{
							{Label: "2-1", Value: "2<::>1"},
						},
					},
				),
			),
		)},
	}
}
