package controllers

import (
	"example-beego/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// UsersController allows REST-based operations for user models
type UsersController struct {
	beego.Controller
}

// @Title List Users
// @Description List Users
// @Param   Authorization  header string  true "Authorization Token"
// @Success 200 {object} utils.ResponseApi
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router / [get]
func (res *UsersController) Index() {
	utils.Response(&res.Controller, "Berhasil Mengambil data")
}
