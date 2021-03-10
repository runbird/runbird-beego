package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"runbird-beego/util"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int //0 正常状态  1 删除
	Createtime int64
	Updatetime int64
}

//--------------数据库操作-----------------

func InsertUser(user User) (int64, error) {
	return util.ModifyDB("insert into user(username,password,status,createtime,updatetime) values(?,?,?,?,?)",
		user.Updatetime, user.Password, user.Status, user.Createtime, user.Updatetime)
}

//条件查询
func QueryUserWightCon(conn string) int {
	sql := fmt.Sprintf("select id from users %s ", conn)
	logs.Info(sql)
	row := util.QueryRowDB(sql)
	id := 0

	//TODO
	row.Scan(id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username =%s", username)
	return QueryUserWightCon(sql)
}

func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
