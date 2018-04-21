package model


import (
	"fmt"
	"time"
	
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/logs"
	"github.com/sherlockhua/goproject/seckill/common"
	//"time"
)



//声明一些全局变量
var (
    pool          *redis.Pool
)

func initRedis(conf *common.SkillConf) (err error) {
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

func Init(conf *common.SkillConf) (err error) {

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

	secProxyData.conf = conf
	return
}

func SecInfo(product_id int) (a *common.Activity, err error) {
	a, err = secProxyData.GetActivity(product_id)
	if err != nil {
		return
	}
	return
}


func SecKill(productId int, userId int64, userIp string) (result *common.SecKillResult, err error) {
	
	req := &common.SecKillRequest{
		ProductId: productId,
		UserId: userId,
		UserIp: userIp,
		ResultChan: make(chan *common.SecKillResult, 1),
	}

	err = secProxyData.AddRequest(req)
	if err != nil {
		if err == ErrAlreadySaleout {
			result = &common.SecKillResult{
				UserId: userId,
				ProductId:productId,
				Status: common.ActivitySaleOut,
			}
			err = nil
			return
		}
		return
	}

	timer := time.NewTicker(10 *time.Second)
	select {
	case <- timer.C:
		err = fmt.Errorf("timeout")
	case result = <- req.ResultChan:
		return
	}
	return
}

func loadProductInfo() (err error) {
	
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