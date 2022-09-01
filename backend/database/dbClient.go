package database

import (
	"log"
	"todo-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect(connectionString string) {
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")
	}
	log.Println("connected to database")
}

func Migrate() {
	DB.AutoMigrate(&models.Todo{})
	log.Println("database migration completed")
}
