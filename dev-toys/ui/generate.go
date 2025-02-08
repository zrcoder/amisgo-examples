package ui

import (
	"encoding/base64"

	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"

	ac "github.com/zrcoder/amisgo/comp"
	am "github.com/zrcoder/amisgo/model"
)

func (u *UI) JsonGraph() ac.Form {
	return u.EditorImg("json", sampleJson, func(s any) (any, error) {
		src := s.(string)
		buf, err := util.Json2Svg([]byte(src))
		return regularSvgData(buf.Bytes(), err)
	})
}

func (u *UI) JsonViewer() ac.Form {
	return u.EditorJson(sampleJson)
}

func (u *UI) Qrcode() ac.Form {
	return u.EditorQrCoder()
}

func (u *UI) Hash() ac.Form {
	return u.Form().AutoFocus(true).WrapWithPanel(false).Body(
		u.App.Editor().Language("text").Name("editor").AllowFullscreen(false).Options(am.Schema{"fontSize": 14}),
		u.Flex().ClassName("w-full").Items(
			u.Button().Label("▼").TransformMultiple(
				func(d am.Schema) (am.Schema, error) {
					return util.Hash([]byte(d.Get("editor").(string)))
				},
				"editor",
			),
		),
		u.InputText().Name("md5").Label("MD5").Disabled(true),
		u.InputText().Name("sha1").Label("SHA1").Disabled(true),
		u.InputText().Name("sha256").Label("SHA256").Disabled(true),
		u.InputText().Name("sha512").Label("SHA512").Disabled(true),
	)
}

func (u *UI) Json2struct() ac.Form {
	return u.DualEditor(jsonCfg, comp.EditorCfg{Lang: "go"}, "", func(input any) (output any, err error) {
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
