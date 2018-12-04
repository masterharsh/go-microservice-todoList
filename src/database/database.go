package database

import (
	"HarshCoding/Go-todo-microservice/src/models"

	"github.com/jinzhu/gorm"
)

// Db connection
var Db *gorm.DB

// Init function for migration
func Init() {
	//open a db connection
	var err error
	Db, err = gorm.Open("mysql", "root:Optimus123@/todoDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	Db.AutoMigrate(&models.TodoModel{})
}
