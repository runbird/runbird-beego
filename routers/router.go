package routers

import (
	"github.com/astaxie/beego"
	"runbird-beego/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/login", &controllers.ExitController{})

	beego.Router("/article/add", &controllers.AddArticleController{})
}
