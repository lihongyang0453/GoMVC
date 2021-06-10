package routers

import (
	"GoMVC/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//常用三种注册路由形式
	//1、固定路由，这种路由注册出来的只能响应restful API 不能识别自定义的方法名称
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Get", &controllers.MainController{})

	//1.1、正则路由
	// beego.Router("/role/getByID/?:id", &controllers.RoleController{})  //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123",? 代表可为空
	// beego.Router("/role/getByID/:id([0-9]+)", &controllers.RoleController{}) //匹配正整数
	// beego.Router("/role/getByName/:username([\\w]+)", &controllers.RoleController{}) //匹配字符串
	// beego.Router("/role/SaveModel/", &controllers.RoleController{}) //
	//1.2、自定义路由、
	// beego.Router("/role/getByID", &controllers.RoleController{},"*:getByID")  
	// beego.Router("/role/getByID", &controllers.RoleController{},"get:getByID") 
	// beego.Router("/role/getByName", &controllers.RoleController{},"get,getByName") 
	// beego.Router("/role/SaveModel", &controllers.RoleController{},"post:SaveModel") 
	// beego.Router("/role/SaveModel", &controllers.RoleController{},"get,post:SaveModel") //多个 HTTP Method 指向同一个函数的示例：
	//beego.Router("/role/SaveModel/", &controllers.RoleController{},"get:GetFunc;post:PostFunc") //路由可以分别对应get请求和post请求的


	//2、自动匹配，把需要路由的控制器注册到自动路由，除了前缀两个 /:controller/:method 的匹配之外，
	//剩下的 url beego 会帮你自动化解析为参数，保存在 this.Ctx.Input.Params 当中
	beego.AutoRouter(&controllers.UserController{})

	//3、注解路由,然后在具体的controller 中的 method 方法上面写上 router 注释（// @router）就可以了
	beego.Include(&controllers.LoginController{})

	

}
