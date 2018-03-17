package main

import (
	"fmt"
	"sync"
	"os/signal"
	"os"
	"syscall"
)

var waitGroup sync.WaitGroup

func produce(ch chan <- string, exitChan chan bool) {

	var i int
	var exit bool
	for {
		str := fmt.Sprintf("hello %d", i)
		select {
		case ch <- str:
		case exit = <- exitChan:
		}
		if exit {
			fmt.Printf("user notify produce exited\n")
			break
		}
	}
	close(ch)
	waitGroup.Done()
}

func consume(ch <-chan string) {

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
	var exitChan chan bool = make (chan bool, 1)
	var sinalChan chan os.Signal = make(chan os.Signal, 1)
	waitGroup.Add(2)
	signal.Notify(sinalChan, syscall.SIGUSR2)

	go produce(ch, exitChan)
	go consume(ch)

	<- sinalChan
	exitChan <- true
	waitGroup.Wait()
}