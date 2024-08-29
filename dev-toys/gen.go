package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"

	"amisgo-examples/dev-toys/comp"
	"amisgo-examples/dev-toys/util"

	ac "github.com/zrcoder/amisgo/comp"
)

var jsonGraph = comp.EditorImg("json", func(s any) (any, error) {
	src := s.(string)
	buf, err := util.Json2Svg([]byte(src))
	return regularSvgData(buf.Bytes(), err)
})

var qrcode = ac.Form().ColumnCount(2).WrapWithPanel(false).Body(
	ac.Wrapper().Style(ac.Schema{"width": "50%"}).Body(
		comp.Editor(comp.EditorCfg{Lang: "text", Name: "editor"}),
	),
	ac.Flex().Style(ac.Schema{"width": "50%"}).AlignItems("center").Items(
		ac.QRCode().Name("qrcode").Value("${editor}").CodeSize(256).Level("M"),
	),
).Actions()

var hash = ac.Form().WrapWithPanel(false).Body(
	ac.Editor().Language("text").Name("editor").AllowFullscreen(false),
	ac.Flex().Style(ac.Schema{"width": "100%"}).Items(
		ac.Button().Icon("fa fa-arrow-down").TransformMultiple("$editor", "done", func(a any) (ac.Data, error) {
			input := []byte(a.(string))

			h := md5.New()
			if _, err := h.Write(input); err != nil {
				return nil, err
			}
			resMd5 := hex.EncodeToString(h.Sum(nil))

			h = sha1.New()
			if _, err := h.Write(input); err != nil {
				return nil, err
			}
			resSha1 := hex.EncodeToString(h.Sum(nil))

			h = sha256.New()
			if _, err := h.Write(input); err != nil {
				return nil, err
			}
			resSha256 := hex.EncodeToString(h.Sum(nil))

			h = sha512.New()
			if _, err := h.Write(input); err != nil {
				return nil, err
			}
			resSha512 := hex.EncodeToString(h.Sum(nil))

			return ac.Data{
				"md5":    resMd5,
				"sha1":   resSha1,
				"sha256": resSha256,
				"sha512": resSha512,
			}, nil
		}),
	),
	ac.InputText().Name("md5").Label("MD5").Disabled(true),
	ac.InputText().Name("sha1").Label("SHA1").Disabled(true),
	ac.InputText().Name("sha256").Label("SHA256").Disabled(true),
	ac.InputText().Name("sha512").Label("SHA512").Disabled(true),
)

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
