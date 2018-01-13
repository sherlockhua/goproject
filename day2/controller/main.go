package main


import(
	"fmt"
	"strings"
)

func Test() int {
	return 100
}

func TestFor() {
	var a int
	for a = 0; a < 10; a++ {
		fmt.Println(a)
	}
}


func TestForV2() {
	var a int = 0
	for a < 10 {
		fmt.Println(a)
		a++
	}
}

func TestForStr() {

	str := "hello world"
	for index, val := range str {
		fmt.Printf("index:%d val:%c val2:%c\n", index, val, str[0])
	}
}

func StrOperator (){
	str := "hello "
	fmt.Printf("len:%d\n", len(str))

	str2 := str +"world"
	fmt.Printf("str2:%s\n", str2)

	str3 := "the,character,represented,by,the corresponding Unicode code point"
	result := strings.Split(str3, ",")
	fmt.Printf("result:%v\n", result)

	str5 := strings.Join(result, ",")
	fmt.Printf("str5:%s\n", str5)

	isContain := strings.Contains(str3, "represented33")
	fmt.Printf("contain:%t\n", isContain)

	str4 := "baidu.com"
	index := strings.Index(str4, "du")
	fmt.Printf("index:%d\n", index)

	if ret := strings.HasPrefix(str4, "http://"); ret == false {
		str4 = "http://" +str4
	}

	fmt.Printf("str4：%s\n", str4)

	
}

func TestScanf() {
	var number int
	var str string 
	fmt.Scanf("%d %s", &number, &str)
	fmt.Println(number, str)
}

func main(){
	TestScanf()
	//StrOperator()
	//TestForV2()
	//TestForStr()
	/*
	if a := Test(); a > 10 {
		fmt.Printf("good %d\n", a)
	}*/
}