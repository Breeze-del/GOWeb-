package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
)

type TopicController struct {
	beego.Controller
}

// 不加*是访问不到views下的资源的
func (c *TopicController) Get() {
	c.Data["IsTopic"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err.Error())
	}
	c.Data["topics"] = topics
	c.TplName = "topic.html"
}

func (c *TopicController) Post() {
	// 数据库操作之前先检查是否是处于登陆的状态
	if !(checkAccount(c.Ctx)) {
		c.Redirect("/login", 302)
		return
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	var err error
	err = models.AddTopic(title, content)
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}

// 自动路由 /topic/add 后面相当于跟的就是方法名字
func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
}
