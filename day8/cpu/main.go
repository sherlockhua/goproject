package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU()
	fmt.Printf("cpu num:%d\n", num)
	runtime.GOMAXPROCS(1)
}