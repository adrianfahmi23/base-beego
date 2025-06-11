package controllers

import (
	"encoding/json"
	"example-beego/models"
	"example-beego/utils"
	"log"

	beego "github.com/beego/beego/v2/server/web"
)

// UserController allows REST-based operations for user models
type UserController struct {
	beego.Controller
}

// @Title List Users
// @Description List Users
// @Param   Authorization  header string  true "Authorization Token"
// @Success 200 {object} utils.ResponseApi
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router / [get]
func (res *UserController) Index() {
	utils.Response(&res.Controller, "Berhasil Mengambil data")
}

// @Title Store User
// @Description Post User
// @Param   Authorization  header string  true "Authorization Token"
// @Param   body body models.UserForm  true "Data User"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /store [post]
func (res *UserController) Store() {
	request := models.UserForm{}

	if err := json.Unmarshal(res.Ctx.Input.RequestBody, &request); err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Ada yang salah dengan parsing data. " + err.Error(),
		}, 400)
		return
	}

	if err := models.StoreUser(request); err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Ada yang salah saat memasukan data. " + err.Error(),
		}, 500)
		return
	}

	utils.Response(&res.Controller, utils.ResponseMessage{
		Message: "Berhasil menambahkan data",
	})
}

func (res *UserController) Delete() {
	id := res.Ctx.Input.Param(":id")

	log.Println(id)
}
