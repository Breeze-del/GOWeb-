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
	tid := c.Input().Get("tid")
	var err error
	if len(tid) == 0 {
		// 不存在 是添加操作
		err = models.AddTopic(title, content)
	} else {
		// 存在tid 说明是修改操作
		err = models.ModifyTopic(tid, title, content)
	}
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

// 预览文章
func (c *TopicController) View() {
	topic, err := models.GetTopic(c.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err.Error())
		c.Redirect("/", 302)
		return
	}
	c.Data["topic"] = topic
	// 不知道为什么模板data topicId不能识别
	c.Data["tid"] = c.Ctx.Input.Param("0")
	c.TplName = "topic_view.html"
}

// 修改文章
func (c *TopicController) Modify() {
	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err.Error())
		c.Redirect("/", 302)
		return
	}
	c.Data["topic"] = topic
	c.Data["tid"] = tid
	c.TplName = "topic_modify.html"
}
