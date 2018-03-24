package main

import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
)

var client *clientv3.Client
var logConfChan chan string

func initEtcd(addr []string, keyfmt string, timeout time.Duration) (err error) {

	var key []string
	for _, ip := range ipArrays {
		key = append(key, fmt.Sprintf(keyfmt, ip))
	}

	logConfChan = make(chan string, 8)
	fmt.Println("etcd watch key:", key)

	client, err = clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: timeout,
	})
	if err != nil {
		logs.Warn("init etcd client failed, err:%v", err)
		return
	}
	
	logs.Debug("init etcd succ")
	waitGroup.Add(1)
	go WatchEtcd(key)
	return
}

func WatchEtcd(keys []string) {
	var watchChans []clientv3.WatchChan
	for _, key := range keys {
		rch := client.Watch(context.Background(), key)
		watchChans = append(watchChans, rch)
	}

	for {
		for _, watchC := range watchChans {
			select {
			case wresp := <- watchC:
				for _, ev := range wresp.Events {
					logs.Debug("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
					logConfChan <- string(ev.Kv.Value)
				}
			default:
			}
		}

		time.Sleep(time.Second)
	}

	waitGroup.Done()
}

func GetLogConf() chan string {
	return logConfChan
}