package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *MainController) Get() {
	user := User{
		ID:        1,
		Name:      "Rey",
		Email:     "rey@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.Data["json"] = map[string]interface{}{
		"code":    200,
		"message": "Success",
		"data":    user,
	}
	c.ServeJSON()
}
