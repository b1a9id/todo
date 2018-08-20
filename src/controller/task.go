package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"database/sql"
	"log"
	"github.com/b1a9id/todo/src/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TaskGET(c echo.Context) error {
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

	for result.Next()  {
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
	return c.JSON(http.StatusOK, tasks)
}

func TaskPOST(c echo.Context) error {
	dbDriver := "mysql"
	dbUser := "root"
	dbName := "gwa"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser + "@/" + dbName + dbOption)
	if err != nil {
		log.Fatal(err)
	}

	title := c.FormValue("title")
	now := time.Now()

	task := &model.Task{
		Title: title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err2 := task.Save(db)
	if err2 != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, task)
}
