package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/zrcoder/amisgo-examples/doc/docs"

	"github.com/zrcoder/amisgo/comp"
)

func Menu() (any, error) {
	options, err := getMenuOptions(docs.FS, ".")
	if err != nil {
		return nil, err
	}

	return comp.InputTree().Name("menu").Value("amis.md").Options(
		options...,
	), nil
}

func Doc() any {
	return comp.Markdown().
		Name("doc").
		Src(fmt.Sprintf("%s?%s=${menu}", docsApi, docQuery))
}

func Page(menu, doc any) any {
	return comp.Page().Title(comp.Image().InnerClassName("no-border").Height("24px").Alt("Amisgo").Src("https://raw.githubusercontent.com/zrcoder/amisgo-assets/refs/heads/main/logo-with-text.svg")).Aside(menu).Body(doc)
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
