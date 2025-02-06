package ui

import (
	"encoding/base64"
	"html"
	"net/url"

	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"
	am "github.com/zrcoder/amisgo/model"
)

type Encoders struct {
	*amisgo.App
}

func NewEncoders(app *amisgo.App) *Encoders { return &Encoders{App: app} }

func (e *Encoders) Base64ED() any {
	return comp.DualEditor(
		e.App,
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

func (e *Encoders) UrlED() any {
	return comp.DualEditor(
		e.App,
		comp.EditorCfg{}, comp.EditorCfg{}, "Url",
		func(input any) (output any, err error) {
			return url.QueryEscape(input.(string)), nil
		},
		func(input any) (output any, err error) {
			return url.QueryUnescape(input.(string))
		})
}

func (e *Encoders) HtmlED() any {
	return comp.DualEditor(
		e.App,
		comp.EditorCfg{}, comp.EditorCfg{}, "Html",
		func(input any) (output any, err error) {
			return html.EscapeString(input.(string)), nil
		},
		func(input any) (output any, err error) {
			return html.UnescapeString(input.(string)), nil
		})
}

func (e *Encoders) Decqr() any {
	var qrData []byte
	return e.Form().ColumnCount(3).WrapWithPanel(false).Body(
		e.Flex().Style(am.Schema{"width": "45%"}).AlignItems("center").Items(
			e.InputImage().Name("img").Upload(int64(10*(1<<20)), func(data []byte) (path string, err error) {
				qrData = data
				return "", err
			}),
		),
		e.Flex().Style(am.Schema{"width": "10%"}).AlignItems("center").Items(
			e.Action().ActionType("submit").Label("▶︎").Reload("out"),
		),
		e.Service().Style(am.Schema{"width": "45%"}).
			Name("out").
			GetData(func() (any, error) {
				if qrData == nil {
					return "", nil
				}
				decodecQr, err := util.DecodeQr(qrData)
				return am.Schema{"decqr": decodecQr}, err
			}).Body(
			comp.Editor(e.App, comp.EditorCfg{Name: "text", ReadOnly: true, Value: "${decqr}"}),
		),
	)
}
