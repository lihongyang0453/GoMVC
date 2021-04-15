package controllers

import beego "github.com/beego/beego/v2/server/web"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) refresh() {
	this.XSRFExpire = 7200 //动态更新过期时间
	//this.EnableXSRF = false
}
