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

func InitSearchBuffers(keywords string) {
	for i := 0; i < 5; i++ {
		searchBufferMap[i] = make(chan File, 5000)
		go processSearchBuffer(i, keywords)
	}
}

func processSearchBuffer(index int, keywords string) {
	// log.Println("Starting search buffer nr:", index)
	for {
		WordSearch(<-searchBufferMap[index], keywords)
	}
}

func WordSearch(v File, keywords string) {

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

	wordList := strings.Split(keywords, ",")
	scanner := bufio.NewScanner(file)
	var line string
	var foundKeyword bool
	lineNumber := 1
	for scanner.Scan() {
		line = scanner.Text()
		for _, keyword := range wordList {
			if strings.Contains(line, keyword) {
				foundKeyword = true
				v.Results.Hits[lineNumber] = line
			}
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
