package main

import (
	"encoding/base64"

	"amisgo-examples/dev-toys/comp"

	qr "github.com/skip2/go-qrcode"
	"github.com/zrcoder/ttoy/pkg/generator"
)

var jsonViewer = comp.EditorImg("json", func(s any) (any, error) {
	src := s.(string)
	buf, err := generator.Json2Svg([]byte(src))
	return regularSvgData(buf.Bytes(), err)
})

var qrcode = comp.EditorImg("text", func(input any) (any, error) {
	q, err := qr.New(input.(string), qr.Medium)
	if err != nil {
		return nil, err
	}
	return regularPngData(q.PNG(-1))
})

func regularSvgData(input []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	res := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(input))
	return res, nil
}

func regularPngData(input []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	res := "data:image/png;base64," + base64.StdEncoding.EncodeToString([]byte(input))
	return res, nil
}