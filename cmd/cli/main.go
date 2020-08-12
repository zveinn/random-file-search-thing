package main

import (
	"log"
	"os"

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

	fileMap := files.CreateFileMap(os.Args[1])
	// sortedSizes := files.SortBySize(currentDirectory)
	// prettyprint.PrintFileMap(currentDirectory)
	files.WordSearch(fileMap, os.Args[2])
	// meow
}
