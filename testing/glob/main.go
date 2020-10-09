package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("C:\\Users\\Notandi\\go\\src\\github.com\\zveinn\\golang-lessons-for-beginners\\*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files) // contains a list of all files in the current directory
}
