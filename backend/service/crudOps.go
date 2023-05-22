package service

import (
	"gorm.io/gorm"

	"todo-task/configuration"
	"todo-task/models"
)

type Tasks []models.Task
type Task models.Task

var db *gorm.DB = *(configuration.GetDbConnection())

// create
func AddNewTask(taskPayload models.TaskPayload) Task {
	task := Task{Title: taskPayload.Title,
		Todo:     taskPayload.Todo,
		Priority: taskPayload.Priority,
		UserId:   taskPayload.UserId}

	db.Create(&task)
	return task
}

// read
func GetAllTasks() Tasks {
	var tasks Tasks
	// SELECT * FROM task WHERE name LIKE '%jin%';
	db.Where("user_id = ? AND deleted_at IS NULL", 1).Find(&tasks)
	return tasks
}

// get a task with the id taskId
func GetTask(taskId uint) Task {
	var task Task
	db.Where("id = ?", taskId).Find(&task)
	return task
}

// update
func UpdateTask(taskPayload models.TaskPayload, taskId uint) Task {
	task := GetTask(taskId)
	task.Title = taskPayload.Title
	task.Todo = taskPayload.Todo
	task.Priority = taskPayload.Priority
	db.Save(&task)
	return task
}

// delete
func DeleteTask(taskId uint) {
	// DELETE FROM users WHERE id = 10;
	db.Delete(&Task{}, taskId)
}
