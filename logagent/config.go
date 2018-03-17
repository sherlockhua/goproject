package main

import (
	"strings"
	"fmt"
	"github.com/sherlockhua/goproject/config"
)

type AppConfig struct {
	LogPath string
	LogLevel string
	kafkaAddr string
	LogFiles []string
}

var appConfig = &AppConfig{}

func initConfig(filename string) (err error) {
	conf, err := config.NewConfig(filename)
	if err != nil {
		return
	}

	logPath, err := conf.GetString("log_path")
	if err != nil || len(logPath) == 0{
		return
	}

	logLevel, err := conf.GetString("log_level")
	if err != nil || len(logPath) == 0{
		return
	}

	kafkaAddr, err := conf.GetString("kafka_addr")
	if err != nil || len(logPath) == 0{
		return
	}

	logFiles, err := conf.GetString("log_files")
	if err != nil || len(logPath) == 0{
		return
	}

	arr := strings.Split(logFiles, ",")
	for _, v := range arr {
		str := strings.TrimSpace(v)
		if len(str) == 0 {
			continue
		}

		appConfig.LogFiles = append(appConfig.LogFiles, str)
	}

	appConfig.kafkaAddr = kafkaAddr
	appConfig.LogLevel = logLevel
	appConfig.LogPath = logPath

	fmt.Printf("load config succ, data:%v\n", appConfig)
	return
}
