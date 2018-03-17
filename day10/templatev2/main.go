package main

import (
	"fmt"
	//"os"
	"text/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var (
	gtemp *template.Template
)

func init() {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	gtemp = t
}

func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	
	var persons []*Person
	for i := 0; i < 10; i++{
		p := &Person{
			Name: fmt.Sprintf("stu%d", i),
			Age: i*10,
		}

		persons = append(persons, p)
	}

	//执行模板渲染
	gtemp.Execute(w, persons)
}

func main() {
	
	http.HandleFunc("/user_info", handleUserInfo)
	http.ListenAndServe(":8080", nil)
}
