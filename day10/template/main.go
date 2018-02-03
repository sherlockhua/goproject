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
	t, err := template.ParseFiles("./index2.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	gtemp = t
}

func handleUserInfo(w http.ResponseWriter, r *http.Request) {
	/*p := &Person{
		Name:"mary", 
		Age: 20,
	}*/

	m := make(map[string]int, 10)
	m["user_name"] = 33
	m["age"] = 388

	gtemp.Execute(w, m)
}

func main() {
	
	http.HandleFunc("/user_info", handleUserInfo)
	http.ListenAndServe(":8080", nil)
}
