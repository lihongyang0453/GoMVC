package controllers

import (
	"GoMVC/models"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *LoginController) StaticBlock() {
	this.ServeJSON()
}

// @router /all/:key [get]
func (this *LoginController) AllBlock() {
	u := models.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
	}
}
