package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	if !checkCookie(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	/*op := this.Input().Get("op")

	switch op {
	case "add":
		name := this.Input().Get("category")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}*/

	this.Data["isLogin"] = "true"
	this.TplNames = "category.html"
	this.Data["isCategory"] = "true"
	var err error
	this.Data["categories"], err = models.QueryCategories(false)
	if err != nil {
		beego.Error(err)
	}
}

func (this *CategoryController) Add() {
	categotyName := this.Input().Get("category")
	err := models.AddCategory(categotyName)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Redirect("/category", 302)
	return
}

func (this *CategoryController) Delete() {
	categoryId := this.Ctx.Input.Param("0")
	err := models.DeleteCategory(categoryId)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Redirect("/category", 302)
	return
}
