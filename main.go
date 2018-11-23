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
	beego.Run()
}
