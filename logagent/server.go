package main

import (
	"sync"
	"github.com/hpcloud/tail"
	"github.com/astaxie/beego/logs"
	"fmt"
	"encoding/json"
	"strings"
)

var waitGroup sync.WaitGroup

type TailObj struct {
	tail *tail.Tail
	offset int64
	/*
	filename string
	service string
	sendRate int
	*/
	logConf LogConfig
}

type TailMgr struct {
	tailObjMap map[string]*TailObj
	lock sync.Mutex
}

var tailMgr *TailMgr

func NewTailMgr() (*TailMgr) {
	return  &TailMgr {
		tailObjMap:make(map[string]*TailObj, 16),
	}
}

func (t *TailMgr) AddLogFile(conf LogConfig) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	_, ok := t.tailObjMap[conf.LogPath]
	if ok {
		err = fmt.Errorf("duplicate filename:%s", conf.LogPath)
		return
	}

	tail, err := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
		})

	tailObj := &TailObj{
		logConf: conf,
		offset:0,
		tail:tail,
	}

	t.tailObjMap[conf.LogPath] = tailObj
	go tailObj.readLog()
	return
}

func (t *TailMgr) reloadConfig(logConfArr []LogConfig) (err error) {

	for _, conf := range logConfArr {
		tailObj, ok := t.tailObjMap[conf.LogPath]
		if !ok {
			err = t.AddLogFile(conf)
			if err != nil {
				logs.Error("add log file failed, err:%v", err)
				continue
			}
			continue
		}
		
		tailObj.logConf = conf
		t.tailObjMap[conf.LogPath] = tailObj
	}
	return
}


func (t *TailMgr) Process() {
	for conf := range GetLogConf() {
		logs.Debug("log conf:%v", conf)
		var logConfArr []LogConfig
		err := json.Unmarshal([]byte(conf), &logConfArr)
		if err != nil {
			logs.Error("unmarshal failed, err:%v conf:%s", err, conf)
			continue
		}

		err = t.reloadConfig(logConfArr)
		if err != nil {
			logs.Error("reload config from etcd failed, err:%v", err)
			continue
		}

		logs.Debug("reload from etcd succ, config:%v", logConfArr)
	}
	/*
	for _, tailObj := range t.tailObjMap {
		waitGroup.Add(1)
		go tailObj.readLog()
	}
	*/
}

func (t *TailObj) readLog() {
	for line := range t.tail.Lines {
		if line.Err != nil{
			logs.Error("read line failed, err:%v", line.Err)
			continue
		}

		str := strings.TrimSpace(line.Text)
		if (len(str) == 0 || str[0] == '\n') {
			continue
		}
		kafkaSender.addMessage(line.Text)
	}
	waitGroup.Done()
}


func RunServer() {
	tailMgr = NewTailMgr()
	/*
	var logfiles []string
	for _, filename := range logfiles {
		err := tailMgr.AddLogFile(filename)
		if err != nil {
			logs.Error("add log file %s failed, err:%v", filename, err)
			continue
		}
		logs.Debug("add log file %s succ", filename)
	}
	*/
	tailMgr.Process()
	waitGroup.Wait()
}