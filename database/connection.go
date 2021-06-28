package database

import (
	"github.com/cryonayes/StajProje/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Connect() {

	dsn := "cryonayes:mysql3366@tcp(localhost:3306)/go_share?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Database connection failed!")
	}
	DBConn = connection
	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic("Database migration error!")
	}
}
