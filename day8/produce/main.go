package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func produce(ch chan <- string) {
	ch <- "hello1"
	ch <- "hello2"
	ch <- "hello3"

	close(ch)
	waitGroup.Done()
}

func consume(ch <- chan string) {

	for {
		
		str, ok := <- ch
		if !ok {
			fmt.Printf("ch is closed")
			break
		}
		fmt.Printf("value:%s\n", str)
	}
	waitGroup.Done()
}

func main() {
	var ch chan string = make(chan string)
	waitGroup.Add(2)
	go produce(ch)
	go consume(ch)
	waitGroup.Wait()
}