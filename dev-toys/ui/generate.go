package ui

import (
	"encoding/base64"

	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"

	"github.com/zrcoder/amisgo"
	am "github.com/zrcoder/amisgo/model"
)

type Generators struct {
	*amisgo.App
}

func NewGenerators(app *amisgo.App) *Generators { return &Generators{App: app} }

func (g *Generators) JsonGraph() any {
	return comp.EditorImg(g.App, "json", sampleJson, func(s any) (any, error) {
		src := s.(string)
		buf, err := util.Json2Svg([]byte(src))
		return regularSvgData(buf.Bytes(), err)
	})
}

func (g *Generators) JsonViewer() any {
	return comp.EditorJson(g.App, sampleJson)
}

func (g *Generators) Qrcode() any {
	return g.Form().AutoFocus(true).ColumnCount(2).WrapWithPanel(false).Body(
		g.Wrapper().ClassName("w-1/2").Body(
			comp.Editor(g.App, comp.EditorCfg{Lang: "text", Name: "editor"}),
		),
		g.Wrapper().ClassName("w-1/2").Body(
			g.Flex().ClassName("h-full").AlignItems("center").Items(
				g.QRCode().Name("qrcode").Value("${editor}").CodeSize(256).Level("M").BackgroundColor("white").ForegroundColor("#333"),
			),
		),
	)
}

func (g *Generators) Hash() any {
	return g.Form().AutoFocus(true).WrapWithPanel(false).Body(
		g.Editor().Language("text").Name("editor").AllowFullscreen(false).Options(am.Schema{"fontSize": 14}),
		g.Flex().ClassName("w-full").Items(
			g.Button().Label("▼").TransformMultiple(
				func(d am.Schema) (am.Schema, error) {
					return util.Hash([]byte(d.Get("editor").(string)))
				},
				"editor",
			),
		),
		g.InputText().Name("md5").Label("MD5").Disabled(true),
		g.InputText().Name("sha1").Label("SHA1").Disabled(true),
		g.InputText().Name("sha256").Label("SHA256").Disabled(true),
		g.InputText().Name("sha512").Label("SHA512").Disabled(true),
	)
}

func (g *Generators) Json2struct() any {
	return comp.DualEditor(g.App, jsonCfg, comp.EditorCfg{Lang: "go"}, "", func(input any) (output any, err error) {
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
