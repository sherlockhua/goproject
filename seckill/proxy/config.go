package main

import (
	"github.com/astaxie/beego"
	"github.com/sherlockhua/goproject/seckill/proxy/model"
	"fmt"
)

var proxyConf model.ModelConf

func initConfig() (err error){
	redisAddr := beego.AppConfig.String("redis::redis_addr")
	if (len(redisAddr) == 0) {
		err = fmt.Errorf("invalid redis addr")
		return
	}

	sendQueue := beego.AppConfig.String("redis::send_request_queue_name")
	if (len(redisAddr) == 0) {
		err = fmt.Errorf("invalid redis addr")
		return
	}

	recvQueue := beego.AppConfig.String("redis::recv_resp_queue_name")
	if (len(redisAddr) == 0) {
		err = fmt.Errorf("invalid redis addr")
		return
	}

	logPath := beego.AppConfig.String("logs::log_path")
	if (len(logPath) == 0) {
		err = fmt.Errorf("invalid logPath")
		return
	}

	logLevel := beego.AppConfig.String("logs::log_level")
	if (len(logLevel) == 0) {
		err = fmt.Errorf("invalid log level")
		return
	}

	sendNum, err := beego.AppConfig.Int("redis::send_queue_thread_num")
	if (err != nil) {
		err = fmt.Errorf("invalid send_queue_thread_num")
		return
	}
	
	recvNum, err := beego.AppConfig.Int("redis::recv_queue_thread_num")
	if (err != nil) {
		err = fmt.Errorf("invalid recv_queue_thread_num")
		return
	}
	
	etcdAddr := beego.AppConfig.String("etcd::addr")
	if (len(etcdAddr) == 0) {
		err = fmt.Errorf("invalid etcdAddr")
		return
	}

	etcdProductKey := beego.AppConfig.String("etcd::product_key")
	if (len(etcdProductKey) == 0) {
		err = fmt.Errorf("invalid etcdProductKey")
		return
	}

	proxyConf.RedisAddr = redisAddr
	proxyConf.SendQueueName = sendQueue
	proxyConf.RecvQueueName = recvQueue

	proxyConf.LogPath = logPath
	proxyConf.LogLevel = logLevel

	proxyConf.RecvQueueThreadNum = recvNum
	proxyConf.SendQueueThreadNum = sendNum

	proxyConf.EtcdAddr = etcdAddr
	proxyConf.EtcdProductKey = etcdProductKey
	
	return
}