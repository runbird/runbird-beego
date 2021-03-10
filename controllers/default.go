package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Post() {
	logs.Info("decrepted log")
	body := this.Ctx.Output.Body([]byte("post method...."))
	if body == nil {
		os.Exit(1)
	}
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
