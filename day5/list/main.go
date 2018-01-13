package main

import(
	"fmt"
)

type People struct {
	Age int
	Name string
	Next *Student
}

type Student struct {
	 Age int
	Name string
	Next *People
}

func testList(){
	var s Student
	s.Age = 100
	s.Name = "abc"
	s.Next = new(People)
	/*s.Next = &People{
		Age:100,
		Name:"efg",
	}*/
	s.Next.Age = 1000
	s.Next.Name = "efg"
	s.Next.Next = new(Student)
	s.Next.Next.Age = 100
	s.Next.Next.Name = "999"

	fmt.Printf("s:%+v\n", s)
	fmt.Printf("next:%v\n", *(s.Next))
	fmt.Printf("people.next:%#v\n", *(s.Next.Next))

	fmt.Printf("list header:%#v\n", s)
	
		fmt.Printf("data:%#v\n", *(s.Next))
		fmt.Printf("data:%#v\n", *(s.Next.Next))
		fmt.Printf("data:%#v\n", s.Next.Next.Next)
		//s.Next = s.Next.Next
}

func main(){
	//testList()
	//createList()
	//testCreateInHeader()
	testCreateInTail()
}