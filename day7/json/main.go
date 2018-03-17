package main


import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"io/ioutil"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"` 
	Score float32 `json:"score"`
}

func testWriteFile(){
	var stus []*Student
	for i := 0; i <1000; i++{
		stu := &Student {
			Name:fmt.Sprintf("stu%d", i),
			Age:rand.Intn(100),
			Score:rand.Float32()*100,
		}
		stus = append(stus, stu)
	}

	data, err := json.Marshal(stus)
	if err != nil {
		fmt.Printf("json failed, err:%v\n", err)
		return
	}

	file, err := os.OpenFile("d:/stu.dat", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("file failed, err:%v\n", err)
		return
	}
	defer file.Close()

	n, err := file.Write(data)
	if err != nil {
		fmt.Printf("write failed, err:%v\n", err)
		return
	}
	fmt.Printf("write %d succ\n", n)
}

func testReadFile(){
	file, err := os.OpenFile("d:/stu.dat", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("file failed, err:%v\n", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("read all failed, err:%v\n", err)
	}

	var stus []*Student
	err = json.Unmarshal(data, &stus)
	if err != nil {
		fmt.Printf("json unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("total student:%d\n", len(stus))
	for i := 0; i <10; i++{
		fmt.Printf("student:%#v\n", stus[i])
	}
}

func main(){
	//testWriteFile()
	testReadFile()
}