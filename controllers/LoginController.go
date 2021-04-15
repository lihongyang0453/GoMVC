package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

// //用户如果没有进行注册，那么就会通过反射来执行对应的函数，如果注册了就会通过interface来进行执行函数，性能上面会提升很多。
// func (c *LoginController) URLMapping() {
// 	c.Mapping("Index", c.Index)
// 	c.Mapping("Login", c.Login)

// }

// // @router /Index [get
// func (this *LoginController) Index() {
// 	this.Data["Website"] = "beego.e"
// 	this.Data["Email"] = "astaxie@gmail.cm"
// 	//this.TplName = "index.tpl
// }

// // @router /Login [post
// func (this *LoginController) Login() {
// 	//this.ServeJSON()
// 	this.Ctx.WriteString("Login")
// }

func (c *LoginController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *LoginController) StaticBlock() {
	this.Ctx.WriteString("Staticlock")
}

// @router /all/:key [get]
func (this *LoginController) AllBlock() {
	this.Ctx.WriteString("Staticlock")
}
