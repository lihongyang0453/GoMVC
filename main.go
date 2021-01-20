package main

import (
	_ "GoMVC/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
	//beego.ErrorHandler("404",$filter.page_not_found)
	//beego.InsertFilter("*", beego.BeforeRouter, $filter.FilterUser)
}
