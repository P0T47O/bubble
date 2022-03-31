package controller

import (
	"bubble/database"
	"bubble/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Gettodo(c *gin.Context) {
	DB := database.GetDB()
	var todolist []model.Todo
	err := DB.Find(&todolist).Error //查找表中所有数据
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}
func Addtodo(c *gin.Context) {
	DB := database.GetDB()
	//获取参数
	var todo model.Todo
	c.BindJSON(&todo)
	err := DB.Create(&todo).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
func Edittodo(c *gin.Context) {
	DB := database.GetDB()
	id, _ := c.Params.Get("id")
	err := DB.Model(model.Todo{}).Where("id=?", id).Update("status", true).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, nil)
	}

}
func Deletetodo(c *gin.Context) {
	DB := database.GetDB()
	id, _ := c.Params.Get("id")
	err := DB.Where("id=?", id).Delete(&model.Todo{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

}
