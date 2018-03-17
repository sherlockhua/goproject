package main

import (
	"sync"
	"github.com/hpcloud/tail"
	"github.com/astaxie/beego/logs"
	"fmt"
	"strings"
)

var waitGroup sync.WaitGroup

type TailObj struct {
	tail *tail.Tail
	offset int64
	filename string
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

func (t *TailMgr) AddLogFile(filename string) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	_, ok := t.tailObjMap[filename]
	if ok {
		err = fmt.Errorf("duplicate filename:%s", filename)
		return
	}

	tail, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
		})

	tailObj := &TailObj{
		filename:filename,
		offset:0,
		tail:tail,
	}

	t.tailObjMap[filename] = tailObj
	return
}

func (t *TailMgr) Process() {
	for _, tailObj := range t.tailObjMap {
		waitGroup.Add(1)
		go tailObj.readLog()
	}
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
	for _, filename := range appConfig.LogFiles {
		err := tailMgr.AddLogFile(filename)
		if err != nil {
			logs.Error("add log file %s failed, err:%v", filename, err)
			continue
		}
		logs.Debug("add log file %s succ", filename)
	}

	tailMgr.Process()
	waitGroup.Wait()
}