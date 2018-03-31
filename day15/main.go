package main


import (
	"fmt"
)
type Person struct {
	Name string
	Sex string
	Age int
}
func (p *Person) SetName(name string) *Person{
	p.Name = name
	return p
}
func (p *Person) SetSex(sex string)*Person{
	p.Sex = sex
	return p
}
func (p *Person) SetAge(age int) *Person{
	p.Age = age
	return p
}

func (p *Person) Print(){
	fmt.Printf("name:%s\nsex:%s\nage:%d\n", p.Name, p.Sex, p.Age)
}
func main(){
	p := &Person{}
	/*
	p.SetAge(18)
	p.SetName("xiaoming")
	p.SetSex("man")
	*/
	p.SetAge(18).SetName("xiaoming").SetSex("man").Print()

	p.Print()
}