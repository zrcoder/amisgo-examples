package main

import (
	"encoding/json"

	"github.com/zrcoder/amisgo"
)

func main() {
	const amisJSON = `{
		"type": "page",
		"title": "Hello",
		"body": "World!"
	}`

	app := amisgo.New()
	app.Mount("/", json.RawMessage(amisJSON))
	app.Run(":8080")
}
