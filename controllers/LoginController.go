package controllers

import (
	OrmModel "GoMVC/models/OrmModel"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

//用户如果没有进行注册，那么就会通过反射来执行对应的函数，如果注册了就会通过interface来进行执行函数，性能上面会提升很多。
func (c *LoginController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Login", c.Login)

}

// @router /Index/:key [get]
func (this *LoginController) Index() {
	user := new(OrmModel.User)
	var totalCount int
	user.GetList(1, 10, nil, "", &totalCount)
	this.TplName = "Index.tpl"
}

// @router /Login/:key [post]
func (this *LoginController) Login() {
	user := new(OrmModel.User)
	user.CreatedTime = time.Now()

	var totalCount int
	list := user.GetList(1, 10, nil, "", &totalCount)
	if len(list) > 0 {

	}
	//this.ServeJSON()
}
