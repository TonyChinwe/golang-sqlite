package repository

import (
	"github.com/tonychinwe/libraryone/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:root@tcp(127.0.0.1)/jpa1?parseTime=true"

func InitDb() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.Book{}, &models.Author{}, &models.Subject{}, &models.Genre{}, &models.Level{}, &models.Category{}, &models.Admin{})

}
