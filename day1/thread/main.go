package main

import(
	"fmt"
	"time"
)

func PrintOdd(n int) {
	//var i int
	for i := 0; i < n; i++ {
		if i % 2 == 1 {
			fmt.Printf("%d\n", i)
		}
	}
}

func PrintEven(n int) {
	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}

func main() {
	go PrintOdd(10)
	go PrintEven(10)

	time.Sleep(10*time.Second)
}