package controllers

import(
	"github.com/astaxie/beego"
)

type UserController struct{
	beego.Controller
}

func (c *UserController) ProFile() {
	c.TplName = "user/ProFile.html"
}

