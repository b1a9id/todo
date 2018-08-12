package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/b1a9id/todo/src/model"
	"database/sql"
	"time"
	"fmt"
	"net/http"
	"log"
)

func TasksGET(c *gin.Context)  {
	dbDriver := "mysql"
	dbUser := "root"
	dbName := "gwa"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser + "@/" + dbName + dbOption)
	if err != nil {
		log.Fatal(err)
	}
	
	result, err := db.Query("SELECT * FROM task ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tasks := []model.Task{}

	for result.Next() {
		task := model.Task{}
		var id uint
		var createdAt, updatedAt time.Time
		var title string

		err = result.Scan(&id, &createdAt, &updatedAt, &title)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt
		task.Title = title
		tasks = append(tasks, task)
	}
	fmt.Println(tasks)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
