package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myapp/controllers"
	"myapp/models"
	_ "myapp/routers"
)

func init() {
	models.RegisterDB()
}
func main() {
	// 开启debug模式
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
