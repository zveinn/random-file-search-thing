package main

import (
	"errors"
	"strings"

	"github.com/karrick/godirwalk"
)

func main() {
	_ = godirwalk.Walk("./../..", &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if strings.Contains(osPathname, ".git") {
				return errors.New("skipping: " + osPathname)
			}
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			if err.Error() == "path filter" {
				return godirwalk.SkipNode
			}
			return godirwalk.Halt
		},
	})
}
