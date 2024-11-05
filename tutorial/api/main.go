package main

import (
	"github.com/zrcoder/amisgo"
)

func main() {
	ServeApi()

	ag := amisgo.New().Register("/", page)
	panic(ag.Run(":80"))
}
