package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	tid := this.Input().Get("topicId")
	nickname := this.Input().Get("nickname")
	content := this.Input().Get("content")

	err := models.AddComment(tid, nickname, content)
	if err != nil {
		this.Redirect("/", 302)
		return
	}

	this.Data["Comments"], err = models.GetAllComments(tid)
	this.Redirect("/topic/view/"+tid, 302)
	return
}

func (this *ReplyController) Delete() {
	tid := this.Ctx.Input.Param("0")
	id := this.Ctx.Input.Param("1")

	err := models.DeleteComment(id)
	if err != nil {
		this.Redirect("/", 302)
		return
	}

	this.Data["Comments"], err = models.GetAllComments(tid)
	this.Redirect("/topic/view/"+tid, 302)
	return
}
