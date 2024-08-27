package util

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/yosssi/gohtml"
	"github.com/zrcoder/cdor"
	"gopkg.in/yaml.v3"
)

func Json(data []byte) (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(nil)
	err := json.Indent(buf, data, "", "    ")
	return buf, err
}

func Toml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := toml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	err = toml.NewEncoder(buf).Encode(obj)
	return buf, err
}

func Yaml(data []byte) (*bytes.Buffer, error) {
	var obj any
	err := yaml.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(buf)
	encoder.SetIndent(2)
	err = encoder.Encode(&obj)
	return buf, err
}

func Html(data []byte) (*bytes.Buffer, error) {
	res := gohtml.FormatBytes(data)
	return bytes.NewBuffer(res), nil
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
	data, err := c.Gen()
	return bytes.NewBuffer(data), err
}
