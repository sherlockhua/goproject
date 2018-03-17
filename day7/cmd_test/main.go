package main

import (
	//"os"
	"fmt"
	"flag"
)

var (
	conf string
	level int
)

func init(){
	flag.StringVar(&conf, "c", "D:/etc/test.conf", "请指定配置文件路径")
	flag.IntVar(&level, "l", 8, "请指定日志级别")
	flag.Parse()
}

func main(){
	fmt.Printf("conf is:%s\n", conf)
	fmt.Printf("level is:%d\n", level)
}