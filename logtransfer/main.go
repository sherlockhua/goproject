package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/sherlockhua/goproject/config"
	"encoding/json"
)

type AppConfig struct {
	logPath string
	logLevel string
	kafkaAddr string
	esAddr string
	esThreadNum int
}

var appConfig AppConfig

func initConfig(confPath string) (err error) {
	conf, err := config.NewConfig(confPath)
	if err != nil {
		return
	}

	logPath, err := conf.GetString("log_path")
	if len(logPath) == 0 || err != nil{
		return fmt.Errorf("get log_path failed,invalid logPath, err:%v", err)
	}

	appConfig.logPath = logPath
	logLevel, err := conf.GetString("log_level")
	if len(logLevel) == 0 || err != nil {
		return fmt.Errorf("get logLevel failed,invalid logLevel, err:%v", err)
	}

	appConfig.logLevel = logLevel

	kafkaAddr, err := conf.GetString("kafka_addr")
	if len(kafkaAddr) == 0 || err != nil {
		return fmt.Errorf("get kafkaAddr failed,invalid kafkaAddr, err:%v", err)
	}

	appConfig.kafkaAddr = kafkaAddr

	esAddr, err := conf.GetString("es_addr")
	if len(kafkaAddr) == 0 || err != nil {
		return fmt.Errorf("get es_addr failed,invalid es_addr, err:%v", err)
	}

	appConfig.esAddr = esAddr

	esThreadNum := conf.GetIntDefault("es_thread_num", 8)
	appConfig.esThreadNum = esThreadNum

	return
}

func getLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "trace":
		return logs.LevelTrace
	case "warn":
		return logs.LevelWarning
	case "info":
		return logs.LevelInformational
	case "error":
		return logs.LevelError
	default:
		return logs.LevelDebug
	}
}

func initLog(logPath string, logLevel string) (err error) {
	config := make(map[string]interface{})
	config["filename"] = logPath
	config["level"] = getLevel(logLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}


func main(){
	err := initConfig("./conf/app.conf")
	if err != nil {
		panic(fmt.Sprintf("init config failed, err:%v", err))
	}

	err = initLog(appConfig.logPath, appConfig.logLevel)
	if err != nil {
		panic(fmt.Sprintf("init log failed, err:%v", err))
	}

	logs.Debug("init log succ, config:%#v", appConfig)
	err = initKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	err = initES(appConfig.esAddr)
	if err!= nil {
		logs.Error("init es failed, err:%v", err)
		return
	}

	err = Run(appConfig.esThreadNum)
	if err != nil {
		logs.Error("run es failed, err:%v", err)
		return
	}
	logs.Debug("run exited")
}