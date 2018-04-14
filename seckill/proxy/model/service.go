package model


import (
	"fmt"
	"time"
	
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/logs"
	//"time"
)


type ModelConf struct {
	RedisAddr string
	RedisPasswd string
	
	EtcdAddr string
	EtcdProductKey string

	SendQueueName string
	RecvQueueName string

	SendQueueThreadNum int
	RecvQueueThreadNum  int

	LogPath string
	LogLevel string
}

//声明一些全局变量
var (
    pool          *redis.Pool
)

func initRedis(conf *ModelConf) (err error) {
	pool = newPool(conf.RedisAddr, conf.RedisPasswd)
	conn := pool.Get()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("connect to redis failed, err:%v", err)
		return
	}

	logs.Debug("connect to redis succ")
	return
}

func Init(conf *ModelConf) (err error) {

	err = initRedis(conf)
	if err != nil {
		return
	}

	err = initEtcd(conf)
	if err != nil {
		return 
	}

	err = initRecvThread(conf)
	if err != nil {
		return
	}

	err = initSendThread(conf)
	if err != nil {
		return
	}

	err = loadProductInfo()
	if err != nil {
		return
	}
	return
}

func SecInfo(product_id int) (a *Activity, err error) {
	a, err = secProxyData.GetActivity(product_id)
	if err != nil {
		return
	}
	return
}


func SecKill(productId, userId int, userIp string) (result *SecKillResult, err error) {
	
	req := &SecKillRequest{
		ProductId: productId,
		UserId: userId,
		UserIp: userIp,
		resultChan: make(chan *SecKillResult, 1),
	}

	err = secProxyData.AddRequest(req)
	if err != nil {
		logs.Error("add request failed, err:%v", err)
		return
	}

	timer := time.NewTicker(10 *time.Second)
	select {
	case <- timer.C:
		err = fmt.Errorf("timeout")
	case result = <- req.resultChan:
		return
	}
	return
}

func loadProductInfo() (err error) {
	secProxyData = &SecProxyData {
		activityMap: make(map[int]*Activity, 128),
		requestChan: make(chan *SecKillRequest, 10000),
	}
/*
	productConf := <- GetProductChan()

	var activityArr []*Activity
	err = json.Unmarshal([]byte(productConf), &activityArr)
	if err != nil {
		return
	}

	err = secProxyData.UpdateActivity(activityArr)
	if err != nil {
		return
	}
*/
	go secProxyData.Reload()
	return
}