package main

import (
	"bubble/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	//RESTful风格
	r.GET("/v1/todo", controller.Gettodo)           //获取待办事项
	r.POST("/v1/todo", controller.Addtodo)          //添加待办事项
	r.PUT("/v1/todo/:id", controller.Edittodo)      //修改待办事项
	r.DELETE("/v1/todo/:id", controller.Deletetodo) //删除待办事项
	return r
}
