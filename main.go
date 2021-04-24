package main

import (
	"github.com/astaxie/beego"
	_ "runbird-beego/routers"
	"runbird-beego/utils"
)

func main() {
	utils.InitMysql()
	//打开session 要么在配置文件中 设置 要么在main中设置 二选一
	//beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
