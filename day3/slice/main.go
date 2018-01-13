package main


import(
	"fmt"
)

func test1(){
	var a [5]int 
	b := a[1:3]
	a[0] = 100
	a[1] = 200

	fmt.Printf("b:%#v\n", b)
}

func test2(){
	var a [5]int 
	b := a[1:3]
	b[0] = 100
	b[1] = 200

	fmt.Printf("b:%#v\n", a)
}

func test3(){
	var a [5]int 
	b := a[1:3]
	//越界访问会panic
	b[100] = 100
	
	fmt.Printf("b:%#v\n", a)
}

func Sum(b []int) int {
	var sum int
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}

	b[0] = 100
	return sum
}

func SumArray(b [100]int) int {
	var sum int
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}

	b[0] = 100
	return sum
}

func testSliceCap() {
	
	a := make([]int, 5, 10)
	a[4] = 100
	b := a[2:3]
	//b[9] = 100

	fmt.Printf("a=%#v, len(a) = %d, cap(a)=%d\n", a, len(a), cap(a))
	fmt.Printf("b=%#v, len(b) = %d, cap(b)=%d\n", b, len(b), cap(b))
}

func testAppend() {
	var a []int
	a = make([]int, 5)
	var b[]int = []int{10,11,12,14}
	a = append(a, b...)
	fmt.Printf("a:%#v\n", a)
}

func testArray() {
	var a  = [...]int{1,2,3,4,5,6,7,8}
	fmt.Println(a)
}

func testStrSlice() {
	var str = "hello world"
	var b []byte = []byte(str)
	b[0] = 'a'
	str1 := string(b)
	fmt.Printf("str1:%s, %d\n", str1, len(str))
}

func testStrReverse() string{
	str := "hello world"
	b := []byte(str)

	for i := 0; i < len(b)/2;i++ {
		//t := b[i]
		//b[i] = b[len(b)-i-1]
		//b[len(b)-i-1] = t
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	str1 := string(b)
	fmt.Println(str1)
	return str1
}

func testStrReverseUtf8() string{
	str := "hello world我们爱中国"
	b := []rune(str)

	for i := 0; i < len(b)/2;i++ {
		//t := b[i]
		//b[i] = b[len(b)-i-1]
		//b[len(b)-i-1] = t
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	str1 := string(b)
	fmt.Println(str1)

	fmt.Printf("len(str)=%d, len(rune)=%d\n", len(str), len(b))
	return str1
}

func main(){
	//test1()
	//test2()
	//test3()
	//var a [100]int = [100]int{1,2,3,4,5}
	/*
	var a [100]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	//result := Sum(a[:])
	result := SumArray(a)
	fmt.Printf("sum=%d\n", result)
	fmt.Printf("a:%#v\n", a)
	*/

	//testSliceCap()
	//testAppend()
	//testStrSlice()
	//testStrReverse()
	testStrReverseUtf8()
}