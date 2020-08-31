package files

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/karrick/godirwalk"
)

var searchBufferMap = make(map[int]chan File)

func InitSearchBuffers(keyword string) {
	for i := 0; i < 5; i++ {
		searchBufferMap[i] = make(chan File, 5000)
		go processSearchBuffer(i, keyword)
	}
}

func processSearchBuffer(index int, keyword string) {
	// log.Println("Starting search buffer nr:", index)
	for {
		WordSearch(<-searchBufferMap[index], keyword)
	}
}

func WordSearch(v File, keyword string) {

	stat, err := os.Stat(v.Name)
	// if the file is 200mb or bigger, we continue
	if stat.Size() > GlobalConfig.MaxFileSize*1000000 {
		GlobalWaitGroup.Done()
		return
	}

	file, err := os.Open(v.Name)
	if err != nil {
		log.Println(err)
		GlobalWaitGroup.Done()
		return
	}

	scanner := bufio.NewScanner(file)
	var line string
	var foundKeyword bool
	lineNumber := 1
	for scanner.Scan() {
		line = scanner.Text()
		if strings.Contains(line, keyword) {
			foundKeyword = true
			v.Results.Hits[lineNumber] = line
		}
		lineNumber++
	}
	file.Close()

	if foundKeyword {
		printBufferMap[rand.Intn(len(printBufferMap))] <- v
	} else {
		GlobalWaitGroup.Done()
	}
}

func WalkDirectories(dir string) {
	_ = godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(osPathname string, info *godirwalk.Dirent) error {

			for _, v := range GlobalConfig.Ignore {
				if strings.Contains(osPathname, v) {
					return godirwalk.SkipThis
				}
			}

			if !info.IsDir() {
				GlobalWaitGroup.Add(1)
				searchBufferMap[rand.Intn(len(searchBufferMap))] <- File{
					Name:  osPathname,
					IsDir: info.IsDir(),
					Results: SearchResults{
						Hits: make(map[int]string),
					},
				}
			}

			return nil
		},
		Unsorted: true,
	})

	return
}
