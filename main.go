package main

import (
	"github.com/astaxie/beego"
	_ "runbird-beego/routers"
	"runbird-beego/utils"
)

func main() {
	utils.InitMysql()
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
