package controllers

import (
	"encoding/json"
	"example-beego/models"
	"example-beego/utils"
	"strconv"
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
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Gagal mendapatkan data",
		}, 401)
		return
	}

	val, err := utils.GenerateJWT(models.User{
		ID:    1,
		Name:  "Fahmi Adrian",
		Email: "adrian.fahmi23@gmail.com",
	})

	if err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Gagal Generate Token",
		}, 401)
		return
	}

	utils.Response(&res.Controller, utils.ResponseApi[map[string]any]{
		Message: "Berhasil login",
		Data: map[string]any{
			"token": val,
		},
	}, 200)
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
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Token tidak ditemukan",
		}, 401)
		return
	}

	token := strings.Replace(authorizationHeader, "Bearer ", "", -1)

	claims, err := utils.VerifyJWT(token)

	if err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Token signature tidak cocok",
		}, 401)
		return
	}

	idInt64, _ := strconv.Atoi(claims.ID) // You should handle the error in production code
	utils.Response(&res.Controller, utils.ResponseApi[models.User]{
		Message: "Berhasil mengambil data",
		Data: models.User{
			ID:    idInt64,
			Name:  claims.Username,
			Email: claims.Email,
		},
	}, 200)
}
