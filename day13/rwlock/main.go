package main


import (
	"time"
	"sync"
	"fmt"
)

var rwLock sync.RWMutex
var lock sync.Mutex
var w sync.WaitGroup
var count int

func main() {
	w.Add(1)
	start := time.Now().UnixNano()
	go func() {
		for i := 0; i < 1000;i++ {
			//rwLock.Lock()
			lock.Lock()
			count++
			time.Sleep(5*time.Millisecond)
			lock.Unlock()
			//rwLock.Unlock()
		}
		w.Done()
	}()

	//for i := 0; i < 1000000; i++ {
	for i := 0; i < 16; i++ {
		w.Add(1)
		go func() {
			for i := 0; i < 5000;i++ {
				//rwLock.RLock()
				lock.Lock()
				time.Sleep(1*time.Millisecond)
				lock.Unlock()
				//rwLock.RUnlock()
			}
			w.Done()
		}()
	}
	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start)/1000/1000)
	//fmt.Println(count)
}