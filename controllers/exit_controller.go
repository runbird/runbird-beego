package controllers

type ExitController struct {
	BaseController
}

//登出
func (this *ExitController) Get() {
	this.DelSession("loginuser")
	this.Redirect("/", 302)
}
