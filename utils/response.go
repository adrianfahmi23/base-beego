package utils

import (
	"github.com/beego/beego/v2/server/web"
)

type ResponseApi[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
}

type ResponseMeta[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data"`
	Meta    Meta   `json:"meta"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func Response[T any](ctx *web.Controller, data T, code ...int) {
	statusCode := 200
	if len(code) > 0 {
		statusCode = code[0]
	}

	ctx.Ctx.Output.SetStatus(statusCode)
	ctx.Data["json"] = data
	ctx.ServeJSON()
}
