package main

import (
	"log"
	"os"

	"github.com/zkynet/golang-lessons-for-beginners/src/files"
	"github.com/zkynet/prettyprint"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("You need to specify the directory you want to start from..")
		os.Exit(1)
	}
	currentDirectory := files.CreateFileMap(os.Args[1])
	sortedSizes := files.SortBySize(currentDirectory)
	prettyprint.PrintFileMapBySize(currentDirectory, sortedSizes)
}
