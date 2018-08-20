package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"database/sql"
	"log"
	"github.com/b1a9id/todo/src/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"strconv"
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

func TaskPATCH(c echo.Context) error {
	dbDriver := "mysql"
	dbUser := "root"
	dbName := "gwa"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser + "@/" + dbName + dbOption)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task, err := model.TaskByID(db, uint(id))
	if err != nil {
		panic(err.Error())
	}

	title := c.FormValue("title")
	now := time.Now()

	task.Title = title
	task.UpdatedAt = now

	err2 := task.Update(db)
	if err != nil {
		panic(err2.Error())
	}
	return c.JSON(http.StatusOK, task)
}

func TaskDELETE(c echo.Context) error {
	dbDriver := "mysql"
	dbUser := "root"
	dbName := "gwa"
	dbOption := "?parseTime=true"
	db, err := sql.Open(dbDriver, dbUser + "@/" + dbName + dbOption)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	task, err := model.TaskByID(db, uint(id))
	if err != nil {
		panic(err.Error())
	}

	err = task.Delete(db)
	if err != nil {
		panic(err.Error())
	}

	return c.String(http.StatusOK, "deleted")
}
