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
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err.Error())
	}
	c.Data["topics"] = topics
	c.TplName = "home.html"
}
