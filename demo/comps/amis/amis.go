package amis

import (
	"github.com/zrcoder/amisgo"
	"github.com/zrcoder/amisgo-examples/demo/comps/app"
	"github.com/zrcoder/amisgo/schema"
)

const (
	dynamicContent = `
{
   "label": "弹框",
   "type": "button",
   "actionType": "dialog",
   "dialog": {
     "title": "弹框",
     "body": "这是个简单的弹框。"
   }
 }`
)

func Demos(a *amisgo.App) []app.Demo {
	return []app.Demo{
		{Name: "Base", View: a.Amis().Schema(a.Tpl().Tpl("amis render"))},
		{Name: "Dynamic", View: a.Group().Body(
			a.Editor().Language("json").Name("amis").Value(dynamicContent),
			a.Amis().Name("amis"),
		)},
		{Name: "Props", View: a.Amis().Props(schema.Schema{"tpl": "amis render"}).Value(a.Tpl())},
	}
}
