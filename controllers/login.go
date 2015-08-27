package controllers

import (
	/*	"crypto/md5"
		"encoding/hex"*/
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "login.html"
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("account", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 301)
		return
	}
}

func (this *LoginController) Post() {
	/*	s := md5.New()*/
	account := this.Input().Get("account")
	pwd := this.Input().Get("pwd")
	aotuLogin := this.Input().Get("isAutoLogin") == "on"

	if beego.AppConfig.String("adminname") == account &&
		beego.AppConfig.String("adminpwd") == pwd {
		maxAge := 0
		if aotuLogin {
			maxAge = 1<<20 - 1
		}
		/*s.Write([]byte(pwd))
		md5pwd := hex.EncodeToString(s.Sum(nil))*/
		this.Ctx.SetCookie("account", account, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}

func checkCookie(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("account")
	if err != nil {
		return false
	}
	account := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("adminname") == account &&
		beego.AppConfig.String("adminpwd") == pwd
}
