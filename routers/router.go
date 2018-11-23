package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"myapp/controllers"
)

func init() {
	// beego.Router(“/user/:username([\\w]+)“, &controllers.RController{})
	//正则字符串匹配 //例如对于URL”/user/astaxie”可以匹配成功，此时变量”:username”值为”astaxie”
	//c.Ctx.Input.Param(":username") 取出正则路由中的值

	// 自定义路由匹配 beego.Router("/staticblock/:key", &CMSController{}, "get:StaticBlock")
	// 第三个参数为路由匹配方法 *: 为任何访问的匹配这个方法

	// 固定路由方法 方法交给controller取实现
	beego.Router("/", &controllers.MainController{})
	// 基本路由方法
	beego.Get("/hello", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})

	// namespace路由 好处是集中管理路由 可以使用过滤器和判断条件 比较方便
	ns := beego.NewNamespace("/v1",
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Context.Request.UserAgent() != "" {
				return true
			}
			return false
		}),
		beego.NSGet("/hello", func(ctx *context.Context) {
			ctx.Output.Body([]byte("hello world"))
		}),
		beego.NSRouter("/beego", &controllers.MainController{}),
		beego.NSNamespace("/api",
			beego.NSRouter("/index", &controllers.MainController{})),
	)
	// 注册路由
	beego.AddNamespace(ns)
}
