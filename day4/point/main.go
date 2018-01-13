package main


import(
	"fmt"
)

func test1() {
	var a int
	a = 10
	fmt.Println(a)

	var b *int
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", a)
	fmt.Printf("%p\n", &a)

	b = &a
	fmt.Printf("%d\n", *b)

	*b = 100
	fmt.Printf("a=%d\n", a)
}

func modify(a *int) {
	*a = 100
}

func test2() {
	var b int = 1
	var p *int 
	p =  &b
	modify(p)
	fmt.Println(b)
}

func test3() {

	///p 默认初始化nil
	var p *int
	var b int
	p = &b
	*p = 200 //b = 200
	
	if (p == &b) {
		fmt.Printf("equal\n")
	}

	fmt.Printf("%p %p %p\n", p, &b, &p)

	p = new(int)
	*p = 1000
	fmt.Printf("%d\n", *p)
	fmt.Printf("%p %p %p\n", p, &b, &p)

	if (p == &b) {
		fmt.Printf("equal")
	}

	//指针类型的变量初始化：1. 使用其他变量地址给它赋值。 2. 使用new分配
}

func test4(){
	var p *string
	
	//*p = "abc"
	//1. 第一种方式
	p = new(string)

	*p = "abc"
	fmt.Println(*p)

	//2. 第二种方式
	var str string = "hello word"
	p = &str
	fmt.Println(*p)
}

func test5(){
	var a []int
	a  = make([]int, 10)
	a[0] = 100
	fmt.Println(a)

	var p *[]int
	p = new([]int)
	(*p) = make([]int, 10)
	(*p)[0] = 100
	fmt.Println(p)

	p = &a 
	(*p) [0] = 1000
	fmt.Println(a)
}

func modify_arr(a []int) {
	fmt.Printf("modify:%p\n", a)
	a[0] = 100
}

func test6() {
	var a[6]int 
	fmt.Printf("test6:%p\n", &a)
	modify_arr(a[:])
	fmt.Println(a)
}

func test() {
	return
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	test6()
}