package pages

import (
	"encoding/base64"

	"dtoy/comp"
	"dtoy/util"

	ac "github.com/zrcoder/amisgo/comp"
)

var (
	JsonGraph = comp.EditorImg("json", sampleJson, func(s any) (any, error) {
		src := s.(string)
		buf, err := util.Json2Svg([]byte(src))
		return regularSvgData(buf.Bytes(), err)
	})
	Qrcode = ac.Form().AutoFocus(true).ColumnCount(2).WrapWithPanel(false).Body(
		ac.Wrapper().ClassName("w-2/4").Body(
			comp.Editor(comp.EditorCfg{Lang: "text", Name: "editor"}),
		),
		ac.Flex().ClassName("w-2/4").AlignItems("center").Items(
			ac.QRCode().Name("qrcode").Value("${editor}").CodeSize(256).Level("M").BackgroundColor("white").ForegroundColor("#333"),
		),
	)
	Hash = ac.Form().AutoFocus(true).WrapWithPanel(false).Body(
		ac.Editor().Language("text").Name("editor").AllowFullscreen(false).Options(ac.Schema{"fontSize": 14}),
		ac.Flex().ClassName("w-full").Items(
			ac.Button().Icon("fa fa-arrow-down").TransformMultiple("editor", "done", func(input any) (any, error) {
				return util.Hash([]byte(input.(string)))
			}),
		),
		ac.InputText().Name("md5").Label("MD5").Disabled(true),
		ac.InputText().Name("sha1").Label("SHA1").Disabled(true),
		ac.InputText().Name("sha256").Label("SHA256").Disabled(true),
		ac.InputText().Name("sha512").Label("SHA512").Disabled(true),
	)
	Json2struct = comp.DualEditor(jsonCfg, comp.EditorCfg{Lang: "go"}, "", func(input any) (output any, err error) {
		return util.Json2Struct([]byte(input.(string)))
	}, nil)
)

func regularSvgData(input []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	res := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(input))
	return res, nil
}
