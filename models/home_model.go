package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"html/template"
	"runbird-beego/utils"
	"strconv"
	"strings"
)

//前端页面所需对象
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

type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		files, e := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		e = files.Execute(&buffer, homeParam)
		if e != nil {
			logs.Error("home modle execute template error :{}", e)
		}
		htmlHome += buffer.String()
	}
	logs.Info("htmlHome :{}", htmlHome)
	return template.HTML(htmlHome)
}

func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsPara := strings.Split(tags, "&")
	for _, tag := range tagsPara {
		tagLink = append(tagLink, TagLink{
			TagName: tag,
			TagUrl:  "/?tag=" + tag,
		})
	}
	return tagLink
}

//翻页
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	num := GetArticleRowsNum()
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	//总页数
	allPageNum := (num-1)/pageRow + 1
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	//首页
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	//尾页
	//当前页数大于总页数，那么下一页不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}
