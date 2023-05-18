package main

import (
	"fmt"
	"net/http"

)


func main () {
	fmt.Println("Started the project")
	
	http.ListenAndServe(":8080", nil)
}
