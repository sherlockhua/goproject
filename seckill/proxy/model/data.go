package model

import (
	"sync"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)



type Activity struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
}

type SecProxyData struct {
	activityMap map[int]*Activity
	rwLock sync.RWMutex
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