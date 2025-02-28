package main

import (
	_ "embed"
	"encoding/json"

	"github.com/zrcoder/amisgo"
)

//go:embed pages/index.json
var index json.RawMessage

func main() {
	app := amisgo.New()
	app.Mount("/", index)
	app.Run(":8080")
}
