package main

import (
	"time"
	"fmt"
)

var exits [3]bool

func calc(index int) {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
	}
	exits[index] = true
}

func main() {
	start := time.Now().UnixNano()
	go calc(0)
	go calc(1)
	go calc(2)
	
	for {
		if exits[0] && exits[1] && exits[2] {
			break
		}
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	fmt.Printf("finished, cost:%d ms\n", (end - start)/1000/1000)
}