package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["example-beego/controllers:AuthController"] = append(beego.GlobalControllerRouter["example-beego/controllers:AuthController"],
        beego.ControllerComments{
            Method: "CheckToken",
            Router: `/check-login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:AuthController"] = append(beego.GlobalControllerRouter["example-beego/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:MainController"] = append(beego.GlobalControllerRouter["example-beego/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:UsersController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:UsersController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UsersController"],
        beego.ControllerComments{
            Method: "GetByID",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:UsersController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UsersController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/store`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
