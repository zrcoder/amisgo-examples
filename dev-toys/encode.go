package main

import (
	"encoding/base64"
	"html"
	"net/url"

	"amisgo-examples/dev-toys/comp"
	"amisgo-examples/dev-toys/util"

	ac "github.com/zrcoder/amisgo/comp"
)

var (
	base64ED = comp.DualEditor(comp.EditorCfg{}, comp.EditorCfg{}, "Base64",
		func(input any) (output any, err error) {
			return base64.StdEncoding.EncodeToString([]byte(input.(string))), nil
		},
		func(input any) (output any, err error) {
			out, err := base64.StdEncoding.DecodeString(input.(string))
			if err != nil {
				return "", err
			}
			return string(out), nil
		})

	urlED = comp.DualEditor(comp.EditorCfg{}, comp.EditorCfg{}, "Url",
		func(input any) (output any, err error) {
			return url.QueryEscape(input.(string)), nil
		},
		func(input any) (output any, err error) {
			return url.QueryUnescape(input.(string))
		})
	htmlED = comp.DualEditor(comp.EditorCfg{}, comp.EditorCfg{}, "Html",
		func(input any) (output any, err error) {
			return html.EscapeString(input.(string)), nil
		},
		func(input any) (output any, err error) {
			return html.UnescapeString(input.(string)), nil
		})

	qrData []byte
	decqr  = ac.Form().ColumnCount(3).WrapWithPanel(false).Body(
		ac.Flex().Style(ac.Schema{"width": "45%"}).AlignItems("center").Items(
			ac.InputImage().Name("img").Upload(int64(10*(1<<20)), func(data []byte) (path string, err error) {
				qrData = data
				return "", err
			}),
		),
		ac.Flex().Style(ac.Schema{"width": "9%"}).AlignItems("center").Items(
			ac.Action().ActionType("submit").Icon("fa fa-arrow-right").Reload("out"),
		),
		ac.Service().Style(ac.Schema{"width": "45%"}).
			Name("out").
			GetData(func() (any, error) {
				if qrData == nil {
					return "", nil
				}
				decodecQr, err := util.DecodeQr(qrData)
				return ac.Data{"decqr": decodecQr}, err
			}).Body(
			comp.Editor(comp.EditorCfg{Name: "text", ReadOnly: true, Value: "${decqr}"}),
		),
	)
)
