package repositories

import (
	"github.com/iqbalmahad/belajar-dari-zoom.git/entities"
	"gorm.io/gorm"
)

type TodoListRepo interface {
	Create(db *gorm.DB, todoList entities.TodoList) entities.TodoList
	Update(db *gorm.DB, todoList entities.TodoList) error
	// Delete(db *gorm.DB, todoList entities.TodoList)
	// GetByID(db *gorm.DB, todoListId int) (entities.TodoList, error)
	// GetAll(db *gorm.DB) []entities.TodoList
}
