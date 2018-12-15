package routers

import (
	"dmsw/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
