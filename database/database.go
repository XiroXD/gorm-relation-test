package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/XiroXD/gorm-test/models"
)

var DBConn *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DBConn = db

	return db
}

func Migrate(conn *gorm.DB) {
	conn.AutoMigrate(&models.Post{}, &models.Author{})
}
