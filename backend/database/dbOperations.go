package database

import "todo-backend/models"

func GetTodo() []models.Todo {
	var todos []models.Todo
	DB.Find(&todos, "status = ?", true)
	return todos 
}

func AddTodo(todo models.Todo) {
	DB.Create(&todo)
}

func UpdateTodo(id uint, content string) {
	DB.Model(&models.Todo{}).Where("id = ?", id).Update("content", content)
}

func DeleteTodo(id uint) {
	DB.Delete(&models.Todo{}, id)
}

func SoftDeleteTodo(id uint) {
	DB.Model(&models.Todo{}).Where("id = ?", id).Update("status", false)
}