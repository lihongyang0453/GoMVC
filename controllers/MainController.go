package controllers

import (
	logHelper "GoMVC/utils/LogHelper"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	logHelper.LogEmail("send email to me !")
	c.TplName = "index.tpl"
}
