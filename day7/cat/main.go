package main


import (
	"os"
	"fmt"
	//"bufio"
	"io"
)

func cat (filename string) {
	file, err := os.Open(filename) 
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}

func main() {
	if (len(os.Args) == 1) {
		fmt.Println("请指定文件名")
		return
	}

	for i := 1; i <len(os.Args);i++{
		cat(os.Args[i])
	}
}
