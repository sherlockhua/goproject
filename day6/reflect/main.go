package main


import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string
}

func (s *Student) Set(name string , Age int, Sex int) {
	s.Name = name
}

func (s *Student) GetName(name string) {
	s.Name = name
}


func getAllMethod(a interface{}) {
	typeInfo := reflect.TypeOf(a)
	num := typeInfo.NumMethod()
	for i := 0;i <num; i++{
		method := typeInfo.Method(i)
		fmt.Println(method)
	}
}

func getTypeInfo(a interface{}) {
	typeInfo := reflect.TypeOf(a)
	kind := typeInfo.Kind()
	fmt.Println("kind of a:", kind)

	num := typeInfo.NumMethod()
	fmt.Println("method num:", num)

	method, ok := typeInfo.MethodByName("SetName")
	if !ok {
		fmt.Println("not have method SetName")
	} else {
		fmt.Println(method)
	}
	fmt.Println()
	fmt.Println()
}

func testGetTypeInfo() {
	var i int
	getTypeInfo(i)

	var stu Student
	getTypeInfo(&stu)

	var s [5]int
	getTypeInfo(s)
}


func testGetAllMethod() {
	
	var stu Student
	getAllMethod(&stu)
}

func testGetValuInfo() {
	var i int = 100
	valueInfo := reflect.ValueOf(&i)

	valueInfo.Elem().SetInt(200)
	tmp := valueInfo.Interface()
	val := tmp.(*int)
	fmt.Println("val:", val)
	fmt.Println("val of valueInfo:", valueInfo.Elem().Int())
	fmt.Println("type:", valueInfo.Type())
	fmt.Println("kind:", valueInfo.Kind())
	fmt.Println("i=", i)

	var stu Student
	valueInfo = reflect.ValueOf(stu)
	fmt.Println("type:", valueInfo.Type())
	fmt.Println("kind:", valueInfo.Kind())

}

func main() {
	//testGetTypeInfo()
	//testGetAllMethod()
	testGetValuInfo()
}