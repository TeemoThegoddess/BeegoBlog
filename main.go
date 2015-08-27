package main

import (
	"blog/controllers"
	_ "blog/models"
	_ "blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("getLength", controllers.GetStrLength)
	beego.Run()
}
