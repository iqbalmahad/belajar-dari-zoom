package repositories

import (
	"github.com/iqbalmahad/belajar-dari-zoom.git/entities"
	"gorm.io/gorm"
)

type TodoListRepoImplementation struct {
}

func NewTodoListRepo() TodoListRepo {
	return &TodoListRepoImplementation{}
}

func (repo *TodoListRepoImplementation) Create(db *gorm.DB, todoList entities.TodoList) entities.TodoList {
	db.Create(&todoList)
	return todoList
}

func (repo *TodoListRepoImplementation) Update(db *gorm.DB, todoList entities.TodoList) error {
	return db.Model(&todoList).Updates(todoList).Error
}

// func (repo *TodoListRepoImplementation) Delete(db *gorm.DB, todoList entities.TodoList) {

// }

// func (repo *TodoListRepoImplementation) GetByID(db *gorm.DB, todoListId int) (entities.TodoList, error) {
// 	return
// }
// func (repo *TodoListRepoImplementation) GetAll(db *gorm.DB) []entities.TodoList {

// }
