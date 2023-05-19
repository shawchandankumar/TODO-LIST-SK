package crud

import (
	"models"

	"gorm.io/gorm"
)


type Tasks []models.Task
type Task models.Task


// create
func AddNewTask(db *gorm.DB, taskPayload models.TaskPayload) Task {
	task := Task{Title: taskPayload.Title,
				Todo: taskPayload.Todo,
				Priority: taskPayload.Priority,
				UserId: taskPayload.UserId}

	db.Create(&task)
	return task
}

// read
func GetAllTasks(db *gorm.DB) Tasks {
	var tasks Tasks
	// SELECT * FROM task WHERE name LIKE '%jin%';
	db.Where("user_id = ? AND deleted_at IS NULL", 1).Find(&tasks)
	return tasks
}


// get a task with the id taskId
func GetTask(db *gorm.DB, taskId uint) Task {
	var task Task
	db.Where("id = ?", taskId).Find(&task)
	return task
}

// update
func UpdateTask(db *gorm.DB, taskPayload models.TaskPayload, taskId uint) {
	task := GetTask(db, taskId)
	task.Title = taskPayload.Title
	task.Todo = taskPayload.Todo
	task.Priority = taskPayload.Priority
	db.Save(&task)
}

// delete
func DeleteTask(db *gorm.DB, taskId uint) {
	// DELETE FROM users WHERE id = 10;
	db.Delete(&Task{}, taskId)
}
