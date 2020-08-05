package main

import (
	"os"
	"path/filepath"

	"github.com/zkynet/prettyprint"
)

func main() {
	currentDirectory := CreateFileMap(".")
	prettyprint.PrintFileMap(currentDirectory)
}

// CreateFileMap ...
// This function parses the directory given as input and
// returns a map of files, the index being the full
// file path and the value being the file ext.
func CreateFileMap(dir string) map[string]os.FileInfo {
	fileMap := make(map[string]os.FileInfo)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		fileMap[dir+"/"+path] = info
		return nil
	})
	if err != nil {
		panic(err)
	}

	return fileMap
}
