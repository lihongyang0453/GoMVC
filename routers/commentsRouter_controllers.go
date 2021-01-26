package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["GoMVC/controllers:LoginController"] = append(beego.GlobalControllerRouter["GoMVC/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Index",
            Router: "/Index/:key",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["GoMVC/controllers:LoginController"] = append(beego.GlobalControllerRouter["GoMVC/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/Login/:key",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
