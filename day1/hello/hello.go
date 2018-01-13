package main

import (
	"code.oldboy.com/day1/calc"
	"fmt"
)

func main() {

	var sum int
	var sub int
	sum = calc.Add(2, 102)
	sum, sub = calc.Calc(2, 3)

	fmt.Println(sum, sub)
	SayHello()
	//fmt.Println("hello")

}
