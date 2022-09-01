package services

import (
	"todo-backend/database"
	"todo-backend/models"
)

func GetTodo() []models.Todo {
	return database.GetTodo()
}

func AddTodo(t models.Todo) {
	todo := &models.Todo{
		Content:  t.Content,
		Status: true,
	}
	database.AddTodo(*todo)
}

func UpdateTodo(id uint, content string) {
	database.UpdateTodo(id, content)
}

func DeleteTodo(id uint) {
	database.DeleteTodo(id)
}

func SoftDeleteTodo(id uint) {
	database.SoftDeleteTodo(id)
}
