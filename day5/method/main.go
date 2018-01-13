package main


import (
	"fmt"
	"code.oldboy.com/day5/model"
)

type Int int

func Add(a, b int) int{
	return a + b
}

func (i *Int)Add(a, b int)  {
	*i =  Int(a + b)
	return 
}

func (i Int)Sub(a, b int)  {
	i =  Int(a - b)
	return 
}

func testInt() {
	//c := Add(100, 300)
	//fmt.Println(c)

	var a Int
	a.Add(100, 200)
	fmt.Println(a)
	a.Sub(100, 200)
	fmt.Println(a)
}

type Student struct {
	Name string
	Age int
}

func (s *Student) Set(name string, age int) {
	s.Name = name
	s.Age = age
}

func testStudent() {
	var s Student
	s.Set("abc", 100)
	fmt.Println(s)
}

func testModel() {
	school := model.NewSchool("北京大学", "北京海淀区")
	fmt.Printf("school name is:%s\n", school.GetName())
	fmt.Printf("school addr is:%s\n", school.GetAddr())
}

func main() {
	testModel()

}