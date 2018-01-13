package main

import(
	"fmt"
	"time"
)

const (
	Man = 1
	Female = 2
)

func main(){
	now := time.Now()
	second := now.Unix()
	if second % Female == 0 {
		fmt.Println("female")
	} else {
		fmt.Println("man")
	}
}