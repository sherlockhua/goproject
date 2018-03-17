package main

import (
	"time"
	"fmt"
)

func calc() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var p *int
	*p = 100
}

func main() {
	go calc()
	time.Sleep(time.Second*3)
	fmt.Println("progress exited")
}