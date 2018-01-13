package main


import(
	"fmt"
)

func main() {
	pipe := make(chan string, 3)

	pipe <- "a"
	pipe <- "b"

	var a string
	a = <- pipe

	fmt.Printf("%s\n", a)
}