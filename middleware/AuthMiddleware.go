package middleware

import (
	"example-beego/utils"
	"strings"

	"github.com/beego/beego/v2/server/web/context"
)

func AuthMiddleware(ctx *context.Context) {
	authorizationHeader := ctx.Request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(utils.ResponseMessage{
			Message: "Unauthorized",
		}, false, false)
		return
	}

	token := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	_, err := utils.VerifyJWT(token)

	if err != nil {
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(utils.ResponseMessage{
			Message: "Unauthorized",
		}, false, false)
	}
}
