package main


import (
	"fmt"
	"time"
)

func set(p *int) {
	*p = 123
}

func test() {
	
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("panic:%v\n", err)
		}
	}()
	var p *int
	set(p)
	fmt.Printf("*p=%d\n", *p)
}

func main() {
	for {
		test()
		time.Sleep(time.Second)
	}
}