package main


import (
	"fmt"
)

func main(){
	var a interface{}
	var b int = 1000

	a = b
	fmt.Println(a)

	var c string = "hello"
	a = c
	fmt.Println(a)
}