package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"runbird-beego/utils"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
	Updatetime int64
}

func AddArticle(article Article) (int64, error) {
	i, e := InsertArticle(article)
	SetArticleRowsNum()
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

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	num, e := beego.AppConfig.Int("articleListPageNum")
	if e != nil {
		logs.Error("lack of beego.AppConfig.articleListPageNum .. {}", e)
	}
	page--
	logs.Info("-----> FindArticleWithPage :", page)
	return QueryArticleWithPage(page, num)
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticleWithConn(sql)
}

func QueryArticleWithConn(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime,updatetime from article " + sql
	rows, e := utils.QueryDB(sql)
	if e != nil {
		return nil, e
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		var updatetime int64
		createtime = 0
		updatetime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime, &updatetime)
		art := Article{id, title, tags, short, content, author, createtime, updatetime}
		artList = append(artList, art)
	}
	return artList, nil
}

//翻页
//存储表的行数
var articleRowsNum = 0

//只有首次获取行数的时候采取统计表里数据行数
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

func QueryArticleRowNum() int {
	rows, e := utils.QueryDB("select count(*) from article")
	num := 0
	if e != nil {
		rows.Scan(&num)
	}
	return num
}

func SetArticleRowsNum() {
	articleRowsNum = QueryArticleRowNum()
}
