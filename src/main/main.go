package main

import (
	"github.com/gorilla/mux"
	"github.com/b1a9id/todo/src/controller"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	s := router.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/tasks", controller.TaskGET).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
