package main

import (
	"github.com/labstack/echo"
	"github.com/b1a9id/todo/src/controller"
)
func main() {
	e := echo.New()
	g := e.Group("/api/v1")

	g.GET("/tasks", controller.TaskGET)
	g.POST("/tasks", controller.TaskPOST)

	e.Logger.Fatal(e.Start(":8080"))
}
