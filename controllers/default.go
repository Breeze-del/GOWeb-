package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// 这里的值会在index.tpl中被用到
	c.TplName = "index.tpl"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@...d"
	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false
	// 结构体
	type u struct {
		Name string
		Age  int
		Sex  string
	}
	user := &u{
		Name: "Joe",
		Age:  18,
		Sex:  "Male",
	}
	c.Data["User"] = user
	// Slice
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	c.Data["nums"] = nums
	// 模板变量
	c.Data["TplVar"] = "hey guys"
	// bee-go 内置模板函数
	c.Data["html"] = "<div>hello beego</div>"
}
