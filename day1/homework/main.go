package main


import (
	"fmt"
)



func main() {
	
	var str string
	str = "hello world"
	fmt.Printf("%s\n", str)

	var num  int
	num = 100

	fmt.Printf("%d %b %x\n", num, num, num)
	fmt.Printf("%b\n", num)
	fmt.Printf("%x\n", num)

	var t float32
	t = 3.2
	fmt.Printf("%f\n", t)
}