package main


import (
	"fmt"
	"time"
)

func test() {
	var intChan chan int = make(chan int, 1)
	go func() {
		intChan <- 10
	}()

	result := <- intChan
	fmt.Printf("result:%d\n", result)
}

func testNoBufChan() {
	
	var intChan chan int = make(chan int, 2)
	go func() {
		
		fmt.Printf("begin input to chan\n")
		intChan <- 10
		intChan <- 10
		fmt.Printf("end input to chan\n")
	}()

	//result := <- intChan
	//fmt.Printf("result:%d\n", result)
	time.Sleep(10*time.Second)
}

func main() {
	testNoBufChan()
}