package main

import (
	"sync"
	"fmt"
	"time"
)

var lock sync.Mutex
var w sync.WaitGroup
var count int

func main() {
	start := time.Now().UnixNano()
	w.Add(1)
	go func() {
		for i := 0; i < 1000000;i++ {
			lock.Lock()
			count++
			lock.Unlock()
		}
		w.Done()
	}()

	for i := 0; i < 1000000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}

	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start)/1000/1000)
	fmt.Println(count)
}