package main

import (
	"github.com/b1a9id/todo/src/controller"

	"github.com/gin-gonic/gin"
)


func main()  {
	router := gin.Default()

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
