package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myapp/models"
	_ "myapp/routers"
)

// init函数自动调用
func init() {
	// 注册数据库
	models.RegisterDB()
}
func main() {
	// 开启debug模式
	orm.Debug = true
	// 自动建表 将结果在terminal上显示
	orm.RunSyncdb("default", false, true)

	//需要输入两个路径  调试得时候才能加载到views和static文件
	//beego.BConfig.WebConfig.ViewsPath="F:/Go_Web01/src/myapp/views"
	//beego.SetStaticPath("/static","F:/Go_Web01/src/myapp/static")
	beego.Run()
}
