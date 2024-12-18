package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/zrcoder/amisgo-examples/doc/docs"

	"github.com/zrcoder/amisgo/comp"
)

const defaultDoc = "amisgo.md"

func Menu() (any, error) {
	options, err := getMenuOptions(docs.FS, ".")
	if err != nil {
		return nil, err
	}

	return comp.InputTree().
		Name("menu").
		ClassName("no-border").
		Value(defaultDoc).
		Options(
			options...,
		), nil
}

func Doc() any {
	return comp.Markdown().
		Src(fmt.Sprintf("%s?%s=${menu}", docsApi, docQuery))
}

func Page(menu, doc any) any {
	return comp.Wrapper().
		Body(
			comp.Page().
				Name("doc-main").
				Aside(
					comp.Image().
						Src("/static/logo-with-text.svg").
						Alt("Amisgo").
						InnerClassName("border-none").
						Height("28px"),
					menu,
				).
				AsideClassName("w-80").
				AsideResizor(true).
				Body(
					doc,
				),
		)
}

func getMenuOptions(f fs.FS, rootPath string) ([]any, error) {
	var options []any

	err := fs.WalkDir(f, rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == rootPath || !d.IsDir() && !strings.HasSuffix(path, ".md") {
			return nil
		}

		label := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		opt := comp.Option().Label(label).Value(path)
		options = append(options, opt)

		if !d.IsDir() {
			return nil
		}

		opt["children"], err = getMenuOptions(f, path)
		if err != nil {
			return err
		}
		return fs.SkipDir
	})

	return options, err
}
