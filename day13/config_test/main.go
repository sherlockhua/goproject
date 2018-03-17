package main

import (
	"time"
	
	"fmt"
	"github.com/sherlockhua/goproject/config"
)

type AppConfig struct {
	port int
	kafkaAddr string
}

func (a *AppConfig) Callback(conf *config.Config) {
	port, err := conf.GetInt("server_port")
	if err != nil {
		fmt.Println("get port failed, err:", err)
		return
	}

	a.port = port
	fmt.Println("port:", a.port)

	kafkaAddr, err  := conf.GetString("kafka_addr")
	if err != nil {
		fmt.Println("get kafkaAddr failed, err:", err)
		return
	}

	a.kafkaAddr = kafkaAddr
	fmt.Println("kafkaAddr:", a.kafkaAddr)
}

var appConfig = &AppConfig {}

func main() {
	conf, err := config.NewConfig("./config.conf")
	if err != nil {
		fmt.Println("parse config failed")
		return
	}

	conf.AddNotifyer(appConfig)
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

	for {
		fmt.Println("port:", appConfig.port)
		fmt.Println("kafkaAddr:", appConfig.kafkaAddr)
		time.Sleep(5 * time.Second)
	}
}