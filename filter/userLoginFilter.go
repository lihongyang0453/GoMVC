package filter

// import (
// 	"context"
// 	"net/http"
// 	"text/template"

// 	beego "github.com/beego/beego/v2/server/web"
// )

// var FilterUser = func(ctx *context.Context) {
// 	url:=ctx.Request.RequestURI
// 	_, ok := ctx.Input.Session("uid").(int)
// 	if !ok {
// 		ctx.Redirect("/login", 302)
// 	}
// }

// func page_not_found(rw http.ResponseWriter, r *http.Request) {
// 	t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/404.html")
// 	data := make(map[string]interface{})
// 	data["content"] = "page not found"
// 	t.Execute(rw, data)
// }
