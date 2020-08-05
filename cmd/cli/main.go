package main

import (
	"github.com/zkynet/golang-lessons-for-beginners/src/files"
	"github.com/zkynet/prettyprint"
)

func main() {
	currentDirectory := files.CreateFileMap(".")
	prettyprint.PrintFileMap(currentDirectory)
}
