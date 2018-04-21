package main

import (
	"github.com/sherlockhua/goproject/seckill/common"
	//"github.com/astaxie/beego"
	//"fmt"
)

var logicConf *common.SkillConf

func initConfig() (err error){
	logicConf, err = common.ReadConfig()
	return
}