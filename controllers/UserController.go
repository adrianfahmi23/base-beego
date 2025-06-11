package controllers

import (
	"encoding/json"
	"example-beego/models"
	"example-beego/utils"
	"reflect"

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

	users, err := models.GetAllUser()

	if err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Ada yang salah saat mengambil data user",
		}, 500)
		return
	}

	utils.Response(&res.Controller, utils.ResponseApi[[]models.User]{
		Message: "Berhasil mendapatkan data",
		Data:    users,
	})
}

// @Title Get One User
// @Description Get One User
// @Param   Authorization  header string  true "Authorization Token"
// @Param   username path string  true "Username"
// @Success 200 {object} utils.ResponseApi
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /:id [get]
func (res *UserController) Show() {
	id := res.Ctx.Input.Param(":id")

	user, err := models.GetOneUser(id)

	if err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Ada yang salah mengambil data. " + err.Error(),
		}, 500)
		return
	}

	utils.Response(&res.Controller, utils.ResponseApi[models.User]{
		Message: "Berhasil mendapatkan data",
		Data:    user,
	})
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
		}, 500)
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

// @Title Update User
// @Description Update User
// @Param   Authorization  header string  true "Authorization Token"
// @Param   id path string  true "Id User"
// @Param   body body models.UserForm  true "Data Update User"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /update/:id [put]
func (res *UserController) Update() {
	request := models.UserForm{}
	id := res.Ctx.Input.Param(":id")
	if err := json.Unmarshal(res.Ctx.Input.RequestBody, &request); err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Ada yang salah dengan parsing data. " + err.Error(),
		}, 500)
		return
	}

	user, err := models.GetOneUserById(id)

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "User tidak ditemukan",
		})
		return
	}

	if !utils.CheckPasswordHash(request.Password, user.Password) {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Password salah",
		})
		return
	}

	err = models.UpdateUser(map[string]interface{}{
		"Username": request.Username,
		"Name":     request.Name,
		"Email":    request.Email,
	}, user.ID)

	if err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Ada yang salah saat mengubah data. " + err.Error(),
		}, 500)
		return
	}

	utils.Response(&res.Controller, utils.ResponseMessage{
		Message: "Berhasil mengubah data",
	})
}

// @Title Update Status User
// @Description Update Status User
// @Param   Authorization  header string  true "Authorization Token"
// @Param   id path string  true "Id User"
// @Param   status formData int  true "Status 0 | 1"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /update-status/:id [put]
func (res *UserController) UpdateStatus() {
	id := res.Ctx.Input.Param(":id")
	status, _ := res.GetInt("status")

	err := models.UpdateUser(map[string]interface{}{
		"Status": status,
	}, id)

	if err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Ada yang salah saat mengubah data. " + err.Error(),
		}, 500)
		return
	}

	utils.Response(&res.Controller, utils.ResponseMessage{
		Message: "Berhasil mengubah data",
	})
}

// @Title Delete User
// @Description Delete User
// @Param   id path string  true "Id User"
// @Success 200 {object} utils.ResponseMessage
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /delete/:id [delete]
func (res *UserController) Delete() {
	id := res.Ctx.Input.Param(":id")

	if err := models.DeleteUser(id); err != nil {
		utils.Response(&res.Controller, utils.ResponseMessage{
			Message: "Ada yang salah saat memasukan data. " + err.Error(),
		}, 500)
		return
	}

	utils.Response(&res.Controller, utils.ResponseMessage{
		Message: "Berhasil menghapus data",
	})
}
