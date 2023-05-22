package configuration

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"todo-task/models"
)

var DB **gorm.DB = nil

func Init() {
	dsn := "root:root@tcp(127.0.0.1:3307)/todo-list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("There is some error connecting to database")
		return
	} else {
		fmt.Println("Successfully connected to the database")
	}

	user1 := models.User{}
	task1 := models.Task{}

	db.AutoMigrate(&user1, &task1)
	DB = &db
}

func GetDbConnection() **gorm.DB {
	Init()
	return DB
}
