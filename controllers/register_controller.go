package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

//注册
func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	logs.Info(username, password, repassword)
	fmt.Println(username, password, repassword)

	//注册之前检测是否已经被注册
	mod.Q
}
