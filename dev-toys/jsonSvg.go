package main

import (
	"encoding/base64"

	"amisgo-examples/dev-toys/comp"

	"github.com/zrcoder/ttoy/pkg/generator"
)

var jsonViewer = comp.EditorImg("json", func(s any) any {
	src := s.(string)
	buf, err := generator.Json2Svg([]byte(src))
	res, err := regularSvgData(buf.Bytes(), err)
	if err != nil {
		panic(err)
	}
	return res
})

func regularSvgData(input []byte, err error) (string, error) {
	return "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(input)), err
}
