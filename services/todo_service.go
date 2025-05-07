package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/gizemcetincaglar/todo_app/models"
)

var todoList []models.Todo

func GetTodosByUser(userID string) []models.Todo {
	var userTodos []models.Todo
	for _, todo := range todoList {
		if todo.UserID == userID && todo.DeletedAt == nil {
			userTodos = append(userTodos, todo)
		}
	}
	return userTodos
}

func GetAllTodos() []models.Todo {
	var activeTodos []models.Todo
	for _, todo := range todoList {
		if todo.DeletedAt == nil {
			activeTodos = append(activeTodos, todo)
		}
	}
	return activeTodos
}


func CreateTodo(userID, name string) models.Todo {
	todo := models.Todo{
		ID:         uuid.New().String(),
		UserID:     userID,
		Name:       name,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Completion: 0.0,
	}
	todoList = append(todoList, todo)
	return todo
}

func DeleteTodo(todoID string) bool {
	for i, todo := range todoList {
		if todo.ID == todoID && todo.DeletedAt == nil {
			now := time.Now()
			todoList[i].DeletedAt = &now
			return true
		}
	}
	return false
}
