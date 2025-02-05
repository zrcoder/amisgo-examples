package pages

import (
	"encoding/base64"
	"html"
	"net/url"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"
	am "github.com/zrcoder/amisgo/model"
)

func Base64ED(app *amisgo.App) any {
	return comp.DualEditor(
		app,
		comp.EditorCfg{}, comp.EditorCfg{}, "Base64",
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
}

func UrlED(app *amisgo.App) any {
	return comp.DualEditor(
		app,
		comp.EditorCfg{}, comp.EditorCfg{}, "Url",
		func(input any) (output any, err error) {
			return url.QueryEscape(input.(string)), nil
		},
		func(input any) (output any, err error) {
			return url.QueryUnescape(input.(string))
		})
}

func HtmlED(app *amisgo.App) any {
	return comp.DualEditor(
		app,
		comp.EditorCfg{}, comp.EditorCfg{}, "Html",
		func(input any) (output any, err error) {
			return html.EscapeString(input.(string)), nil
		},
		func(input any) (output any, err error) {
			return html.UnescapeString(input.(string)), nil
		})
}

func Decqr(app *amisgo.App) any {
	var qrData []byte
	return app.Form().ColumnCount(3).WrapWithPanel(false).Body(
		app.Flex().Style(am.Schema{"width": "45%"}).AlignItems("center").Items(
			app.InputImage().Name("img").Upload(int64(10*(1<<20)), func(data []byte) (path string, err error) {
				qrData = data
				return "", err
			}),
		),
		app.Flex().Style(am.Schema{"width": "10%"}).AlignItems("center").Items(
			app.Action().ActionType("submit").Label("▶︎").Reload("out"),
		),
		app.Service().Style(am.Schema{"width": "45%"}).
			Name("out").
			GetData(func() (any, error) {
				if qrData == nil {
					return "", nil
				}
				decodecQr, err := util.DecodeQr(qrData)
				return am.Schema{"decqr": decodecQr}, err
			}).Body(
			comp.Editor(app, comp.EditorCfg{Name: "text", ReadOnly: true, Value: "${decqr}"}),
		),
	)
}
