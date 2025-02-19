package util

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/ChimeraCoder/gojson"
	decqr "github.com/tuotoo/qrcode"
	"github.com/yassinebenaid/bunster/analyser"
	"github.com/yassinebenaid/bunster/generator"
	"github.com/yassinebenaid/bunster/lexer"
	"github.com/yassinebenaid/bunster/parser"
	"github.com/yosssi/gohtml"
	"github.com/zrcoder/amisgo/model"
	"github.com/zrcoder/cdor"
	"gopkg.in/yaml.v3"
)

func Json(data string) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := json.Indent(buf, []byte(data), "", "    ")
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func Toml(data string) (string, error) {
	var obj any
	err := toml.Unmarshal([]byte(data), &obj)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	err = toml.NewEncoder(buf).Encode(obj)
	if err != nil {
		return "", err
	}
	return buf.String(), err
}

func Yaml(data string) (string, error) {
	var obj any
	err := yaml.Unmarshal([]byte(data), &obj)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)
	err = encoder.Encode(&obj)
	if err != nil {
		return "", err
	}
	return buf.String(), err
}

func Html(data string) (string, error) {
	return gohtml.Format(data), nil
}

func Json2Yaml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(writer)
	err = encoder.Encode(obj)
	return writer, err
}

func Yaml2Json(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(obj)
	return writer, err
}

func Json2Toml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	err = toml.NewEncoder(writer).Encode(obj)
	return writer, err
}

func Toml2Json(data []byte) (*bytes.Buffer, error) {
	var obj any
	_, err := toml.Decode(string(data), &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(obj)
	return writer, err
}

func Toml2Yaml(data []byte) (*bytes.Buffer, error) {
	var obj any
	_, err := toml.Decode(string(data), &obj)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(buf)
	err = encoder.Encode(obj)
	return buf, err
}

func Yaml2Toml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	writer := bytes.NewBuffer(nil)
	err = toml.NewEncoder(writer).Encode(obj)
	return writer, err
}

func Json2Svg(data []byte) (*bytes.Buffer, error) {
	return toSvg(func(c *cdor.Cdor) {
		c.Json(string(data))
	})
}

func toSvg(fn func(*cdor.Cdor)) (*bytes.Buffer, error) {
	c := cdor.Ctx()
	fn(c)
	c.Theme(200) // dark theme
	data, err := c.Gen()
	return bytes.NewBuffer(data), err
}

var StructOption = struct {
	Name          string
	Pkg           string
	Tags          []string
	ConvertFloats bool
	SubStruct     bool
}{
	Name:          "Root",
	Pkg:           "model",
	Tags:          []string{"json"},
	ConvertFloats: true,
	SubStruct:     true,
}

func Json2Struct(input []byte) (string, error) {
	buf := bytes.NewBuffer(input)
	out, err := gojson.Generate(
		buf,
		gojson.ParseJson,
		StructOption.Name,
		StructOption.Pkg,
		StructOption.Tags,
		StructOption.SubStruct,
		StructOption.ConvertFloats)
	return string(out), err
}

func Hash(input []byte) (model.Schema, error) {
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

	return model.Schema{
		"md5":    resMd5,
		"sha1":   resSha1,
		"sha256": resSha256,
		"sha512": resSha512,
	}, nil
}

func DecodeQr(img []byte) (string, error) {
	buf := bytes.NewBuffer(img)
	res, err := decqr.Decode(buf)
	if err != nil {
		return "", err
	}
	return res.Content, nil
}

func Shell2Go(data []byte) (*bytes.Buffer, error) {
	script, err := parser.Parse(lexer.New(data))
	if err != nil {
		return nil, err
	}

	if err := analyser.Analyse(script); err != nil {
		return nil, err
	}

	program := generator.Generate(script)
	res := bytes.NewBuffer(nil)
	res.WriteString(program.String())
	return res, nil
}
