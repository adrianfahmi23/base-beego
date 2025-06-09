package main

import (
	_ "example-beego/routers"
	"example-beego/utils"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.WebConfig.CommentRouterPath = "controllers"
	utils.InitDB()

	beego.Run()
}
