package main

import(
	"fmt"
)

type Int int

type Test struct {
	A int
	b int

}

type Student struct {
	Age int 
	Name string
	Sex string
	Grade string
	Score int
	t Int
	a Test
	c *int
}

func testInt() {
	var a Int
	var b int
	a = Int(b)

	fmt.Println(a,  b)
}
func testStruct () {
	var s Student
	s.Age = 18
	s.Name = "tom"
	s.Score = 80
	s.Sex = "man"
	s.a.A = 100
	s.c = new(int)
	*(s.c) = 100

	fmt.Printf("name:%s age:%d score:%d sex:%s c=%d\n", s.Name, s.Age, s.Score, s.Sex, *(s.c))
	fmt.Printf("%+v\n", s)
	s1 := s
	s1.Name = "jim"
	*(s1.c) = 200

	fmt.Printf("name:%s age:%d score:%d sex:%s c=%d\n", s.Name, s.Age, s.Score, s.Sex, *(s.c))
	fmt.Printf("%+v\n", s)
	
	var p1 *int = new(int)
	p2 := p1
	*p2 = 100
	fmt.Printf("s1=%d\n", *p1)

	var p3 = new (Student)
	(*p3).Score = 100

	p4 := p3
	//语法糖，底层会转成 (*p4).Score=1000的形式
	p4.Score = 1000
	fmt.Printf("p3=%+v\n", *p3)
}

func main() {

	//testInt()
	testStruct() 
}