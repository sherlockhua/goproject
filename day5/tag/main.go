package main


import (
	"fmt"
	"encoding/json"
)

type Student struct {
	Name string `json:"name"`
	Age  int  `json:"age"`
	Sex string `json:"sex"`
}

func main() {
	var s Student
	s.Age = 200
	s.Name = "abc"
	s.Sex = "man"

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("json marshal failed, err:%v", err)
		return
	}

	fmt.Printf("json data:%s\n", data)

	var s1 Student 
	err = json.Unmarshal(data, &s1)
	if err != nil {
		fmt.Printf("json Unmarshal failed, err:%v", err)
		return
	}
	fmt.Printf("s1:%#v\n", s1)
}