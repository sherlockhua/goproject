package main


import(
	"fmt"
)

func binary_test() {
	var num int
	num = 1 << 10
	fmt.Printf("%v\n", num)

	num = num >> 10
	fmt.Printf("%v\n", num)
}