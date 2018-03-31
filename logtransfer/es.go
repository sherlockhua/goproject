package main

import (
	"github.com/astaxie/beego/logs"
	"sync"
	elastic "gopkg.in/olivere/elastic.v2"
)

var waitGroup sync.WaitGroup
var client *elastic.Client

func initES(addr string) (err error){
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.12.3:9200/"))
	if err != nil {
		logs.Error("connect to es error:%v", err)
		return
	}

	logs.Debug("connect es success")
	return
	/*
	tweet := Tweet{User: "olivere name", Message: "Take Five"}
	_, err = client.Index().
		Index("twitter").
		Type("tweet").
		BodyJson(tweet).
		Do()
	if err != nil {
		// Handle error
		panic(err)
		return
	}
	fmt.Println("insert succ")
	*/
}

func Run(threadNum int) (err error){

	for i := 0; i < threadNum; i++{
		waitGroup.Add(1)
		go sendToEs()
	}
	waitGroup.Wait()
	return
}

type EsMessage struct {
	Message string
}

func sendToEs() {
	for msg := range GetMessage() {
		var esMsg EsMessage
		esMsg.Message = msg.line
		logs.Debug("begin send to es succ")
		
		_, err := client.Index().
			Index(msg.topic).
			Type(msg.topic).
			BodyJson(esMsg).
			Do()
		if err != nil {
			logs.Error("send to es failed, err:%v", err)
			continue
		}
	
		logs.Debug("send to es succ")
	}
	waitGroup.Done()
}