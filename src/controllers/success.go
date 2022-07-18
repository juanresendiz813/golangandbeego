package controllers

import (
	"github.com/astaxie/beego"
)

type SuccessController struct {
	beego.Controller
}

func (c *SuccessController) Get() {
	c.Data["Website"] = "me.me"
	c.Data["Email"] = "info@gmail.com"
	c.TplName = "success.tpl"
}
