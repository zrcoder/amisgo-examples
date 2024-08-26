package main

import "github.com/zrcoder/amisgo/comp"

var diff = comp.Form().Title("").Static(true).Actions().Body(
	comp.DiffEditor().
		Label("Diff Editor").
		Name("diff").
		ReadOnly(false).
		Size("xxl"),
)
