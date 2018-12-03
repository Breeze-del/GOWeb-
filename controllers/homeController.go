package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	// 首页阴影强调
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	// 获取到cate参数
	cate := c.Input().Get("cate")
	topics, err := models.GetAllTopics(cate, false)
	if err != nil {
		beego.Error(err.Error())
	}
	category, err1 := models.GetAllCategories()
	if err1 != nil {
		beego.Error(err1.Error())
	}
	c.Data["Category"] = category
	c.Data["topics"] = topics
	c.TplName = "home.html"
}
