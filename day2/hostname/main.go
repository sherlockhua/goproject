package main

import(
	"fmt"
	"os"
)

func swapv2(a int, b int) (int, int) {
	return b, a
}

func swap(a int, b int) (int, int) {
	var c int
	c = a
	a = b
	b = c

	return a, b
}

func main(){
	name, ret := os.Hostname()
	fmt.Printf("%s %v\n", name, ret)

	val := os.Getenv("PATH")
	fmt.Printf("%s\n", val)

	var (
		a int = 100
		b int = 300
	)
	//a, b = swapv2(100, 300)
	a, b = b, a
	fmt.Printf("a=%d b=%d\n", a, b)
}