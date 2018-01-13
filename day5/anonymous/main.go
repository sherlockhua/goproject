package main


import(
	"fmt"
)

type People struct {
	Name string
	Age int
}

func (p *People) Format() string {
	return fmt.Sprintf("name=%s&age=%d", p.Name, p.Age)
}

type Student struct{
	Score int
	People
	Name string
}

func (p *Student) Format() string {
	return fmt.Sprintf("name=%s&age=%d", p.Name, p.Age)
}


func test1() {
	var s Student
	//s.People.Name = "abc"
	//s.People.Age = 100
	s.Name = "abc"
	s.People.Name = "cdg"
	s.Age = 100
	s.Score = 100

	fmt.Printf("%#v\n", s)
}

func testMethod() {
	var s Student
	s.Age = 200
	s.People.Name = "abc"
	ret := s.People.Format()
	fmt.Println("format result:", ret)
}

func main() {
	testMethod()
}