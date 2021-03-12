package controllers

import (
	"github.com/astaxie/beego/logs"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	logs.Info("IsLogin: ", this.IsLogin, this.LoginUser)
	page, _ := this.GetInt("page", 0)
	if page < 0 {
		page = 1
	}
	this.Data["page"] = 1
	this.Data["HasFooter"] = true

	this.TplName = "home.html"
}
