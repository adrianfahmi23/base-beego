package controllers

import (
	"encoding/json"
	"example-beego/models"
	"net/http"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// UsersController allows REST-based operations for user models
type UsersController struct {
	beego.Controller
}

// @Title Get All User
// @Description Get Product list
// @Success 200 {object} models.User
// @Param	category_id		query	int	false		"category id"
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router / [get]
func (u *UsersController) Get() {
	u.Data["json"] = models.GetAll()
	u.ServeJSON()
}

// @Title Get User by id
// @Description Get Product list
// @Success 200 {object} models.User
// @Param	id query int "id"
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /:id [get]
func (u *UsersController) GetByID() {
	strId := u.Ctx.Input.Param(":id")
	userId, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		u.serveBadRequest()
	} else {
		user, err := models.GetOne(userId)
		if err != nil {
			u.serveNotFound()
		} else {
			u.Data["json"] = user
			u.ServeJSON()
		}
	}
}

// @Title Store User
// @Description Store User
// @Accept  json
// @Produce  json
// @Param   Authorization  header string  true "Authorization Token"
// @Param   body body models.User "User Value"
// @Success 200 {object} models.User
// @Failure 400 Bad Request
// @Failure 500 Server Error
// @router /store [post]
func (u *UsersController) Post() {
	user := models.User{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error("Could not parse 'Post' request payload, %s.", err)
		u.serveBadRequest()
	} else {
		user.ID = models.AddOne(user)
		u.Data["json"] = user
		u.ServeJSON()
	}
}

func (u *UsersController) Put() {
	user := models.User{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error("Could not parse 'Put' request payload, %s.", err)
		u.serveBadRequest()
	} else {
		err = models.Update(user)
		if err != nil {
			logs.Error("Could not find user in 'Put' request, %s.", err)
			u.serveNotFound()
		} else {
			u.Data["json"] = user
			u.ServeJSON()
		}
	}
}

func (u *UsersController) Delete() {
	strId := u.Ctx.Input.Param(":id")
	userId, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		logs.Error("Could not parse 'Delete' request payload %s.", err)
		u.serveBadRequest()
	} else {
		// For simplicity, we don't block shared objects; simply checking if they exist.
		// In real life application the below code will be a race condition for globally shared
		// objects.
		user, err := models.Delete(userId)
		if err != nil {
			logs.Error("Could not delete user: %s.", err)
			u.serveNotFound()
		} else {
			u.Data["json"] = user
			u.ServeJSON()
		}
	}
}

// Simple wrapper around "request failed due to 400" functionality.
func (u *UsersController) serveBadRequest() {
	u.
		serveResponse("Malformed request data provided, could not continue.",
			http.StatusBadRequest)
}

// Simple wrapper around "request failed not found 404" functionality.
func (u *UsersController) serveNotFound() {
	u.
		serveResponse("Data specified with URL not found.",
			http.StatusNotFound)
}

// Simple wrapper around "request failed due to 500" functionality.
func (u *UsersController) serveInternalServerError() {
	u.serveResponse("Server encountered unexpected error while processing your request.",
		http.StatusInternalServerError)
}

// Carry out real response write, with specified data and response string. For simplicity, only string
// and used only from error handlers.
func (u *UsersController) serveResponse(message string, status int) {
	u.Ctx.ResponseWriter.WriteHeader(status)
	u.Data["json"] = models.RequestError{
		ErrorCode: status,
		Message:   message,
	}

	u.ServeJSON()
}
