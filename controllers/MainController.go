package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	//userService := new(service.UserService)
	// var totalCount int
	// list := userService.GetList(1, 10, nil, "", &totalCount)
	// if len(list) > 0 {

	// }
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
