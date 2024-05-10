package repositories

import (
	"testing"

	"github.com/iqbalmahad/belajar-dari-zoom.git/database"
	"github.com/iqbalmahad/belajar-dari-zoom.git/entities"
)

func TestTodoListRepoCreate(t *testing.T) {
	todoList1 := entities.TodoList{
		Title:       "Hobbies",
		Description: "Suka belajar coding",
		Completed:   false,
	}

	db := database.DBConnection()
	NewTodoListRepo().Create(db, todoList1)
}

func TestTodoListRepoUpdate(t *testing.T) {
	var todoList1 entities.TodoList
	db := database.DBConnection()
	db.First(&todoList1, 1)
	todoList1.Title = "Hobbies"
	todoList1.Description = "suka ngoding dan suka main volly"
	todoList1.Completed = true
	NewTodoListRepo().Update(db, todoList1)
}
