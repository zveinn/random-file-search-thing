package main

import (
	"fmt"
	"log"
	"time"
)

var buffer1 = make(chan int, 31)
var buffer2 = make(chan int, 31)

func main() {
	fmt.Println(buffer1, len(buffer1))
	for i := 0; i <= 10; i++ {
		go putInbuffer1(i)
	}

	fmt.Println(buffer1, len(buffer1))
	go process()
	go process2()
	for {
		time.Sleep(30 * time.Second)
	}
}

func putInbuffer1(x int) {
	select {
	case buffer1 <- x:
	default:
		fmt.Println("no message sent")
	}

}

func process() {
	for {
		item := <-buffer1
		log.Println("BUFFER1:", item)
		select {
		case buffer2 <- item:
		default:
			fmt.Println("no message sent")
		}
	}
}

func process2() {
	for {
		item := <-buffer2
		log.Println("BUFFER2:", item)
	}
}
