package main


import (
	"io/ioutil"
	"fmt"
	"bytes"
	"github.com/sherlockhua/goproject/short_url/model"
	"encoding/json"
	"net/http"
)

func getShortUrl() (short string) {
	var buffer bytes.Buffer
	var req model.Long2ShortRequest 
	req.OriginUrl = "https://app.getpostman.com/app/download/win64?_ga=152.36091171.223405465.1520069059-1132478787.1520069059"
	data, _ := json.Marshal(req)

	buffer.WriteString(string(data))
	resp, err := http.Post("http://localhost:18888/trans/long2short", "application/json", &buffer)
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
	var response model.Long2ShortResponse
	/*
	var respMap map[string]interface{} = make(map[string]interface{},10)
	err = json.Unmarshal(result, &respMap)
	fmt.Printf("err:%v short url:%s\n", err,  respMap)

	dataField := respMap["data"]
	dataFieldMap := dataField.(map[string]interface{})
	shortField := dataFieldMap["short_url"]
	shortFieldUrl := shortField.(string)
	*/
	err = json.Unmarshal(result, &response)
	
	short =  response.ShortUrl
	return
}


func getOriginUrl(shortUrl string) (originUrl string) {
	var buffer bytes.Buffer
	var req model.Short2LongRequest 
	req.ShortUrl = shortUrl
	data, _ := json.Marshal(req)

	buffer.WriteString(string(data))
	resp, err := http.Post("http://localhost:18888/trans/short2long", "application/json", &buffer)
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(result))
	var response model.Short2LongResponse
	/*
	var respMap map[string]interface{} = make(map[string]interface{},10)
	err = json.Unmarshal(result, &respMap)
	fmt.Printf("err:%v short url:%s\n", err,  respMap)

	dataField := respMap["data"]
	dataFieldMap := dataField.(map[string]interface{})
	shortField := dataFieldMap["short_url"]
	shortFieldUrl := shortField.(string)
	*/
	err = json.Unmarshal(result, &response)
	
	originUrl =  response.OriginUrl
	return
}


func main() {
	
	shortUrl := getShortUrl()
	fmt.Printf("short url:%s\n", "http://localhost:18888/"+shortUrl)
	originUrl := getOriginUrl(shortUrl)
	fmt.Printf("origin url:%s\n", originUrl)
}