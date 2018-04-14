package model


import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var etcdClient *clientv3.Client
var productChan chan string

func initEtcd(conf *ModelConf) (err error) {

	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{conf.EtcdAddr},
		DialTimeout: 3 * time.Second,
	})

	if err != nil {
		logs.Warn("init etcd client failed, err:%v", err)
		return
	}

	productChan = make(chan string, 16)
	ctx, cancel := context.WithTimeout(context.Background(), 2 *time.Second)
	//   etcd_key /seckill/product/conf
	resp, err := etcdClient.Get(ctx, conf.EtcdProductKey)
	cancel()
	if err != nil {
		logs.Warn("get key %s failed, err:%v", conf.EtcdProductKey, err)
		return
	}

	for _, ev := range resp.Kvs {
		logs.Debug(" %q : %q\n",  ev.Key, ev.Value)
		productChan <- string(ev.Value)
	}
	
	go WatchEtcd(conf.EtcdProductKey)
	return
}

func WatchEtcd(key string) {
	
	for {
		rch := etcdClient.Watch(context.Background(), key)
		
		for resp := range rch {
			for _, ev := range resp.Events {
				logs.Debug("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				if ev.Type == clientv3.EventTypePut {
					productChan <- string(ev.Kv.Value)
				}
			}
		}
	}

}

func GetProductChan() chan string {
	return productChan
}