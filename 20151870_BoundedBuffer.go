package main

import (
	"fmt"
	"time"
	"math/rand"
)

const BUFSIZE int = 5
const MAXDATA int = 100

func producer(buffer chan int, done chan bool) {
	for value := 1; value <= MAXDATA; value++{ 

		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		fmt.Printf("Producer> %d\n", value)
		buffer <- value
		fmt.Printf("Producer< %d\n", value)
	}
	done <- true
}

func consumer(buffer chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		value := <- buffer
		fmt.Printf("Consumer= %d\n", value)
		if value == MAXDATA {
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	buffer := make(chan int, BUFSIZE)
	done := make(chan bool)

	go producer(buffer, done)
	consumer(buffer)
	<-done
}