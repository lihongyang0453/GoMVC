package controllers

import (
	//beego "github.com/beego/beego/v2/server/web"
	model "GoMVC/models/BaseModel"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	CurrentUserInfo model.UserInfo
}

func (this *BaseController) Refresh() {
	this.XSRFExpire = 7200 //动态更新过期时间
	//this.EnableXSRF = false
}
