package main

import(
	"fmt"
	_ "code.oldboy.com/day2/mysql_driver"
)

func init() {
	fmt.Println("init called")
}

func main(){
	fmt.Println("main called")
}