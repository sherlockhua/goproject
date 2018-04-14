package main

import (
	"github.com/astaxie/beego/logs"
	"encoding/json"
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

func initLog(logPath string, logLevel string) (err error) {
	config := make(map[string]interface{})
	config["filename"] = logPath
	config["level"] = getLevel(logLevel)
	configStr, err := json.Marshal(config)
	if err != nil {
		return
	}
	
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}