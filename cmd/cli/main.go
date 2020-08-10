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
	fileMap := files.CreateFileMap(os.Args[1])
	// sortedSizes := files.SortBySize(currentDirectory)
	// prettyprint.PrintFileMap(currentDirectory)
	files.WordSearch(fileMap, "meow")
}
