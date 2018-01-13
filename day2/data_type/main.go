package main

import(
	"fmt"
	//"code.oldboy.com/day1/calc"
	"time"
)

func int_test() {
	var num int8 = 16
	var num3 int
	var num4 int32
	var b bool = false
	fmt.Printf("%d %t\n", num, b)

	num3 = int(num4)
	fmt.Printf("%d\n", num3)

	num3 = num3 + int(num4)
	fmt.Printf("%%%d\n", num3)

	fmt.Printf("%T\n", num)
}

func Test() int{
	return 9
}

func main() {
	//result := calc.Add(2, 3)
	//fmt.Println(result)
	//weibo_test()
	//binary_test()
	//int_test()
	Test()
	time.Sleep(time.Second*100)
}