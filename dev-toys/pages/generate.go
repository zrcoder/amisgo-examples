package pages

import (
	"encoding/base64"

	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"

	"github.com/zrcoder/amisgo"
	am "github.com/zrcoder/amisgo/model"
)

func JsonGraph(app *amisgo.App) any {
	return comp.EditorImg(app, "json", sampleJson, func(s any) (any, error) {
		src := s.(string)
		buf, err := util.Json2Svg([]byte(src))
		return regularSvgData(buf.Bytes(), err)
	})
}

func Qrcode(app *amisgo.App) any {
	return app.Form().AutoFocus(true).ColumnCount(2).WrapWithPanel(false).Body(
		app.Wrapper().ClassName("w-2/4").Body(
			comp.Editor(app, comp.EditorCfg{Lang: "text", Name: "editor"}),
		),
		app.Flex().ClassName("w-2/4").AlignItems("center").Items(
			app.QRCode().Name("qrcode").Value("${editor}").CodeSize(256).Level("M").BackgroundColor("white").ForegroundColor("#333"),
		),
	)
}

func Hash(app *amisgo.App) any {
	return app.Form().AutoFocus(true).WrapWithPanel(false).Body(
		app.Editor().Language("text").Name("editor").AllowFullscreen(false).Options(am.Schema{"fontSize": 14}),
		app.Flex().ClassName("w-full").Items(
			app.Button().Label("▼").TransformMultiple(
				func(d am.Schema) (am.Schema, error) {
					return util.Hash([]byte(d.Get("editor").(string)))
				},
				"editor",
			),
		),
		app.InputText().Name("md5").Label("MD5").Disabled(true),
		app.InputText().Name("sha1").Label("SHA1").Disabled(true),
		app.InputText().Name("sha256").Label("SHA256").Disabled(true),
		app.InputText().Name("sha512").Label("SHA512").Disabled(true),
	)
}

func Json2struct(app *amisgo.App) any {
	return comp.DualEditor(app, jsonCfg, comp.EditorCfg{Lang: "go"}, "", func(input any) (output any, err error) {
		return util.Json2Struct([]byte(input.(string)))
	}, nil)
}

func regularSvgData(input []byte, err error) (string, error) {
	if err != nil {
		return "", err
	}
	res := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(input))
	return res, nil
}
