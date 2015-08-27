package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	if !checkCookie(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.TplNames = "topic.html"
	this.Data["isLogin"] = "true"
	this.Data["isTopic"] = "true"
	var err error
	this.Data["topics"], err = models.QueryTopics(false)
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Add() {
	this.Data["isLogin"] = "true"
	var err error
	this.Data["Categories"], err = models.QueryCategories(true)
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "addTopic.html"
}

func (this *TopicController) Post() {
	if !checkCookie(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.Data["isLogin"] = "true"
	topicName := this.Input().Get("topicName")
	topicContent := this.Input().Get("topicContent")
	topicId := this.Input().Get("topicId")
	category := this.Input().Get("category")

	var err error
	if len(topicId) == 0 {
		if len(category) != 0 {
			err = models.AddTopic(topicName, category, topicContent)
		} else {
			this.Redirect("/category", 302)
			return
		}
	} else {
		err = models.ModifyTopic(topicId, topicName, category, topicContent)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 301)
	return
}

func (this *TopicController) View() {
	this.Data["isLogin"] = checkCookie(this.Ctx)
	this.TplNames = "viewTopic.html"
	tid := this.Ctx.Input.Param("0")
	var err error

	this.Data["Topic"], err = models.QueryTopic(tid, false)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Tid"] = this.Ctx.Input.Param("0")
	this.Data["Comments"], err = models.GetAllComments(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
}

func (this *TopicController) Modify() {
	if !checkCookie(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.Data["isLogin"] = "true"
	tid := this.Ctx.Input.Param("0")
	var err error
	this.Data["Topic"], err = models.QueryTopic(tid, true)
	this.Data["Categories"], err = models.QueryCategories(true)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.TplNames = "modifyTopic.html"
	this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
	tid := this.Ctx.Input.Param("0")
	err := models.DeleteTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Redirect("/", 302)
	return
}
