package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	// 获取到op然后进行处理
	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
		// 不用渲染操作，所以直接return
		return
	case "del":
		name := c.Input().Get("id")
		// 删除分类的所有文章
		category := c.Input().Get("category")
		if len(name) == 0 {
			break
		}
		err := models.DeleteCategory(name)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteTopics(category)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
		return
	}
	c.Data["IsCategory"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	var err error
	c.Data["categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "category.html"
}
