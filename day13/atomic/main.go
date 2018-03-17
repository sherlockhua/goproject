package main

import (
	"sync/atomic"
	"sync"
	"fmt"
	"time"
)

var w sync.WaitGroup
var count int32

func main() {
	w.Add(1)
	start := time.Now().UnixNano()
	go func() {
		for i := 0; i < 1000000;i++ {
			atomic.AddInt32(&count, 1)
		}
		w.Done()
	}()

	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(&count, 1)
	}

	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start)/1000/1000)
	fmt.Println(count)
}