package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/sherlockhua/goproject/seckill/common"
)

func main(){
	
	err := initConfig()
	if err != nil {
		panic(fmt.Sprintf("init config failed, err:%v", err))
	}

	err = common.InitLog(logicConf.LogPath, logicConf.LogLevel)
	if err != nil {
		panic(fmt.Sprintf("init logger failed, err:%v", err))
	}

	logs.Debug("init logger succ")
	err = initEtcd(logicConf)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}

	err = initRedis(logicConf)
	if err != nil {
		logs.Error("init redis failed, err:%v", err)
		return
	}

	skillLogic.logicConf = logicConf
	skillLogic.Reload()
	skillLogic.Run()
	beego.Run()
}

