package controller

import (
	"net/http"
	"database/sql"
	"log"
	"github.com/b1a9id/todo/src/model"
	"time"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func TaskGET(w http.ResponseWriter, r *http.Request) {
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

	outputJson, err := json.Marshal(&tasks)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(outputJson))

	log.Println(r.URL.Path)
}
