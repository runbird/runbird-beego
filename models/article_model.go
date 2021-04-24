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

func FindArticleWithPage(page int) ([]Article, error) {
	//配置文件中每页获取文章的数量
	num, err := beego.AppConfig.Int("articleListPageNum")
	if err != nil {
		logs.Error("配置文件中每页获取文章的数量异常:{}", err)
	}
	page--
	logs.Debug(">>>>>>>> page:{}", page)
	return QueryArticleWithPaeg(page, num)
}

//分页查询数据库
//limit分页查询 limit m,n
//m代表从多少位开始获取和id值无关
//n代表获取多少条数据
func QueryArticleWithPaeg(page, num int) ([]Article, error) {
	logs.Debug("limit %d , %d", page*num, num)
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	return queryArticleWithConn(sql)
}

func queryArticleWithConn(sql string) ([]Article, error) {
	sql = "select id,titile,tags,content,author,createtime from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tages := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		var updatetime int64
		updatetime = 0
		rows.Scan(&id, &title, &tages, &short, &content, &author, &createtime, &updatetime)
		art := Article{id, title, author, tages, short, content, createtime, updatetime}
		artList = append(artList, art)
	}
	return artList, nil
}
