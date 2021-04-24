package models

import (
	"bytes"
	"github.com/astaxie/beego/logs"
	"html/template"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	UpdateTime string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

//我们需要将从数据库中查询出来的数据，转为对应的结构体对象，所以先设计结构体，这里我们需要考虑如果用户是登录状态，那么是可以修改或删除某一篇文章。
//当然，如果没有登录，那么只能查看。所以在设计结构体的时候，我们直接创建了修改和删除的链接字段。
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		logs.Info("---->tag :{}", art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		t, err := template.ParseFiles("views/block/home_block.html")
		if err != nil {
			logs.Error("MakeHomeBlocks have a error:{}", err)
		}
		buffer := bytes.Buffer{}
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	logs.Info("htmlHome:--->", htmlHome)
	return template.HTML(htmlHome)
}

func createTagsLinks(tags string) []TagLink {
	var tagsLink []TagLink
	tagsParam := strings.Split(tags, "&")
	for _, tag := range tagsParam {
		tagsLink = append(tagsLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagsLink
}
