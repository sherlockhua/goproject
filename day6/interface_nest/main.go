package main

import (
	"fmt"
)

type Eater interface {
	Eat()
}

type Talker interface {
	Talk()
}

type Animal interface {
	Eater
	Talker
}

type Dog struct {
}

func (d *Dog) Eat() {
	fmt.Println("eating")
}

func (d *Dog) Talk() {
	fmt.Println("eating")
}

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("eating")
}

func (c *Cat) Talk() {
	fmt.Println("eating")
}

func justify(a Animal) {
	/*
	dog, ok := a.(*Dog)
	if !ok {
		fmt.Println("conver to dog failed")
		return
	}
	dog.Eat()
*/

	switch t := a.(type) {
	case *Dog:
		t.Eat()
		fmt.Printf("t is dog\n")
	case *Cat:
		t.Eat()
		fmt.Printf("t is cat\n")
	}
}

func main() {
	d := &Dog{}
	var a Animal
	a = d
	a.Eat()

	justify(a)
	a = &Cat{}

	justify(a)
}
