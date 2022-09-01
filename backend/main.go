package main

import (
	"fmt"
	"todo-backend/controllers"
	"todo-backend/database"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Init func")
	database.Connect("root:Adisunny123@tcp(127.0.0.1:3306)/todo_go_react?charset=utf8mb4&parseTime=True&loc=Local")
	database.Migrate()
}

func main() {
	fmt.Println("Backend Todo")

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/get", controllers.GetTodo)
		api.POST("/add", controllers.AddTodo)
		api.PUT("/update", controllers.UpdateTodo)
		api.DELETE("/delete", controllers.DeleteTodo)
	}

	return router
}
