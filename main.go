package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Baguette-Enjoyer/gin-start/controller"
)



func main() {
	server := gin.Default()
	server.GET("/getAllTodos", controller.GetAllTodos)
	server.GET("/getTodosByName",controller.FindTodoByName)
	server.GET("/getTodosBin",controller.GetAllTodosInBin)
	server.POST("/addTodo",controller.AddTodo)
	server.PUT("/updateTodo/:id",controller.UpdateTodo)
	server.DELETE("/deleteTodo/:id",controller.DeleteTodo)
	server.DELETE("/cleanBin",controller.DeleteBin)
	server.Run(":3000")
}
