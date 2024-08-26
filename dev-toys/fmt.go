package main

import "amisgo-examples/dev-toys/comp"

var (
	langs = []string{"json", "yaml", "toml"}

	fmtLeftCfg = &comp.EditorCfg{
		Langs: langs,
		Lang:  "json",
	}
	fmtRightCfg = &comp.EditorCfg{
		Langs: langs,
		Lang:  "yaml",
	}

	fmtEditor = comp.DualEditor(fmtLeftCfg, fmtRightCfg, func(a any) any {
		return "hi"
	})
)
