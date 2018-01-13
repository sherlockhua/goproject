package main


import(
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
	d1 = 2
	d2
	d3 
	e1 = iota
	e2 = iota
	e3 = iota
)


func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)

	fmt.Println(e1)
	fmt.Println(e2)
	fmt.Println(e3)
}