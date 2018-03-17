package main

import (
	"os"
	"bufio"
	"fmt"
)
func stat(str string) {
	var en_count int
	var sp_count int
	var num_count int
	var other_count int
	utf8Arr := []rune(str)
	for _, v := range utf8Arr {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			en_count++
			continue
		}
		if v == ' ' {
			sp_count++
			continue
		}
		if v >= '0' && v<= '9' {
			num_count++
			continue
		}
		other_count++
	} 
	fmt.Printf("en=%d sp=%d num=%d other=%d\n", en_count, sp_count, num_count, other_count)
}


func main(){
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("read string failed, err:%v\n", err)
		return
	}

	stat(line)
}