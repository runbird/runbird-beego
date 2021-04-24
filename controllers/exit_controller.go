package controllers

type ExitController struct {
	BaseController
}

//登出
func (this *ExitController) Get() {
	//清除登录信息
	this.DelSession("loginuser")
	//重定位
	this.Redirect("/", 302)
}
