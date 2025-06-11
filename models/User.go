package models

import (
	"example-beego/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    uint      `json:"status"`
}

type UserForm struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

func GetOneUserById(id string) (User, error) {
	var user User

	if err := DB.First(&user, "id = ?", id).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func GetOneUser(username string) (User, error) {
	var user User

	if err := DB.Model(User{Username: username}).First(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}

func GetAllUser() ([]User, error) {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return []User{}, err
	}

	return users, nil
}

func StoreUser(user UserForm) error {
	hash, _ := utils.HashPassword(user.Password)
	if err := DB.Omit("UpdatedAt", "Status", "CreatedAt").Create(&User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: hash,
	}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(update map[string]interface{}, id string) error {

	if err := DB.Model(&User{}).Where("id = ?", id).Updates(update).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {

	if err := DB.Delete(User{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
