package main

import (
	_ "GoMVC/models/OrmModel"
	_ "GoMVC/routers"

	_ "GoMVC/utils/LogHelper"

	//"github.com/astaxie/beego"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//注册过滤器
	//beego.ErrorHandler("404",$filter.page_not_found)
	//beego.InsertFilter("*", beego.BeforeRouter, $filter.FilterUser)

	beego.Run()

}
