package models

import (
	"github.com/Baguette-Enjoyer/gin-start/database"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

type Todo struct {
	gorm.Model
	UUId 	uuid.UUID 	`json:"uuid:"`
	Name 	string 	`json:"name"`
	IsDone  bool    `json:"isDone"`
}

func init(){
	database.Config()
	db = database.GetDB()
	db.AutoMigrate(&Todo{})
	db.Table("todobins").AutoMigrate(&TodoBin{})
}
func (todo *Todo) AddTodo()  *Todo {
	db.Create(&todo)
	return todo
}
func GetAllTodos() []Todo{
	todosArr := []Todo{}
	db.Find(&todosArr)
	return todosArr	
}

func GetTodosById(id int) (*Todo,*gorm.DB){
	var newTodo Todo
	db.Table("todos").Where("id = ?",id).Find(&newTodo)
	return &newTodo,db
}

func GetTodosByName(name string) ([]Todo){
	var newTodo []Todo
	var queryName = "%" + name + "%";
	db.Table("todos").Where("name LIKE ? ",queryName).Find(&newTodo)
	return newTodo
}
func DeleteTodoById(id int) *Todo {
	var deleteTodo Todo
	var newTodo Todo
	db.Table("todos").Where("id = ?",id).Find(&newTodo).Delete(&deleteTodo)
	go func (){
		newBinItem := TodoBin{}
		newBinItem.Name = newTodo.Name
		db.Table("todobins").Create(&newBinItem)
	}()
	return &deleteTodo
}

