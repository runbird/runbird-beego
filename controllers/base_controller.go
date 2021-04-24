package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	LoginUser interface{}
}

//判断是否登录
func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginuser")
	logs.Info("get loginuser session: ", loginuser)
	if loginuser != nil {
		this.IsLogin = true
		this.LoginUser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
