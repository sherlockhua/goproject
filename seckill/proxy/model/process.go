package model

import (
	"fmt"
)

func initRecvThread(conf *ModelConf) (err error) {
	if (conf.RecvQueueThreadNum <= 0) {
		err = fmt.Errorf("invalid thread num:%d", conf.RecvQueueThreadNum)
		return 
	}

	for  i := 0; i < conf.RecvQueueThreadNum; i++ {
		go RecvHandle()
	}
	return
}


func initSendThread(conf *ModelConf) (err error) {

	if (conf.SendQueueThreadNum <= 0) {
		err = fmt.Errorf("invalid thread num:%d", conf.SendQueueThreadNum)
		return 
	}

	for  i := 0; i < conf.SendQueueThreadNum; i++ {
		go SendHandle()
	}
	return
}

func RecvHandle() {

}

func SendHandle() {

}