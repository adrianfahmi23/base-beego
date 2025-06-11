package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    uint      `json:"status"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}

func GetOne(username string) (Auth, error) {
	var user Auth

	if err := DB.First(&user, username).Error; err != nil {
		return Auth{}, err
	}

	return user, nil
}

func GetAll() ([]Auth, error) {
	var users []Auth

	if err := DB.Find(&users).Error; err != nil {
		return []Auth{}, err
	}

	return users, nil
}

func Store(user User) error {

	if err := DB.Omit("UpdatedAt", "Status", "CreatedAt").Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func Update(update map[string]interface{}, id string) error {

	if err := DB.Model(&User{}).Where("id = ?", id).Updates(update).Error; err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {

	if err := DB.Delete(User{}, id).Error; err != nil {
		return err
	}

	return nil
}
