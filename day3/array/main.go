package main



import(
	"fmt"
	"math/rand"
	//"strings"
)

func test1() {
	var a [25]int8
	length := len(a)
	for i := 0; i < length; i++ {
		fmt.Printf("%p\n", &a[i])
	}
	
	for index, _ := range a {
		fmt.Printf("a[%d]=%d\n", index, a[index])
	}
}

func test2() {
	var a [5]int = [5]int {1,2,3,4,5}
	var b [5] int
	b = a
	fmt.Printf("b=%v\n", b)
	b[0]  = 200
	fmt.Printf("b=%v\n", b)
	fmt.Printf("a=%v\n", a)
}

func test3() {
	var a [5] int = [5]int{1,2,3,4,5}
	fmt.Printf("%v\n", a)

	var b = [...]int{1,3,4,5,7, 8}
	fmt.Printf("%v\n", b)

	var c = [5]int{1,3,4}
	fmt.Printf("%v\n", c)

	var d [5]string = [5]string{1:"abc", 4:"efg"}
	fmt.Printf("%#v\n", d)
}

func test4() {
	var a [4][2]int 
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = (i+1)*(j+1)
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
}

func genRand() {
	var a [100]int
	for i := 0; i < len(a); i++ {
		//赋值
		a[i] = rand.Int()
	}

	for i := 0; i < len(a); i++ {
		//取下标=i的元素的值
		fmt.Printf("%d\n", a[i])
	}
}

func genRandStr() {
	var a [100]string
	var b string = "0123456789我爱中国"
	//var runeArr = []rune(b)
	for i := 0; i < len(a); i++ {
		var str string
		for j := 0; j < 4; j++ {
			index := rand.Intn(len(b))
			//格式化并返回格式化后的字符串
			str = fmt.Sprintf("%s%c",  str, b[index])
		}
		a[i] = str
		fmt.Printf("a[%d]=%s\n", i, a[i])
	}
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//genRand()

	str := [5]int{1,2,3,4,5}
	fmt.Println(str)
	//genRandStr()
}