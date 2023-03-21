package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB

func Config(){
	d,err := gorm.Open("mysql", "long3112:3112@tcp(127.0.0.1:3306)/todos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db;
}