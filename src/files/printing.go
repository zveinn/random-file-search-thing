package files

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var printBufferMap = make(map[int]chan File)

func InitPrintBuffers(keyword string) {
	for i := 0; i < 1; i++ {
		printBufferMap[i] = make(chan File, 50000)
		go processPrintBuffer(i, keyword)
	}
}

func processPrintBuffer(index int, keyword string) {
	// log.Println("Starting print buffer nr:", index)
	var file File
	for {
		file = <-printBufferMap[index]
		// log.Println("WE FOUND THE WORD IN FILE:", file.Name)
		color.Green("FILE: " + file.Name)
		for i, v := range file.Results.Hits {
			v = strings.Replace(v, keyword, color.YellowString(keyword), -1)
			fmt.Println(color.GreenString("("+strconv.Itoa(i)+"): ") + v)
		}
		GlobalWaitGroup.Done()
	}
}
