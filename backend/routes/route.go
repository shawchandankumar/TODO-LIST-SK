package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"gorm.io/gorm"

	"configuration"
	"crud"
	"models"
)

var dB *gorm.DB = nil
var TaskPayload models.TaskPayload

func StartServer() {

	r := mux.NewRouter()

    r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Welcome</h1>")
	})

	// configuring the database
	configuration.DbConfig()
	dB = *(configuration.GetDbConnection())

	r.HandleFunc("/tasks", getAllTasks).Methods("GET")
	r.HandleFunc("/tasks", createNewTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
    
    http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(r))
}

func extractTaskId (r *http.Request) uint {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])
	
	if err != nil {
		panic(err)
	}

	return uint(taskId)
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(crud.GetAllTasks(dB))
}

func createNewTask(w http.ResponseWriter, r *http.Request) {
	
	err := json.NewDecoder(r.Body).Decode(&TaskPayload)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(crud.AddNewTask(dB, TaskPayload))
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&TaskPayload)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	
	taskId := extractTaskId(r)
	
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(crud.UpdateTask(dB, TaskPayload, uint(taskId)))
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := extractTaskId(r)
	crud.DeleteTask(dB, uint(taskId))
}
