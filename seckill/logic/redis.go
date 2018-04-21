package main


import (
	"github.com/astaxie/beego/logs"
	"github.com/sherlockhua/goproject/seckill/common"
    "github.com/garyburd/redigo/redis"
    "time"
)

var pool *redis.Pool

//初始化一个pool
func newPool(server, password string) *redis.Pool {
    return &redis.Pool{
        MaxIdle:     64,
        MaxActive:   1000,
        IdleTimeout: 240 * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                return nil, err
			}
			/*
            if _, err := c.Do("AUTH", password); err != nil {
                c.Close()
                return nil, err
            }*/
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if time.Since(t) < time.Minute {
                return nil
            }
            _, err := c.Do("PING")
            return err
        },
    }
}

func initRedis(conf *common.SkillConf) (err error) {
	pool = newPool(conf.RedisAddr, conf.RedisPasswd)
	conn := pool.Get()
	defer conn.Close()

	_, err = conn.Do("PING")
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
		return
	}

	logs.Debug("connect to redis succ")
	return
}