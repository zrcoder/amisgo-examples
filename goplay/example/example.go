package example

import (
	"embed"
	"strings"

	"github.com/zrcoder/amisgo/schema"
)

//go:embed code/*
var fs embed.FS

const defaultCodeKey = "Hello, World!"

var (
	keys = []string{
		defaultCodeKey,
		// "Conway's Game of Life",
		"Fibonacci Closure",
		"Peano Inteagers",
		"Concurrent pi",
		"Concurrent Prime Sieve",
		"Peg Solitaire Solver",
		"Tree Comparison",
		"Clear Screen",
		"Http Server",
		"Display Image",
		// "Multiple Files",
		"Sleep",
		// "Test Function",
		"Generic index",
	}
	replacer = strings.NewReplacer(
		" ", "-",
		"'", "_",
		",", "",
		"!", "",
	)
)

func Get() (options []any, defaultCode string, err error) {
	for _, key := range keys {
		dir := replacer.Replace(key)
		data, err := fs.ReadFile("code/" + dir + "/main.go")
		if err != nil {
			return nil, "", err
		}
		val := string(data)
		if key == defaultCodeKey {
			defaultCode = val
		}
		options = append(options, schema.Schema{"label": key, "value": val})
	}
	return
}
