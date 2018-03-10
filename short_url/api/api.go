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
	
	var response model.ResponseHeader
	response.Code = code
	response.Message =  getMessage(code)

	data, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("{\"code\":500, \"message\": \"server busy\"}"))
		return
	}

	w.Write(data)
}

func responseSuccess(w http.ResponseWriter, data interface{}) {
	/*
	m := make(map[string]interface{}, 16)
	m["code"] = code
	m["message"] = getMessage(code)
	m["data"] = data
*/
	dataByte, err := json.Marshal(data)
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

	responseSuccess(w, resp)
}


func Short2Long(w http.ResponseWriter, r*http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read all failded, ", err)
		responseError(w, 1001)
		return
	}

	var req model.Short2LongRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		fmt.Println("read all failded, ", err)
		responseError(w, 1002)
		return
	}

	resp, err := logic.Short2Long(&req)
	if err != nil {
		fmt.Println("read all failded, ", err)
		responseError(w, 1003)
		return
	}

	responseSuccess(w, resp)
}

func main(){

	err := logic.InitDb("root:@tcp(10.0.0.200:3306)/short_url?parseTime=true")
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}

	http.HandleFunc("/trans/long2short", Long2Short)
	http.HandleFunc("/trans/short2long", Short2Long)
	http.ListenAndServe(":18888", nil)


}