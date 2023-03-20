package models

import (
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
	d,err := gorm.Open("mysql", "long3112:3112@tcp(127.0.0.1:3306)/todos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
	db.AutoMigrate(&Todo{})
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
	db.Where("id = ?",id).Find(&newTodo)
	return &newTodo,db
}

func GetTodosByName(name string) ([]Todo){
	var newTodo []Todo
	var queryName = "%" + name + "%";
	db.Where("name LIKE ? ",queryName).Find(&newTodo)
	return newTodo
}

