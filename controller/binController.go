package controller

import (
	"encoding/json"
	"fmt"

	"github.com/Baguette-Enjoyer/gin-start/models"
	"github.com/gin-gonic/gin"
)

func GetAllTodosInBin(c *gin.Context){
	todosbinarr := models.GetAllTodosInBin()
	res,err := json.Marshal(&todosbinarr)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(c.Writer,"%s \n",res)
}

func DeleteBin(c *gin.Context){
	models.CleanBin()
	emptyArr := []models.TodoBin{}
	res,err := json.Marshal(&emptyArr)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(c.Writer,"%s \n",res)
}