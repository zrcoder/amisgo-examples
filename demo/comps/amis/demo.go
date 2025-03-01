package amis

import (
	"github.com/zrcoder/amisgo"
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

func Demos(app *amisgo.App) map[string]any {
	return map[string]any{
		"Base": app.Amis().Schema(app.Tpl().Tpl("amis render")),
		"Dynamic": app.Group().Body(
			app.Editor().Language("json").Name("amis").Value(dynamicContent),
			app.Amis().Name("amis"),
		),
		"Props": app.Amis().Props(schema.Schema{"tpl": "amis render"}).Value(app.Tpl()),
	}
}
