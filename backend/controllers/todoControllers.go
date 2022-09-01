package controllers

import (
	"net/http"
	"todo-backend/models"
	"todo-backend/services"

	"github.com/gin-gonic/gin"
)

func GetTodo(ctx *gin.Context) {
	todos := services.GetTodo()
	ctx.IndentedJSON(http.StatusFound, todos)
}

func AddTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		panic("error creating todo")
	}
	services.AddTodo(todo)
	ctx.IndentedJSON(http.StatusCreated, todo)
}

func UpdateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		panic("error creating todo")
	}
	services.UpdateTodo(todo.ID, todo.Content)
	ctx.IndentedJSON(http.StatusCreated, todo)
}

func DeleteTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		panic("error creating todo")
	}
	services.SoftDeleteTodo(todo.ID)
	ctx.IndentedJSON(http.StatusAccepted, todo)
}