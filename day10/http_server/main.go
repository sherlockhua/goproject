package main

import (
	"io"
	"net/http"
)

const form = `<html>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"> 
<body><form action="#" method="post" name="bar">
<div> 姓名：<input type="text" name="username1"/></div>
<div> 密码：<input type="text" name="password"/></div>
<input type="submit" value="登录"/>
</form></html></body>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}
func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Server", "Go 1.9.2 Server")
	w.Header().Set("Niubi", "hahahaha")

	var p *int
	*p = 10
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		request.ParseForm()
		io.WriteString(w, request.FormValue("username"))
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("password"))
	}
}
func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
	}
}
