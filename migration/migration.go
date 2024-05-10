package migration

import (
	"github.com/iqbalmahad/belajar-dari-zoom.git/database"
	"github.com/iqbalmahad/belajar-dari-zoom.git/entities"
	"gorm.io/gorm"
)

func DBMigration() *gorm.DB {
	db := database.DBConnection()

	todoList := entities.TodoList{}
	db.AutoMigrate(&todoList)

	return db
}
