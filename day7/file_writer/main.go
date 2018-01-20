package main


import (
	"os"
	"fmt"
)

func testWriteFile() {

	file , err := os.OpenFile("D:/mylog.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file failed:%v\n", err)
		return
	}
	defer file.Close()

	for i := 0; i < 10; i++{
		file.WriteString(fmt.Sprintf("hello %d\n", i))
	}

}

func main() {
	testWriteFile()
}