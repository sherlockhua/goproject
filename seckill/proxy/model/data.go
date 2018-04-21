package model

import (
	"time"
	"sync"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"fmt"
	"errors"
	"github.com/sherlockhua/goproject/seckill/common"
)


var (
	ErrAlreadySaleout = errors.New("product sale out")
)

type SecProxyData struct {
	activityMap map[int]*common.Activity
	conf *common.SkillConf
	rwLock sync.RWMutex

	userMapLock sync.Mutex
	userRequestMap map[string] *common.SecKillRequest
	requestChan chan *common.SecKillRequest
}

var secProxyData *SecProxyData

func init() {
	NewSecProxyData()
}

func NewSecProxyData() {
	secProxyData = &SecProxyData {
		activityMap: make(map[int]*common.Activity, 128),
		requestChan: make(chan *common.SecKillRequest, 10000),
		userRequestMap: make(map[string]*common.SecKillRequest, 100000),
	}
}

func (s *SecProxyData) UpdateActivity(productArray []*common.Activity) (err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	//处理新增的活动配置
	for _, v := range productArray {
		s.activityMap[v.ProductId] = v
	}

	//处理删除的活动配置
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
		var activityArr []*common.Activity
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

func (s *SecProxyData) GetActivity(id int) (activity *common.Activity, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	activity, ok := s.activityMap[id]
	if !ok {
		err = fmt.Errorf("product_id:%d is not exists", id)
		return
	}

	return
}

func (s *SecProxyData) getKey(userId int64, productId int) string {
	return fmt.Sprintf("v1_%d_%d", userId, productId)
}

func (s *SecProxyData) GetRequest(userId int64, productId int) (req* common.SecKillRequest, err error) {
	var ok bool
	s.userMapLock.Lock()
	defer s.userMapLock.Unlock()

	key := s.getKey(userId, productId)
	req, ok = s.userRequestMap[key]
	if !ok {
		err = errors.New("user not exist")
		return
	}
	return
}

func (s *SecProxyData) AddRequest(req *common.SecKillRequest) (err error) {
	
	timer := time.NewTicker(2*time.Second)
	defer func() {
		timer.Stop()
	}()

	s.rwLock.RLock()
	v, ok := s.activityMap[req.ProductId]
	s.rwLock.RUnlock()

	if !ok {
		err = fmt.Errorf("product id %d not exists", req.ProductId)
		return
	}

	if v.Status == common.ActivitySaleOut {
		err = ErrAlreadySaleout
		return
	}

	key := s.getKey(req.UserId, req.ProductId)

	s.userMapLock.Lock()
	s.userRequestMap[key] = req
	s.userMapLock.Unlock()
	
	select {
	case s.requestChan <- req:
	case <-timer.C:
		err = fmt.Errorf("time out")
		return
	}
	return
}