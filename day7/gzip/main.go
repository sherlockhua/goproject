package main

import (
	"io"
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("D:/张导讲命令.txt.gz")
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}
	defer file.Close()

	reader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("gzip failed, err:%v\n", err)
		return
	}
	bufReader := bufio.NewReader(reader)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("read failed, err:%v\n", err)
			return
		}
		fmt.Printf("%s", line)
	}
}