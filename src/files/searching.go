package files

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func WordSearch(fileMap map[string]File, keyword string) (filteredMap map[string]File) {
	filteredMap = make(map[string]File)
	// abs, err := filepath.Abs(osPathname)
	// if err == nil {
	// 	fmt.Println("Absolute:", abs)
	// }
	for _, v := range fileMap {

		stat, err := os.Stat(v.Name)
		// log.Println(stat, err)
		// if the file is 200mb or bigger, we continue
		if stat.Size() > GlobalConfig.MaxFileSize*1000000 {
			continue
		}

		// Open the file
		file, err := os.Open(v.Name)
		if err != nil {
			log.Println(err)
			continue
		}

		scanner := bufio.NewScanner(file)

		var line string
		for scanner.Scan() {
			line = scanner.Text()
			if strings.Contains(line, keyword) {
				log.Println("WE FOUND THE WORD IN FILE:", v.Name)
				log.Println("LINE:", line)
			}
		}
		file.Close()
	}
	return
}
