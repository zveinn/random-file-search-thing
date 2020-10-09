package files

import (
	"log"
	"sync"
	"time"
)

var GlobalWaitGroup = sync.WaitGroup{}

// File ...
type File struct {
	Name    string
	Dir     string
	IsDir   bool
	ModTime time.Time
	Mode    string
	Size    int64
	Results SearchResults
}
type SearchResults struct {
	Hits map[int]string
}

// FullPath ...
func (f *File) FullPath() string {
	return f.Dir + "\\" + f.Name
}

// Print ...
func (f *File) Print() {
	log.Println("FILE INFO:")
	log.Println(f.Name)
}

// Config ...
type Config struct {
	Ignore      []string `json:"ignore"`
	MaxFileSize int64    `json:"maxFileSize"`
}
