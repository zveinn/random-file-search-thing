package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/zkynet/golang-lessons-for-beginners/src/files"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("You need to specify the directory you want to start from..")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		log.Println("You need to specify a search word ...")
		os.Exit(1)
	}

	files.LoadConfig()
	rand.Seed(time.Now().UTC().UnixNano())
	files.InitSearchBuffers(os.Args[2])
	files.InitPrintBuffers(os.Args[2])

	files.WalkDirectories(os.Args[1])
	for {
		time.Sleep(10 * time.Second)
	}
	// meow
	// meow
	// meow
	// meow
	// meow
}
