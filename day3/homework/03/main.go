// 判断是否为回文
package main1

import (
	"fmt"
)

func palindrome(str string) bool {
	// 将输入的string进行反转
	c := make([]string,0)
	for j := len(str) - 1; j >= 0; j-- {
		c = append(c, string(str[j]))
	}
	// 存储正序的字符串
	var str2 string 
	for i := 0; i <= len(str) -1 ; i ++ {
		str2 += fmt.Sprintf("%v", c[i])
	}
	// 进行判断
	if str2 == str {
		fmt.Println("true")
		return true
	}
	fmt.Println("false")
	return false
}

func main() {
	palindrome("adca")
}

