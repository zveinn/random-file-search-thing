package files

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CreateFileMap ...
// This function parses the directory given as input and
// returns a map of files, the index being the full
// file path and the value being the file ext.
func CreateFileMap(dir string) map[string]File {
	fileMap := make(map[string]File)
	fullPath, pathError := os.Getwd()
	if pathError != nil {
		panic(pathError)
	}
	walkError := filepath.Walk(dir, func(path string, info os.FileInfo, innerWalkError error) error {
		if innerWalkError != nil {
			return innerWalkError
		}
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}
		fileMap[info.Name()] = File{
			Name:    info.Name(),
			Size:    info.Size(),
			IsDir:   info.IsDir(),
			ModTime: info.ModTime(),
			Mode:    info.Mode().String(),
			Dir:     fullPath,
		}
		return nil
	})
	if walkError != nil {
		log.Println(walkError.Error())
		os.Exit(1)
	}

	// filter out all .git files.
	FilterOnName(fileMap, func(name string) bool {
		return strings.Contains(name, ".git")
	})

	return fileMap
}
