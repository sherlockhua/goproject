package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/sherlockhua/goproject/short_url/logic"
	"github.com/sherlockhua/goproject/short_url/model"
)

const (
	ErrSuccess = 0
	ErrInvalidParameter = 1001
	ErrServerBusy = 1002
)

func getMessage(code int) (msg string){
	switch code {
	case ErrSuccess:
		msg = "success"
	case ErrInvalidParameter:
		msg = "invalid parameter"
	case ErrServerBusy:
		msg = "server busy"
	default:
		msg = "unknown error"
	}

	return
}


func responseError(w http.ResponseWriter, code int) {
	
	m := make(map[string]interface{}, 16)
	m["code"] = code
	m["message"] = getMessage(code)

	data, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("{\"code\":500, \"message\": \"server busy\"}"))
		return
	}

	w.Write(data)
}

func responseSuccess(w http.ResponseWriter, code int, data interface{}) {
	
	m := make(map[string]interface{}, 16)
	m["code"] = code
	m["message"] = getMessage(code)
	m["data"] = data

	dataByte, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("{\"code\":500, \"message\": \"server busy\"}"))
		return
	}

	w.Write(dataByte)
}

func Long2Short(w http.ResponseWriter, r*http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read all failded, ", err)
		responseError(w, 1001)
		return
	}

	var req model.Long2ShortRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Println("read all failded, ", err)
		responseError(w, 1002)
		return
	}

	resp, err := logic.Long2Short(&req)
	if err != nil {
		fmt.Println("read all failded, ", err)
		responseError(w, 1003)
		return
	}

	responseSuccess(w, 0, resp)
}


func Short2Long(w http.ResponseWriter, r*http.Request) {
	
}

func main(){
	http.HandleFunc("/trans/long2short", Long2Short)
	http.HandleFunc("/trans/short2long", Short2Long)
	http.ListenAndServe(":18888", nil)
}