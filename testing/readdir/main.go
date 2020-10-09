package main

import (
	"io/ioutil"
	"log"
)

func main() {
	files, err := ioutil.ReadDir("C:\\Users\\Notandi\\go\\src\\github.com\\zveinn\\golang-lessons-for-beginners\\")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(files)
}
