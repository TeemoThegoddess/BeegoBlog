package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	var err error
	this.Data["isHome"] = true
	this.Data["isLogin"] = checkCookie(this.Ctx)
	this.TplNames = "home.html"
	this.Data["topics"], err = models.QueryTopics(true)
	if err != nil {
		beego.Error(err)
	}
}
