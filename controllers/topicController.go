package controllers

import "github.com/astaxie/beego"

type TopicContrpller struct {
	beego.Controller
}

func (c TopicContrpller) Get() {
	c.TplName = "topic.html"
}
