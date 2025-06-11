package main

import (
	"example-beego/database"
	"example-beego/models"
	_ "example-beego/routers"
	"example-beego/utils"

	beego "github.com/beego/beego/v2/server/web"
	cors "github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.CommentRouterPath = "controllers"
	utils.InitDB()

	database.DB = utils.DB
	models.DB = utils.DB

	// database.CreateTableUser()

	beego.Run()
}
