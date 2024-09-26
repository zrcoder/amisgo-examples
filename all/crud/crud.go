package crud

import (
	"encoding/json"

	"github.com/zrcoder/amisgo/comp"
)

var Crud = comp.Page().Body(
	comp.Service().GetData(func() (any, error) {
		return items, nil
	}).Body(
		comp.Crud().Name("crud").Source("${rows}").Columns(
			comp.Column().Name("id").Label("ID"),
			comp.Column().Name("engine").Label("Rendering engine"),
			comp.Column().Name("browser").Label("Browser"),
			comp.Column().Name("platform").Label("Platform(s)"),
			comp.Column().Name("grade").Label("CSS grade"),
			comp.Column().Type("operation").Name("operation").Label("Operation").Buttons(
				comp.Button().Label("Detail").Level("link").ActionType("dialog").Dialog(
					comp.Dialog().Body(dialogBody(true)).Confirm(false),
				),
				comp.Button().Label("Modify").ActionType("drawer").Drawer(
					comp.Drawer().Title("Modify").Body(
						dialogBody(false, func(d comp.Data) error {
							data, err := json.Marshal(d)
							if err != nil {
								return err
							}
							var item Item
							err = json.Unmarshal(data, &item)
							if err != nil {
								return err
							}
							for i, v := range items.Rows {
								if v.ID == item.ID {
									items.Rows[i] = item
								}
							}
							return nil
						}),
					),
				),
				comp.Button().Label("Delete").Level("danger").Reload("crud").DisabledOn("this.grade === 'A'").ActionType("dialog").Dialog(
					comp.Dialog().Body(
						dialogBody(true, func(d comp.Data) error {
							id := int(d.Get("id").(float64))
							rows := make([]Item, 0, len(items.Rows))
							for _, v := range items.Rows {
								if v.ID != id {
									rows = append(rows, v)
								}
							}
							items.Rows = rows
							return nil
						})),
				),
			),
		),
		comp.Button().Label("Add").ActionType("dialog").Level("primary").ClassName("m-b-sm").Dialog(
			comp.Dialog().Body(
				comp.Form().Body(
					comp.InputText().Label("Engine").Name("engine"),
					comp.InputText().Label("Platform").Name("platform"),
				).Go(func(d comp.Data) error {
					engine := d.Get("engine").(string)
					platform := d.Get("platform").(string)
					items.Rows = append(items.Rows, Item{Engine: engine, Platform: platform})
					return nil
				}),
			),
		),
	),
)

func dialogBody(disableInput bool, action ...func(d comp.Data) error) any {
	res := comp.Form().Body(
		comp.InputText().Name("id").Label("ID").Disabled(true),
		comp.InputText().Name("engine").Disabled(disableInput).Label("engine"),
		comp.InputText().Name("browser").Disabled(disableInput).Label("browser"),
		comp.InputText().Name("platform").Disabled(disableInput).Label("platform"),
		comp.InputText().Name("version").Disabled(disableInput).Label("version"),
		comp.Control().Label("grade").Disabled(disableInput).Body(
			comp.Tag().Label("${grade}").DisplayMode("normal").Color("active"),
		))

	if action != nil {
		res.Go(action[0])
	}
	return res
}

type Items struct {
	Rows []Item `json:"rows"`
}

type Item struct {
	Engine   string `json:"engine"`
	Browser  string `json:"browser"`
	Platform string `json:"platform"`
	Version  string `json:"version"`
	Grade    string `json:"grade"`
	ID       int    `json:"id"`
	Children []Item `json:"children"`
}

var items = Items{
	Rows: []Item{
		{
			Engine:   "Trident",
			Browser:  "IE 4.0",
			Platform: "Win 95+",
			Version:  "4",
			Grade:    "X",
			ID:       1,
		},
		{
			Engine:   "Trident",
			Browser:  "IE 4.0",
			Platform: "Win 95+",
			Version:  "4",
			Grade:    "Y",
			ID:       1001,
		},
		{
			Engine:   "Trident",
			Browser:  "IE 4.0",
			Platform: "Win 95+",
			Version:  "4",
			Grade:    "Z",
			ID:       1002,
		},
		{
			Engine:   "Trident",
			Browser:  "Safari 6.0",
			Platform: "Win 95+",
			Version:  "4",
			Grade:    "A",
			ID:       2,
		},
		{
			Engine:   "Trident",
			Browser:  "IE 4.0",
			Platform: "Win 95+",
			Version:  "4",
			Grade:    "Y",
			ID:       2001,
		},
		{
			Engine:   "Trident",
			Browser:  "IE 4.0",
			Platform: "Win 95+",
			Version:  "4",
			Grade:    "Z",
			ID:       2002,
		},
	},
}
