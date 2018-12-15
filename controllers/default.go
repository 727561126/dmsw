package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index/login.tpl"
	c.Layout = "index.tpl"

}
func (c *MainController) Login() {
	c.TplName = "main/main.html"
}
