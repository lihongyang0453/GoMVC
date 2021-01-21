package main

import (
	_ "GoMVC/models/orm"
	_ "GoMVC/routers"

	"github.com/astaxie/beego"
)

func main() {

	beego.Run()
	//beego.ErrorHandler("404",$filter.page_not_found)
	//beego.InsertFilter("*", beego.BeforeRouter, $filter.FilterUser)
}
