package model

import (
	"time"
	"sync"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"fmt"
)

const (
	ActivityStart   = 1
	ActivityEnd     = 2
	ActivitySaleOut = 3
	ActivityNotStart = 4
)

type Activity struct {
	ProductId int    `json:"product_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64	 `json:"end_time"`
	Status    int	 `json:"status"`
}

type SecKillResult struct {
	UserId int
	ProductId int
	Token string
	Status int
}

type SecKillRequest struct {
	UserId int
	ProductId int
	UserIp string
	resultChan chan *SecKillResult
}

type SecProxyData struct {
	activityMap map[int]*Activity
	rwLock sync.RWMutex
	requestChan chan *SecKillRequest
}

var secProxyData *SecProxyData

func (s *SecProxyData) UpdateActivity(productArray []*Activity) (err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	for _, v := range productArray {
		s.activityMap[v.ProductId] = v
	}

	for k , _ := range s.activityMap {
		found := false
		for _, v2 := range productArray {
			if v2.ProductId == k {
				found = true
				break
			}
		}

		if found == false {
			delete(s.activityMap, k)
		}
	}
	return
}

func (s *SecProxyData) Reload() {
	for productConf := range  GetProductChan() {
		var activityArr []*Activity
		err := json.Unmarshal([]byte(productConf), &activityArr)
		if err != nil {
			logs.Error("unmarshal failed, err:%v, conf:%s", err, productConf)
			continue
		}
	
		err = s.UpdateActivity(activityArr)
		if err != nil {
			logs.Error("UpdateActivity failed, err:%v, conf:%s", err, productConf)
			continue
		}

		logs.Debug("reload conf from etcd succ, new conf:%s", productConf)
	}
	return
}

func (s *SecProxyData) GetActivity(id int) (activity *Activity, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	activity, ok := s.activityMap[id]
	if !ok {
		err = fmt.Errorf("product_id:%d is not exists", id)
		return
	}

	return
}


func (s *SecProxyData) AddRequest(req *SecKillRequest) (err error) {
	
	timer := time.NewTicker(2*time.Second)
	defer func() {
		timer.Stop()
	}()

	select {
	case s.requestChan <- req:
	case <-timer.C:
		err = fmt.Errorf("time out")
		return
	}
	return
}