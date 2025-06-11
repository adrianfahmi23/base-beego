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

    beego.GlobalControllerRouter["example-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Show",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Store",
            Router: `/store`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateStatus",
            Router: `/update-status/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["example-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["example-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
