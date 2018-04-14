package model


import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/logs"
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
	return
}

func loadProductInfo() (err error) {
	secProxyData = &SecProxyData {
		activityMap: make(map[int]*Activity, 128),
	}

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

	go secProxyData.Reload()
	return
}