package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(a int) {
			fmt.Println(a)
			waitGroup.Done()
		}(i)
	}
	
	waitGroup.Wait()
}