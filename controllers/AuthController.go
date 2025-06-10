package controllers

import (
	"encoding/json"
	"example-beego/models"
	"example-beego/utils"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

// @Title Login
// @Description Login
// @Accept  json
// @Param   body body models.Auth true "Auth Login"
// @Success 200 {object} models.Auth
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /login [post]
func (res *AuthController) Login() {
	request := models.Auth{}

	if err := json.Unmarshal(res.Ctx.Input.RequestBody, &request); err != nil {
		res.Ctx.Output.SetStatus(500)
		res.Data["json"] = err.Error()
		res.ServeJSON()
	}

	val, err := utils.GenerateJWT(models.User{
		ID:      1,
		Name:    "Fahmi Adrian",
		Country: "Indonesia",
		Email:   "adrian.fahmi23@gmail.com",
	})

	if err != nil {
		res.Ctx.Output.SetStatus(500)
		res.Data["json"] = err.Error()
		res.ServeJSON()
	}

	res.Data["json"] = val
	res.ServeJSON()

}

// @Title Check Login
// @Description Check Login
// @Param   Authorization  header string  true "Authorization Token"
// @Success 200 {object} models.Auth
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /check-login [post]
func (res *AuthController) CheckToken() {
	authorizationHeader := res.Ctx.Request.Header.Get("Authorization")

	if !strings.Contains(authorizationHeader, "Bearer") {
		res.Ctx.Output.SetStatus(401)
		res.Data["json"] = "Token tidak ditemukan"
		res.ServeJSON()
	}

	token := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	claims, err := utils.VerifyJWT(token)

	if err != nil {
		res.Ctx.Output.SetStatus(401)
		// res.Data["json"] = "Token signature tidak cocok"
		res.ServeJSON()
	}

	res.Data["json"] = claims
	res.ServeJSON()
}
