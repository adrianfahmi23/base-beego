package main

import (
	_ "example-beego/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

