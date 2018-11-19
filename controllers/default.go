package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("appname: " + beego.AppConfig.String("appname") +
		"\nhttpport: " + beego.AppConfig.String("httpport"))
	c.Ctx.WriteString("\nappname: " + beego.AppPath)
}
