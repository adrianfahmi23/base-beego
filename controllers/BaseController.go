package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Response(data any, code ...int) {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}

	c.Ctx.Output.SetStatus(statusCode)
	c.Data["json"] = data
	c.ServeJSON()
}
