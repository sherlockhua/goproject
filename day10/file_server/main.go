package main

import (
	"fmt"
	//"os"
	"html/template"
	"net/http"
)


var (
	gtemp *template.Template
)

func init() {
	t, err := template.ParseFiles("./views/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	gtemp = t
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	gtemp.Execute(w, nil)
} 


func main() {
	
	http.HandleFunc("/index", handleIndex)
	//http.Handle("/static/", http.FileServer(http.Dir("./static/")))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("./media/"))))
	http.ListenAndServe(":8080", nil)
}
