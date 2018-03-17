package main

import (
	"fmt"
)

type Animal interface {
	Talk()
}

type Dog struct {
	Name string
}

func (d *Dog) Talk () {
	fmt.Println("dog is talk")
}
/*
func (d Dog) Talk () {
	fmt.Println("dog is talk")
}*/

func Talk(a Animal) {
	a.Talk()
}



func main() {
	var d Dog
	//Talk(d)
	Talk(&d)
}