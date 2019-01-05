package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myapp/models"
	_ "myapp/routers"
	"os"
)

// init函数自动调用
func init() {
	// 注册数据库
	models.RegisterDB()
}
func main() {
	// 开启debug模式
	orm.Debug = true
	// 自动建表 第一个参数 数据库名字 第二个参数 是否先删除表再建立表 第三个参数 是否将执行过程打印在terminal
	orm.RunSyncdb("default", false, false)

	//// 需要输入两个路径  调试得时候才能加载到views和static文件
	//beego.BConfig.WebConfig.ViewsPath = "F:/Go_Web01/src/myapp/views"
	//beego.SetStaticPath("/static", "F:/Go_Web01/src/myapp/static")

	// 创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	// 创建静态文件路径 用作展示文件
	// beego.SetStaticPath("/attachment","./attachment")
	beego.Run()
}
