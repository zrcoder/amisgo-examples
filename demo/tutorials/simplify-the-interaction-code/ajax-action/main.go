package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/schema"
)

func main() {
	app := amisgo.New()
	index := app.Page().Body(
		app.Form().WrapWithPanel(false).Body(
			app.InputText().Name("input"),
			app.InputText().Name("output").ReadOnly(true),
			app.Action().Label("Greet").Level("primary").ActionType("ajax").Api(
				app.Api().Url("/convert").Data(schema.Schema{"input": "${input}"}).Set(
					"resp",
					schema.Schema{
						"200": schema.Schema{
							"then": app.EventAction().ActionType("setValue").
								Args(app.EventActionArgs().Value("${resp}")),
						},
					},
				),
			),
		),
	)
	app.Mount("/", index)
	app.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		input, _ := io.ReadAll(r.Body)
		defer r.Body.Close()
		m := map[string]string{}
		json.Unmarshal(input, &m)
		output := "hello " + m["input"]
		resp := schema.SuccessResponse("", schema.Schema{"output": output}) // 这里的 key 值必须是第二个编辑器的 name
		w.Write(resp.Json())
	})

	app.Run(":8888")
}
