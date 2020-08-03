package main

import (
	"log"

	"github.com/zkynetio/classes/l2/src/tasks"
)


func main() {
	log.Println("Launching web server ...")
	
	tasks.CreateTask()
}
