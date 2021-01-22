package controllers

import (
	OrmModel "GoMVC/models/OrmModel"
	"fmt"

	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BaseController
}

//用户如果没有进行注册，那么就会通过反射来执行对应的函数，如果注册了就会通过interface来进行执行函数，性能上面会提升很多。
// func (c *LoginController) URLMapping() {
// 	c.Mapping("Login", c.Login)
// }

func (this *LoginController) Index() {
	this.TplName = "Index.tpl"
}

// @router /Login/:key [get]
func (this *LoginController) Login() {
	

	//this.ServeJSON()
}
