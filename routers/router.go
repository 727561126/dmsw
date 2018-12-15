package routers

import (
	"dmsw/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "Get:Index")
	beego.Router("/login", &controllers.MainController{}, "Get:Login")
	beego.Router("/user", &controllers.UserController{})
}
