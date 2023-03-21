package controller

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Baguette-Enjoyer/gin-start/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllTodos(ctx *gin.Context){
	var todosArr = []models.Todo{}
	todosArr = models.GetAllTodos()
	todoJS,err := json.Marshal(&todosArr)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(ctx.Writer,"%s \n",todoJS)
}

func AddTodo(ctx *gin.Context){
	var newTodo = models.Todo{}
	json.NewDecoder(ctx.Request.Body).Decode(&newTodo)
	uuid,err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	newTodo.UUId = uuid
	newTodoTemp := newTodo.AddTodo()
	resTodo,err :=json.Marshal(newTodoTemp)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(ctx.Writer,"%s \n",resTodo)
}

func FindTodoByName(ctx *gin.Context){
	name := ctx.Query("name")
	todos := models.GetTodosByName(name)
	res,err := json.Marshal(&todos)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(ctx.Writer,"%s \n", res)
}

func UpdateTodo(ctx *gin.Context){
	id := ctx.Params.ByName("id")
	oid,err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	var todoBody models.Todo
	json.NewDecoder(ctx.Request.Body).Decode(&todoBody)
	newTodo,db := models.GetTodosById(oid);
	newTodo.IsDone = !newTodo.IsDone;
	if todoBody.Name != "" {
		newTodo.Name = todoBody.Name
	}
	if todoBody.IsDone != newTodo.IsDone {
		newTodo.IsDone = todoBody.IsDone
	}
	db.Table("todos").Save(&newTodo)
	res,err := json.Marshal(&newTodo)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(ctx.Writer,"%s \n",res)
}

func DeleteTodo(ctx *gin.Context){
	id := ctx.Params.ByName("id");
	oid,err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	todo:= models.DeleteTodoById(oid)
	res,err := json.Marshal(&todo)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(ctx.Writer,"%s \n",res)
}