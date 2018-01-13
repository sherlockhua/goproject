package main

import(
	"fmt"
)

func main() {
	var a []int = make([]int, 10)
	fmt.Println(a[0])

	var b [10]int = [10]int{1,2,3, 8:100}
	fmt.Println(b)

	var c[10]int = b
	fmt.Println(c)
	c[0] = 1000
	fmt.Println(b)

	a = b[:]
	a[0] = 1000
	fmt.Println(b)
	
	a = append(a, 10, 30, 40)
	fmt.Println(a)
	a[0] = 2000

	fmt.Println(b)
}