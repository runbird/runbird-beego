package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"runbird-beego/models"
	"runbird-beego/utils"
	"time"
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
	id := models.QueryUserWithUsername(username)
	if id > 0 {
		logs.Info("registed id:", id)
		this.Data["json"] = map[string]interface{}{"code": 0, "content": "该用户已经注册！"}
		this.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	user := models.User{
		Id:         0,
		Username:   username,
		Password:   password,
		Status:     0,
		Createtime: time.Now().Unix(),
		Updatetime: time.Now().Unix(),
	}
	_, e := models.InsertUser(user)
	if e != nil {
		logs.Error("insert New User have a error...", e)
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	this.ServeJSON()
}
