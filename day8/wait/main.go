package main

import (
	"time"
	"fmt"
	"sync"
)



func calc(index int, waitGroup *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
	}

	waitGroup.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	start := time.Now().UnixNano()
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go calc(i, &waitGroup)
	}

	waitGroup.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("finished, cost:%d ms\n", (end - start)/1000/1000)
}