package main

import (
	_ "GoMVC/models/OrmModel"
	_ "GoMVC/routers"

	_ "GoMVC/utils/LogHelper"

	//"github.com/astaxie/beego"

	//"GoMVC/filter"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//注册过滤器
	//beego.ErrorHandler("404",$filter.page_not_found)
	//beego.InsertFilter("*", beego.BeforeRouter, filter.FilterUser)
	// 跨域处理
	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	// 	AllowCredentials: true,
	// }))

	//beego.SetStaticPath("/static", "static")//注册了 static 目录为静态处理的目录
	beego.BConfig.RunMode = "dev"
	beego.Run()

}
