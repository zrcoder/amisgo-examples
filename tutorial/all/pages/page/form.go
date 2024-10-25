package page

import (
	"fmt"

	"github.com/zrcoder/amisgo/comp"
)

var Form = comp.Page().Title("表单页面").Body(
	comp.Form().Mode("horizontal").
		Body(
			comp.InputText().Label("姓名").Name("name"),
			comp.InputEmail().Label("邮箱").Name("email"),
		).Go(handleSubmit),
)

func handleSubmit(m comp.Data) error {
	fmt.Println("name:", m.Get("name"))
	fmt.Println("email:", m.Get("email"))
	return nil
}
