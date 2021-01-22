package routers

import (
	"GoMVC/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//常用三种注册路由形式
	//1、固定路由，这种路由注册出来的只能响应restful API 不能识别自定义的方法名称
	beego.Router("/", &controllers.MainController{})

	//2、自动匹配，把需要路由的控制器注册到自动路由，除了前缀两个 /:controller/:method 的匹配之外，
	//剩下的 url beego 会帮你自动化解析为参数，保存在 this.Ctx.Input.Params 当中
	beego.AutoRouter(&controllers.UserController{})

	//3、注解路由,然后在具体的controller 中的 method 方法上面写上 router 注释（// @router）就可以了
	beego.Include(&controllers.LoginController{})
}
