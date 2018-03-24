package main

import (
	"time"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"fmt"
)

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

func initLog() (err error) {
	config := make(map[string]interface{})
	config["filename"] = appConfig.LogPath
	config["level"] = getLevel(appConfig.LogLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}

func main() {
	
	err := initConfig("./conf/app.conf")
	if err != nil {
		panic(fmt.Sprintf("init config failed, err:%v", err))
	}

	err = initLog()
	if err != nil {
		return
	}

	logs.Debug("init log succ")
	ipArrays, err = getLocalIP()
	if err != nil {
		logs.Error("get local ip failed, err:%v", err)
		return
	}

	logs.Debug("get local ip succ, ips:%v", ipArrays)
	err = initKafka()
	if err != nil {
		logs.Error("init kafka faild, err:%v", err)
		return
	}

	err = initEtcd(appConfig.etcdAddr, appConfig.etcdWatchKeyFmt,
		 time.Duration(appConfig.etcdTimeout)*time.Millisecond)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}

	RunServer()
}