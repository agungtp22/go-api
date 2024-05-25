package main

import (
	"github.com/gin-gonic/gin"
	"go_restfull/controller"
	"go_restfull/model"
)

func main() {
	r := gin.Default()

	model.ConnectDb()

	r.GET("/api/news", controller.GetNews)
	r.GET("/api/news/:id", controller.GetDetailNews)
	r.POST("/api/news", controller.CreateNews)
	r.PUT("/api/news/:id", controller.UpdateNews)
	r.DELETE("/api/news", controller.DeleteNews)

	r.Run()
}
