package main

import (
	"fmt"
	"regexp"

)

//练习题四：统计一个字符串的字母，数字，空格，其他字符个数
func statisticalCharacterNumber(s string) {
	//示例字符串
	// var s = "MemTotal: 100111 KB"
	//创建正则匹配字母、数字、空格、
	var rNum = regexp.MustCompile(`\d`)
	var rCharacter = regexp.MustCompile("[a-zA-Z]")
	var rBlank = regexp.MustCompile(" ")
	//定义其他字符变量
	var specialcharacter int
	//获取相应的字母、数字、空格的个数
	num := len(rNum.FindAllStringSubmatch(s, -1))
	character := len(rCharacter.FindAllStringSubmatch(s, -1))
	blank := len(rBlank.FindAllStringSubmatch(s, -1))
	fmt.Printf("数字个数%d\n", num)
	fmt.Printf("字母个数%d\n", character)
	fmt.Printf("空格个数%d\n", blank)
	//其他字符
	specialcharacter = len(s) - num - character - blank

	fmt.Printf("其他字符个数%d\n", specialcharacter)
	fmt.Printf("总个数%d\n", len(s))
}

func main () {
	statisticalCharacterNumber("MemTotal: 100111 KB")
}