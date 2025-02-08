package ui

import (
	"encoding/base64"
	"html"
	"net/url"

	ac "github.com/zrcoder/amisgo/comp"

	"github.com/zrcoder/amisgo-examples/dev-toys/comp"
	"github.com/zrcoder/amisgo-examples/dev-toys/util"
)

func (u *UI) Base64ED() ac.Form {
	return u.DualEditor(
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

func (u *UI) UrlED() ac.Form {
	return u.DualEditor(
		comp.EditorCfg{}, comp.EditorCfg{}, "Url",
		func(input any) (output any, err error) {
			return url.QueryEscape(input.(string)), nil
		},
		func(input any) (output any, err error) {
			return url.QueryUnescape(input.(string))
		})
}

func (u *UI) HtmlED() ac.Form {
	return u.DualEditor(
		comp.EditorCfg{}, comp.EditorCfg{}, "Html",
		func(input any) (output any, err error) {
			return html.EscapeString(input.(string)), nil
		},
		func(input any) (output any, err error) {
			return html.UnescapeString(input.(string)), nil
		})
}

func (u *UI) Decqr() ac.Form {
	return u.QrcodeEditor(
		func(data []byte) (path string, err error) {
			u.qrData = data
			return "", err
		},
		func() (string, error) {
			if u.qrData == nil {
				return "", nil
			}
			return util.DecodeQr(u.qrData)
		},
	)
}
