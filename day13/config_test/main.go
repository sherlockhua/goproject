package main

import (
	"time"
	
	"fmt"
	"sync/atomic"
	"github.com/sherlockhua/goproject/config"
)

type AppConfig struct {
	port int
	kafkaAddr string
}

type AppConfigMgr struct {
	config atomic.Value
}

var appConfigMgr = &AppConfigMgr{}

func (a *AppConfigMgr) Callback(conf *config.Config) {
	var appConfig = &AppConfig {}

	port, err := conf.GetInt("server_port")
	if err != nil {
		fmt.Println("get port failed, err:", err)
		return
	}

	appConfig.port = port
	fmt.Println("port:", appConfig.port)

	kafkaAddr, err  := conf.GetString("kafka_addr")
	if err != nil {
		fmt.Println("get kafkaAddr failed, err:", err)
		return
	}

	appConfig.kafkaAddr = kafkaAddr
	fmt.Println("kafkaAddr:", appConfig.kafkaAddr)

	appConfigMgr.config.Store(appConfig)
}

func run() {
	for {
		appConfig := appConfigMgr.config.Load().(*AppConfig)
		fmt.Println("port:", appConfig.port)
		fmt.Println("kafkaAddr:", appConfig.kafkaAddr)
		fmt.Println()
		fmt.Println()
		time.Sleep(5 * time.Second)
	}
}

func main() {
	conf, err := config.NewConfig("./config.conf")
	if err != nil {
		fmt.Println("parse config failed")
		return
	}

	conf.AddNotifyer(appConfigMgr)

	var appConfig = &AppConfig {}
	appConfig.port , err = conf.GetInt("server_port")
	if err != nil {
		fmt.Println("get port failed, err:", err)
		return
	}

	fmt.Println("port:", appConfig.port)

	appConfig.kafkaAddr, err  = conf.GetString("kafka_addr")
	if err != nil {
		fmt.Println("get kafkaAddr failed, err:", err)
		return
	}

	fmt.Println("kafkaAddr:", appConfig.kafkaAddr)
	appConfigMgr.config.Store(appConfig)

	run()
}