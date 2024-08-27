package main

import (
	"encoding/base64"

	"amisgo-examples/dev-toys/comp"
	"amisgo-examples/dev-toys/util"

	ac "github.com/zrcoder/amisgo/comp"
)

var jsonGraph = comp.EditorImg("json", func(s any) (any, error) {
	src := s.(string)
	buf, err := util.Json2Svg([]byte(src))
	return regularSvgData(buf.Bytes(), err)
})

var qrcode = ac.Form().Title("").ColumnCount(2).WrapWithPanel(false).Body(
	ac.Wrapper().Style(ac.Schema{"width": "50%"}).Body(
		comp.Editor(comp.EditorCfg{Lang: "text", Name: "editor"}),
	),
	ac.Flex().Style(ac.Schema{"width": "50%"}).AlignItems("center").Items(
		ac.QRCode().Name("qrcode").Value("${editor}").CodeSize(256).Level("M"),
	),
).Actions()

var json2struct = comp.DualEditor(jsonCfg, comp.EditorCfg{Lang: "go"}, "", func(input any) (output any, err error) {
	return util.Json2Struct([]byte(input.(string)))
}, nil)

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
