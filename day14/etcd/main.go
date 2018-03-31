package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var logconf = `
[
    {
        "topic":"nginx_log",
        "log_path":"D:/opensource/nginx-1.13.10/logs/error.log",
        "service":"account",
        "send_rate":50000
	},
	{
        "topic":"nginx_log",
        "log_path":"D:/opensource/nginx-1.13.10/logs/access.log",
        "service":"nginx",
        "send_rate":50000
	}
]
`

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/192.168.12.3/log_config", logconf)
	_, err = cli.Put(ctx, "/logagent/conf/b", "sample_value1")
	_, err = cli.Put(ctx, "/logagent/conf/c", "sample_value2")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/logagent/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
