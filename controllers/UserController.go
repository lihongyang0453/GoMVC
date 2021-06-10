package controllers

import (
	service "GoMVC/services"
	logHelper "GoMVC/utils/LogHelper"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	c.Data["Website"] = "bbbbbbbbbbbbbeego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.TplName = "index.tpl"
	userService := new(service.UserService)
	var totalCount int
	list := userService.GetList(1, 10, nil, "", &totalCount)
	if len(list) > 0 {
		go logHelper.LogInfo("go")
		logHelper.LogInfo("normal")
	}
	//c.TplName = "index.tpl"
	c.Render()
}

func (this *UserController) Get2() {
	// v := this.GetSession("asta")
	// if v == nil {
	// 	this.SetSession("asta", int(1))
	//    this.Data["num"] = 0
	// } else {
	// 	this.SetSession("asta", v.(int)+1)
	// 	this.Data["num"] = v.(int)
	// }
	//c.TplName = "index.tpl"

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
}
