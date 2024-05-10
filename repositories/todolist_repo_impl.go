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

func (repo *TodoListRepoImplementation) Delete(db *gorm.DB, todoList entities.TodoList) {
	db.Delete(&todoList, todoList.ID)
}

func (repo *TodoListRepoImplementation) GetByID(db *gorm.DB, todoListId int) (entities.TodoList, error) {
	var todoList entities.TodoList
	if err := db.Where("id = ?", todoListId).First(&todoList).Error; err != nil {
		return entities.TodoList{}, err
	}
	return todoList, nil

}

func (repo *TodoListRepoImplementation) GetAll(db *gorm.DB) []entities.TodoList {
	var todoLists []entities.TodoList
	db.Find(&todoLists)
	return todoLists
}
