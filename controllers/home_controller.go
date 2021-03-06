package controllers

import (
	"github.com/astaxie/beego/logs"
	"runbird-beego/models"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	page, _ := this.GetInt("page", 0)
	if page < 0 {
		page = 1
	}
	var artList []models.Article
	artList, _ = models.FindArticleWithPage(page)
	this.Data["pageCode"] = models.ConfigHomeFooterPageCode(page)
	this.Data["HasFooter"] = true

	logs.Info("IsLogin: ", this.IsLogin, this.LoginUser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)

	this.TplName = "home.html"
}
