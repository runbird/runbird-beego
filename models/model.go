package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"runbird-beego/utils"
)

func init() {
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	driver := orm.RegisterDriver(driverName, orm.DRMySQL)
	if driver != nil {
		logs.Error("注册数据库驱动错误")
	}

	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		utils.LogError("连接数据库出错")
		return
	}
	utils.LogInfo("连接数据库成功")
}
