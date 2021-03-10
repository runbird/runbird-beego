package routers

import (
	"github.com/astaxie/beego"
	"runbird-beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
}
