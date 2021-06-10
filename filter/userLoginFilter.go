package filter

import (
	"context"
	"net/http"
	"text/template"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//beego.InsertFilter("*", beego.BeforeRouter, FilterUser)
}

var FilterUser = func(ctx *context.Context) {
	// _, ok := ctx.Input.Session("uid").(int)
	// if !ok && ctx.Request.RequestURI != "/login" {
	// 	ctx.Redirect(302, "/login")
	// }
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}
