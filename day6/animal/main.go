package main


import (
	"fmt"
)

type Animal interface {
	Eat() 
	Talk()
}

type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	fmt.Println(d.Name, "is eating")
}

func (d *Dog) Talk() {
	fmt.Println(d.Name, " is wawa!")
}

type Cat struct {
	Name string
}

func (d *Cat) Eat() {
	fmt.Println(d.Name, "is eating")
}

func (d *Cat) Talk() {
	fmt.Println(d.Name, "is 喵喵!")
}

func Test() {
	var a Animal
	//a.Eat()
	var d Dog
	d.Eat()

	a = &d
	a.Eat()

	var c Cat
	a = &c
	a.Eat()
}

func TestOperator() {
	var animalList []Animal
	d := &Dog{
		Name: "小黄",
	}
	animalList = append(animalList, d)

	d1 := &Dog{
		Name: "旺财",
	}
	animalList = append(animalList, d1)

	c1 := &Cat{
		Name: "小白",
	}
	animalList = append(animalList, c1)

	for _, v := range animalList {
		v.Eat()
		v.Talk()
	}
}

func main() {
	//Test()
	TestOperator()
}