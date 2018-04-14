package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	_ "github.com/sherlockhua/goproject/seckill/proxy/router"
	"github.com/sherlockhua/goproject/seckill/proxy/model"
)

func initModel() (err error) {
	/*
	var modelConf model.ModelConf
	modelConf.RecvQueueName = proxyConf.RecvQueueName
	modelConf.SendQueueName = proxyConf.SendQueueName
	modelConf.RedisAddr = proxyConf.RedisAddr
	*/
	err = model.Init(&proxyConf)
	return
}

func main() {
	
	err := initConfig()
	if err != nil {
		panic(fmt.Sprintf("init config failed, err:%v", err))
	}

	err = initLog(proxyConf.LogPath, proxyConf.LogLevel)
	//beego.SetLogger("file", logConf)
	logs.Debug("init config and log succ, appconfig:%#v", proxyConf)

	err = initModel() 
	if err != nil {
		logs.Error("init model failed, err:%v", err)
		return
	}
	
	beego.Run()
}
