package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"todo-task/controller"
)

func StartServer() {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Welcome</h1>")
	})

	r.HandleFunc("/tasks", controller.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks", controller.CreateNewTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", controller.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")

	http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(r))
}