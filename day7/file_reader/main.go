package main 

import (
	"os"
	"fmt"
	"io"
	"bufio"
	"io/ioutil"
)

func testFile() {
	file, err := os.Open("D:/运维人员如何最大限度避免误删除文件-演示文件.txt")
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}

	defer file.Close()
	var data [1024]byte
	for {
		n, err := file.Read(data[:])
		if err == io.EOF {
			break
		} 
		if err != nil { 
			fmt.Printf("read file error:%s\n", err)
			return
		}

		str := string(data[0:n])
		fmt.Println(str)
	}
}

func testBufIO() {
	file, err := os.Open("D:/运维人员如何最大限度避免误删除文件-演示文件.txt")
	if err != nil {
		fmt.Println("open file failed:", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read file error:%s\n", err)
			return
		}
		fmt.Println(line)
	}
}

func testIoUtil() {
	data, err := ioutil.ReadFile("D:/运维人员如何最大限度避免误删除文件-演示文件.txt")
	if err != nil {
		fmt.Printf("read file failed:%v\n", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}

func main(){
	//testFile()
	//testBufIO()
	testIoUtil()
}