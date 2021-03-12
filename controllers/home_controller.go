package controllers

import (
	"github.com/astaxie/beego/logs"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	logs.Info("IsLogin: ", this.IsLogin, this.LoginUser)
	this.TplName = "home.html"
}
