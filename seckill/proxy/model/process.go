package model

import (
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"fmt"
	"github.com/sherlockhua/goproject/seckill/common"
	"time"
)

func initRecvThread(conf *common.SkillConf) (err error) {
	if (conf.RecvQueueThreadNum <= 0) {
		err = fmt.Errorf("invalid thread num:%d", conf.RecvQueueThreadNum)
		return 
	}

	for  i := 0; i < conf.RecvQueueThreadNum; i++ {
		go RecvHandle()
	}
	return
}


func initSendThread(conf *common.SkillConf) (err error) {

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

	for {
		conn := pool.Get()
		reply, err := conn.Do("BRPOP", secProxyData.conf.RecvQueueName, 0)
		dataArr, err := redis.Strings(reply, err)
		if err != nil {
			logs.Error("rpop from redis failed, err:%v", err)
			conn.Close()
			continue
		}

		if len(dataArr) != 2 {
			logs.Error("invalid data[%#v] rpop from redis failed, err:%v", dataArr, err)
			conn.Close()
			continue
		}

		data := dataArr[1]
		logs.Debug("rpop from redis:%s succ, data:%s", secProxyData.conf.RecvQueueName, data)
		var resp common.SecKillResult
		err = json.Unmarshal([]byte(data), &resp)
		if err != nil {
			logs.Error("unmarshal resp failed, err:%v data:%s", err, data)
			conn.Close()
			continue
		}

		req, err := secProxyData.GetRequest(resp.UserId, resp.ProductId)
		if err != nil {
			logs.Error("not found user, resp:%#v from request map", resp)
			conn.Close()
			continue
		}

		req.ResultChan <- &resp
		conn.Close()
	}
}

func SendHandle() {
	for req := range secProxyData.requestChan {
		
		req.CurTime = time.Now().UnixNano()
		data, err := json.Marshal(req)
		if err != nil {
			logs.Error("marshal data failed, err:%v data:%s", err, string(data))
			continue
		}

		conn := pool.Get()
		_, err = conn.Do("LPUSH", secProxyData.conf.SendQueueName, string(data))
		if err != nil {
			logs.Error("lpush to redis failed, err:%v", err)
			conn.Close()
			continue
		}

		conn.Close()
		logs.Debug("send to redis queue:%s succ, data:%s", secProxyData.conf.RecvQueueName, string(data))
	}
}