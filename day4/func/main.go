package main


import(
	"fmt"
	"sort"
)

type Int int

func add (a, b, c, d int) {

}

func test(){
	a = a + 100
	var b int = 1000
	fmt.Println(a)
	fmt.Println(b)

	if a > 10 {
		var c int = 2000
		b = c
		var a int
		a = 1000
		fmt.Println(c, a)
	}

	
	for i := 0; i <100; i++{
		fmt.Println(i)
	}

	
}

func testType() {
	var a Int
	a = 100
	a = a +100

	var b int
	b = int(a)

	fmt.Println(a, b)
}

var is_login bool
func testOrder(){
	if (is_login == false) {
		return 
	}
}

func testAccount() {
	if (is_login == false) {
		return 
	}
}
type handle func()

func login(run handle) {
	
	if is_login == false {
		return
	}

	//登录通过
	run()
}
func testLogin(){

	login(testAccount)
	login(testOrder)
}

func Add(arg...int) int {
	var sum int
	for i := 0; i < len(arg); i++{
		sum = sum + arg[i]
	}
	return sum
}

func testArg() {
	fmt.Println(Add(1))
	fmt.Println(Add())
	fmt.Println(Add(1,2,3,4,5))
}

func testDefer() {
	var a int = 100
	fmt.Println("before defer:a=", a)
	defer fmt.Println(a)
	if a > 100 {
		return
	}

	a = 200
	defer fmt.Println(a)
	fmt.Println("after defer: a = ", a)
}

func testRecusive() {
	fmt.Println("hello")
	testRecusive()
}

func Adder() func(int)int {

	var x int
	f := func(i int)int {
		x = x + i
		return x
	}

	return f
}

func testClosure() {
	
	f1 := Adder()
	fmt.Println(f1(10))
	fmt.Println(f1(20))
	fmt.Println(f1(30))

	f2 := Adder()
	fmt.Println(f2(10))
	fmt.Println(f2(20))
	fmt.Println(f2(30))
}

func testMap()  {
	var a map[string]int
	a = make(map[string]int, 100)
	a["abc"] = 0
	a["hello"] = 1200
	a["cello"] = 1200
	fmt.Println(a)

	var keys []string

	for k, v := range a {
		fmt.Printf("a[%s] = %d\n", k, v)
		keys = append(keys, k)
	}

	fmt.Println("\n")
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("a[%s]=%d\n", k, a[k])
	}

	//第一种写法，返回val用来接收key "abc"的值，exist对应key "abc"是否在a中存在。
	val, exist := a["abc"]
	fmt.Printf("val=%d ok = %t\n", val, exist)
	if exist {
		fmt.Printf("val = %d\n", val)
	} else {
		fmt.Println("not found")
	}

	//第二种写法，val直接获取key aaa得值，如果aaa不存在，则val为0。这种写法无法区分aaa是否存在
	val = a["aaa"]
	fmt.Println(val)
}

func testMapSlice(){
	s := make([]map[string]int, 10)
	for i := 0;i <len(s); i++{
		s[i] = make(map[string]int, 100)
	}
	s[0]["abc"] = 100
	s[5]["abc"] = 100
	fmt.Println(s)
}

func main(){
	//test()
	//fmt.Println(a)
	//testType()
	//testArg()
	//testDefer()
	//testRecusive()
	//testClosure()
	testMap()
	//testMapSlice()
}