package models

import (
	"github.com/astaxie/beego/logs"
	"runbird-beego/utils"
)

type Article struct {
	Id         int
	Title      string
	Author     string
	Tags       string
	Short      string
	Content    string
	Createtime int64
	Updatetime int64
}

func AddArticle(article Article) (int64, error) {
	i, e := InsertArticle(article)
	return i, e
}

//-----------数据库操作---------------

//插入一篇文章
func InsertArticle(article Article) (int64, error) {
	row, e := utils.ModifyDB("insert into article(title,author,tags,short,content,createtime,updatetime) values(?,?,?,?,?,?,?)",
		article.Title, article.Author, article.Tags, article.Short, article.Content, article.Createtime, article.Updatetime)
	if e != nil {
		logs.Error("insert into article have a error", e)
	}
	return row, e
}
