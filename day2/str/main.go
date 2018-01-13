package main


import(
	"fmt"
)

func main(){
	var str string
	str = "abc\n"

	fmt.Printf("%s\n", str)

	var str2 string
	str2 = `abc
	
	
	hello 
	
	
	\n`
	fmt.Printf("%s\n", str2)
}