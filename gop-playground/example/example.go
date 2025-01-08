package example

import (
	"embed"
	"strings"

	"github.com/zrcoder/amisgo/model"
)

//go:embed *
var FS embed.FS

const defaultCodeKey = "Hello, Go+"

var keys = []string{
	defaultCodeKey,
	"Go+ Basic",
	"Range",
	"Rational",
	"Slice literal",
	"List-Map comprehension",
	"Error Wrap",
}

func Get() (options []any, defaultCode string, err error) {
	for _, key := range keys {
		data, err := FS.ReadFile(key + ".gop")
		if err != nil {
			return nil, "", err
		}
		key = strings.ReplaceAll(key, "-", "/")
		val := string(data)
		if key == defaultCodeKey {
			defaultCode = val
		}
		options = append(options, model.Schema{"label": key, "value": val})
	}
	return
}
