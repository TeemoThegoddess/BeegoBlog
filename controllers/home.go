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
	category := this.Input().Get("category")
	label := this.Input().Get("label")

	/*if len(category) == 0 {
		this.Data["topics"], err = models.QueryTopics(true)
	} else {
		this.Data["topics"], err = models.QueryTopicsByCategory(true, category)
	}

	if len(label) == 0 {
		this.Data["topics"], err = models.QueryTopics(true)
	} else {
		this.Data["topics"], err = models.QueryTopicsByLabel(true, label)
	}*/
	this.Data["topics"], err = models.GetAllTopics(true, label, category)

	this.Data["categories"], err = models.QueryCategories(false)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
}
