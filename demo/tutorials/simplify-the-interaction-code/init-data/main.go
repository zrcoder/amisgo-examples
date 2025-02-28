package main

import (
	"fmt"
	"time"

	"github.com/zrcoder/amisgo"
)

func main() {
	app := amisgo.New()
	index := app.Page().Body("Now: ${date}").InitData(getDate)
	app.Mount("/", index)
	app.Run(":8888")
}

func getDate() (any, error) {
	y, m, d := time.Now().Date()
	mm := time.Now().UnixNano()
	return map[string]string{"date": fmt.Sprintf("%d-%d-%d %d", y, m, d, mm)}, nil
}
