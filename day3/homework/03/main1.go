package main

import (
	"log"
	"fmt"
)

func handler(str string) {
	if len(str) <= 1 {
		log.Fatal(str, " is not string type")
	}
	s := []rune(str)
	for i:=0; i< len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	result := string(s)
	if result == str {
		fmt.Printf("%s is 回文字符串\n", str)
	}
}




func main() {
	fmt.Println("请输入一个字符串：")
	var str string
	fmt.Scanln(&str)
	handler(str)
}