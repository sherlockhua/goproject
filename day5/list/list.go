package main

import (
	"fmt"
)

type Teacher struct {
	Name string
	Age int
	Next *Teacher
}

func NewTeacher(name string, age int) *Teacher {
	p := new(Teacher)
	p.Name = name
	p.Age = age
	return p
}

func createList() {
	var header *Teacher = &Teacher{}
	header.Age = 200
	header.Name = "a"

	fmt.Println("第一次打印")
	printList(header)
	
	/*p := new(Teacher)
	p.Name = "b"
	p.Age = 100*/
	p := NewTeacher("b", 100)
	header.Next = p

	fmt.Println("第二次打印")
	printList(header)

	p = new(Teacher)
	p.Name = "c"
	p.Age = 100

	header.Next.Next = p

	fmt.Println("第三次打印")
	printList(header)
}

func createInHeader(h *Teacher, name string, age int) (*Teacher) {
	p := &Teacher{}
	p.Age = age
	p.Name = name

	p.Next = h
	return p
}

func printList(h *Teacher) {
	for h != nil {
		fmt.Printf("Name:%v Age:%v\n", h.Name, h.Age)
		h = h.Next
	}
}

func testCreateInHeader() {
	var header *Teacher
	header = createInHeader(header, "a", 18)
	header = createInHeader(header, "b", 19)
	header = createInHeader(header, "c", 20)


	header = createInHeader(header, "a", 18)
	header = createInHeader(header, "b", 19)
	header = createInHeader(header, "c", 20)

	printList(header)
}


func createInTail(tail *Teacher, name string, age int) (*Teacher) {
	p := &Teacher{}
	p.Age = age
	p.Name = name

		if tail == nil {
			return p
		}
	tail.Next = p
	return p
}

func testCreateInTail() {
	var header *Teacher
	var tail *Teacher = header

	tail = createInTail(tail, "a", 18)
	if header == nil {
		header = tail
	}

	tail = createInTail(tail, "b", 19)
	tail = createInTail(tail, "c", 20)


	tail = createInTail(tail, "a", 18)
	tail = createInTail(tail, "b", 19)
	tail = createInTail(tail, "c", 20)

	printList(header)
}
