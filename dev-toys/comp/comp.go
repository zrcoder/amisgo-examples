package comp

import "github.com/zrcoder/amisgo"

type Comp struct {
	*amisgo.App
}

func New(app *amisgo.App) *Comp {
	return &Comp{app}
}
