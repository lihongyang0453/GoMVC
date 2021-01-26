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
	logHelper.LogInfo("success")
	logHelper.LogError("error")
	logHelper.LogWarning("warning")
	c.TplName = "index.tpl"
}
