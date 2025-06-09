package utils

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	host := web.AppConfig.DefaultString("db_host", "localhost")
	database := web.AppConfig.DefaultString("db_name", "database")
	user := web.AppConfig.DefaultString("db_user", "root")
	pass := web.AppConfig.DefaultString("db_pass", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
