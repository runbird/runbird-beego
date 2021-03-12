package controllers

import (
	"fmt"
	"runbird-beego/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

/*
当访问/add路径的时候回触发AddArticleController的Get方法
响应的页面是通过TpName
*/
func (this *AddArticleController) Get() {
	this.TplName = "write_article.html"
}

func (this *AddArticleController) Post() {
	title := this.GetString("title")
	author := this.GetString("author")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	fmt.Printf("addcontroller titile:%s, author:%s, tags:%s, short:%s, content:%s ", title, author, tags, short, content)
	article := models.Article{
		Id:         0,
		Title:      title,
		Author:     author,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Createtime: time.Now().Unix(),
		Updatetime: time.Now().Unix(),
	}
	_, err := models.AddArticle(article)
	//返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	this.Data["json"] = response
	this.ServeJSON()
}
