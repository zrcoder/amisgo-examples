package main

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo/comp"
)

func main() {
	count := 0
	page := comp.Page().Body(
		comp.Form().Body(
			comp.Log().Source(logApiPath).Placeholder(""),
		).Go(func(m comp.Data) error {
			Logf("submit button clicked: %d\n", count)
			count++
			return nil
		}),
	)
	ag := amisgo.New().Mount("/", page)
	panic(ag.Run(":80"))
}
