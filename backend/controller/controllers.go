package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"todo-task/models"
	"todo-task/service"
)

var TaskPayload models.TaskPayload

func extractTaskId(r *http.Request) uint {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	if err != nil {
		panic(err)
	}

	return uint(taskId)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.GetAllTasks())
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&TaskPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.AddNewTask(TaskPayload))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&TaskPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskId := extractTaskId(r)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.UpdateTask(TaskPayload, uint(taskId)))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := extractTaskId(r)
	service.DeleteTask(uint(taskId))
}
