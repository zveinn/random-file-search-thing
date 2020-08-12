package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	BuildWindows()
	// BuildMac()
	// BuildLinux()
}

func BuildWindows() {
	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "amd64")
	c := exec.Command("go", "build", "-o", "find.exe", "../cmd/cli/main.go")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func BuildMac() {
	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", "amd64")
	c := exec.Command("go", "build", "-o", "find.mac", "../cmd/cli/main.go")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
func BuildLinux() {
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")
	os.Setenv("CGO_ENABLED", "0")
	c := exec.Command("go", "build", "-o", "find.linux", "../cmd/cli/main.go")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
