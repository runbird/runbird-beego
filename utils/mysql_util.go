package utils

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var db *sql.DB

func InitMysql() {
	logs.Info("InitMysql ....")
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	// orm.RegisterDriver(driverName,orm.DRMySQL)

	//数据库连接
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpwd := beego.AppConfig.String("mysqlpwd")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := mysqluser + ":" + mysqlpwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4"
	logs.Info(dbConn)

	//
	db1, e := sql.Open(driverName, dbConn)
	if e != nil {
		logs.Error(e)
	} else {
		db = db1
		logs.Info("MySQL连接成功...")
	}

	CreateTableWithUser()
	CreateTableWithArticle()
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, e := db.Exec(sql, args...)
	if e != nil {
		logs.Error(e)
		return 0, e
	}

	row, e := result.RowsAffected()
	if e != nil {
		logs.Error(e)
		return 0, e
	}
	return row, nil
}

//创建用户
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10),
		updatetime INT(10)
		);`
	_, e := ModifyDB(sql)
	if e != nil {
		logs.Error("create table users have a error")
	}
}

//创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	_, e := ModifyDB(sql)
	if e != nil {
		logs.Error("create table article have a error")
	}
}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
