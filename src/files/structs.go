package files

import (
	"log"
	"time"
)

// File ...
type File struct {
	Name    string
	Dir     string
	IsDir   bool
	ModTime time.Time
	Mode    string
	Size    int64
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
