package main

import (
	"sync"
	"github.com/sherlockhua/goproject/seckill/common"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"fmt"
	"crypto/md5"
	"encoding/base64"
	"errors"
)

type SkillLogic struct {
	rwLock sync.RWMutex
	activityMap map[int]*common.Activity

	randRate float32
	salt string
	logicConf *common.SkillConf
	reqChan chan *common.SecKillRequest
	respChan chan *common.SecKillResult
}

var skillLogic  = &SkillLogic{
	activityMap: make(map[int]*common.Activity, 10000),
	reqChan: make(chan *common.SecKillRequest, 10000),
	respChan: make(chan *common.SecKillResult, 10000),
	randRate: 1.0,
	salt: "vbTiuz0f8zmY4aLbiuoLJW5yVIdifP50",
}

func (s *SkillLogic) Reload() {
	go func() {
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
	}()	
	return
}


func (s *SkillLogic) UpdateActivity(productArray []*common.Activity) (err error) {
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

func (s *SkillLogic) Run(){
	for i := 0; i < logicConf.RecvQueueThreadNum; i++{
		go s.RecvHandle()
	}

	for i := 0; i < logicConf.SendQueueThreadNum; i++{
		go s.SendHandle()
	}

	for i := 0; i < 16; i++{
		go s.Handle()
	}
}

func (s *SkillLogic) RecvHandle() {
	for {
		conn := pool.Get()
		reply, err := conn.Do("BRPOP", s.logicConf.SendQueueName, 0)
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
		req := &common.SecKillRequest{}
		err = json.Unmarshal([]byte(data), req)
		if err != nil {
			logs.Error("Unmarshal data[%#v] rpop from redis failed, err:%v", dataArr, err)
			conn.Close()
			continue
		}

		select {
		case s.reqChan <- req:
		default:
		}
	}
}


func (s *SkillLogic) SendHandle() {
	
	for resp := range s.respChan {
		data, err := json.Marshal(resp)
		if err != nil {
			logs.Error("marshal failed, err:%v resp:%v", err, resp)
			continue
		}

		conn := pool.Get()
		_, err = conn.Do("LPUSH", s.logicConf.RecvQueueName, string(data))
		if err != nil {
			logs.Error("lpush to redis failed, err:%v", err)
			conn.Close()
			continue
		}
		conn.Close()
	}
}

func (s *SkillLogic) Handle(){

	for req := range s.reqChan {
		err := s.isSaleout(req.ProductId)
		if err != nil {
			resp := &common.SecKillResult {
				UserId: req.UserId,
				ProductId: req.ProductId,
				Status: common.ActivitySaleOut,
				CurTime:req.CurTime,
			}

			s.respChan <- resp
			continue
		}

		num := rand.Intn(100)
		maxRate := int(s.randRate*100.0)
		if num >= maxRate {
			resp := &common.SecKillResult {
				UserId: req.UserId,
				ProductId: req.ProductId,
				Status: common.ActivityRetry,
				CurTime:req.CurTime,
			}

			s.respChan <- resp
			continue
		}


		resp := &common.SecKillResult {
			UserId: req.UserId,
			ProductId: req.ProductId,
			Status: common.ActivitySucc,
			Token: s.getToken(req),
			CurTime:req.CurTime,
		}

		s.respChan <- resp
	}
}

func (s *SkillLogic) getToken(req *common.SecKillRequest) (token string) {
	data := fmt.Sprintf("pid_%d_uid_%d_salt_%s_time_%d", 
		req.ProductId, req.UserId, s.salt, req.CurTime)
	sum := md5.Sum([]byte(data))
	
	token = base64.StdEncoding.EncodeToString(sum[:])
	return
}

func (s *SkillLogic) isSaleout(productId int) (err error) {
	s.rwLock.RLock()
	

	product, ok := s.activityMap[productId]
	if !ok {
		err = errors.New("product not exist")
		logs.Error("product:%d not exist", productId)
		s.rwLock.RUnlock()
		return
	}

	if product.Status == common.ActivitySaleOut {
		err = errors.New("sale out")
		logs.Debug("product:%d sale out", productId)
		s.rwLock.RUnlock()
		return
	}

	s.rwLock.RUnlock()
	
	key := fmt.Sprintf("pid_%d", productId)
	conn := pool.Get()
	curVal, err := redis.Int(conn.Do("INCR", key))
	if err != nil {
		logs.Error("incr failed, err:%v", err)
		err = errors.New("redis err")
		return
	}

	logs.Debug("product:%d curVal:%d total:%d", productId, curVal, product.Count)
	if curVal >= product.Count {
		err = errors.New("sale out")
		return
	}
	return
}