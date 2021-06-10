package controllers

import (
	service "GoMVC/services"
	logHelper "GoMVC/utils/LogHelper"

	//beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	userService := new(service.UserService)
	var totalCount int
	list := userService.GetList(1, 10, nil, "", &totalCount)
	if len(list) > 0 {
		go logHelper.LogInfo("go")
		logHelper.LogInfo("normal")
	}
	// CacheHelper.Add("name", "lhy")
	// CacheHelper.AddNX("name", "lhy")
	// model := userService.FindById(20)
	// if model.DisplayName != "" {

	// }
	// model.Id = 0
	// flag := userService.Insert(model)
	// if flag {

	// }
	// flag := userService.DeleteById(17)
	// if flag {

	// }
	// model.CreatedTime = time.Now()
	// userService.Update(model)
	c.TplName = "index.tpl"
}
