package main

import (
	"github.com/b1a9id/todo/src/controller"

	"github.com/gin-gonic/gin"
)


func main()  {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/tasks", controller.TasksGET)
		v1.POST("/tasks", controller.TaskPOST)
		v1.PATCH("/tasks/:id", controller.TaskPATCH)
	}

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
