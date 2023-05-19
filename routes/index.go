package main

import (
	"fmt"

	"gorm.io/gorm"

	"configuration"
	"crud"
	"models"
)


func main () {
	fmt.Println("Started the project")

	// configuring the database
	configuration.DbConfig()
	// var task models.Task

	var dB *gorm.DB = *(configuration.GetDbConnection())
	

	// fmt.Printf("%+v", crud.GetAllTasks(dB))
	// crud.DeleteTask(dB, 5)
	crud.UpdateTask(dB, models.TaskPayload{Title: "Sleep", 
											Todo: "Going to Sleep for Better Health",
											Priority: 2}, 2)				
	
	// http.ListenAndServe(":8080", nil)
}
