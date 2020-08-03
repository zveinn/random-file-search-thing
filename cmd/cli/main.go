package main

import (
	"log"

	"github.com/zkynet/golang-lessons-for-beginners/src/tasks"
	p "github.com/zkynet/test-module"
)


func main() {
	log.Println("Launching web server ...")
	p.Print("meow")
	tasks.CreateTask()
}
