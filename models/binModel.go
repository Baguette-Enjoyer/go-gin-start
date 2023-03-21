package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type TodoBin struct {
	gorm.Model
	Name 	string 		`json:"name"`
}

// func init(){
// 	db.AutoMigrate(&TodoBin{})
// }

func GetAllTodosInBin() []TodoBin{
	items := []TodoBin{}
	db.Table("todobins").Find(&items)
	return items
}

func (todo *Todo) AddTodoToBin() *TodoBin {
	newBinItem := TodoBin{}
	newBinItem.Name = todo.Name
	db.Table("todobins").Create(&newBinItem)
	return &newBinItem
}

func CleanBin() {
	db.Table("todobins").Delete(&TodoBin{})
}
