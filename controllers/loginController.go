package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	// 获取是不是退出
	// 为true会路由失败 我不知道为什么
	isExit := c.Input().Get("exit")
	fmt.Println(isExit)
	if isExit == "True" {
		// 立即删除cookie保留得信息
		c.Ctx.SetCookie("usname", " ", -1, "/")
		c.Ctx.SetCookie("password", " ", -1, "/")
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	// input方法获得post传递过来的值 以map的方式
	//c.Ctx.WriteString(fmt.Sprint(c.Input()))
	usname := c.Input().Get("usname")
	password := c.Input().Get("password")
	autologin := c.Input().Get("autologin") == "on"

	// 比较配置文件和post表单中account信息
	if beego.AppConfig.String("usname") == usname &&
		beego.AppConfig.String("password") == password {
		// cookie 最大寿命
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}
		// 写入cookie
		c.Ctx.SetCookie("usname", usname, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
	}
	// 重定向到首页去 code301（301表示永久重定向 302暂时性重定向） 访问主页
	c.Redirect("/", 301)
}

// 检查cookies 判断是不是登陆的状态 用于界面显示
func checkAccount(ctx *context.Context) bool {
	//context可以从controller中获得 也可以从context包中获得
	//他是对request（input）和response（output）的封装

	//ck, err := ctx.Request.Cookie("usname")
	//if err != nil {
	//	return false
	//}
	//usname := ck.Value
	//ck, err1 := ctx.Request.Cookie("password")
	//if err1 != nil {
	//	return false
	//}
	//password := ck.Value
	//// 通过cookie获得的account与配置文件中account比较来确定 是不是处于登陆状态
	//return beego.AppConfig.String("usname") == usname &&
	//	beego.AppConfig.String("password") == password
	return true
}
