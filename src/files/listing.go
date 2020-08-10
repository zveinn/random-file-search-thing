package files

import (
	"errors"
	"strings"

	"github.com/karrick/godirwalk"
)

func CreateFileMap(dir string) (fileMap map[string]File) {
	fileMap = make(map[string]File)
	_ = godirwalk.Walk(dir, &godirwalk.Options{

		Callback: func(osPathname string, info *godirwalk.Dirent) error {

			if strings.Contains(osPathname, ".git") {
				return errors.New("skipping: " + osPathname)
			}

			// Save the file info.
			fileMap[osPathname] = File{
				Name:  osPathname,
				IsDir: info.IsDir(),
			}

			return nil
		},

		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},

		Unsorted: true,
	})

	return
}
