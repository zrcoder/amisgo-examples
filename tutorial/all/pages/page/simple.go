package page

import (
	"fmt"
	"time"

	"github.com/zrcoder/amisgo/comp"
)

var Simple = comp.Page().
	Title("标题").
	Remark(comp.Remark().
		Title("标题").
		Body("这是一段描述问题，注意到了没，还可以设置标题。而且只有点击了才弹出来。").
		Icon("question-mark").
		Placement("right").
		Trigger("click").
		RootClose(true),
	).Body(
	"内容部分. 可以使用 \\${var} 获取变量。如: `\\$date`: ${date}</br>",
	comp.Image().Src("/static/amisgo.png").EnlargeAble(true),
).Aside("    边栏部分").
	ClassName("white-space-pre").
	Toolbar("工具栏").
	InitData(getDate)

func getDate() (any, error) {
	y, m, d := time.Now().Date()
	mm := time.Now().UnixNano()
	return map[string]string{"date": fmt.Sprintf("%d-%d-%d %d", y, m, d, mm)}, nil
}
