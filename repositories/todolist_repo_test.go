package repositories

import (
	"testing"

	"github.com/iqbalmahad/belajar-dari-zoom.git/database"
	"github.com/iqbalmahad/belajar-dari-zoom.git/entities"
	"github.com/stretchr/testify/assert"
)

func TestTodoListRepoCreate(t *testing.T) {
	todoList1 := entities.TodoList{
		Title:       "Hobbies",
		Description: "Suka ngoding",
		Completed:   false,
	}

	db := database.DBConnection()
	NewTodoListRepo().Create(db, todoList1)
	assert.NotEmpty(t, todoList1, 1)
}

func TestTodoListRepoUpdate(t *testing.T) {
	var todoList1 entities.TodoList
	db := database.DBConnection()
	db.Where("id = ?", 2).First(&todoList1)
	todoList1.Title = "Hobbies"
	todoList1.Description = "suka ngoding"
	todoList1.Completed = true
	NewTodoListRepo().Update(db, todoList1)
}

func TestTodoListRepoDelete(t *testing.T) {
	var todoList1 entities.TodoList
	db := database.DBConnection()
	db.First(&todoList1, 2)
	NewTodoListRepo().Delete(db, todoList1)
}

func TestTodoListRepoGetById(t *testing.T) {
	var todoList1 entities.TodoList
	db := database.DBConnection()
	db.Where("id = ?", 2).First(&todoList1)
	NewTodoListRepo().GetByID(db, todoList1.ID)
}

func TestTodoListRepoGetAll(t *testing.T) {
	var todoList entities.TodoList
	db := database.DBConnection()
	db.Find(&todoList)
	NewTodoListRepo().GetAll(db)
}
