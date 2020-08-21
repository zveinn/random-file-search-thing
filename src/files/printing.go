package files

import "log"

var printBufferMap = make(map[int]chan File)

func InitPrintBuffers(keyword string) {
	for i := 0; i < 1; i++ {
		printBufferMap[i] = make(chan File, 50000)
		go processPrintBuffer(i, keyword)
	}
}

func processPrintBuffer(index int, keyword string) {
	log.Println("Starting print buffer nr:", index)
	var file File
	for {
		file = <-printBufferMap[index]
		log.Println("WE FOUND THE WORD IN FILE:", file.Name)
		for i, v := range file.Results.Hits {
			log.Println("LINE:", i, v)
		}
	}
}
