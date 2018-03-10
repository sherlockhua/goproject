package main

import (
	_ "beego_short_url/routers"
	"github.com/astaxie/beego"
	"github.com/sherlockhua/goproject/short_url/logic"
)

func initDb() (err error) {
	err = logic.InitDb(beego.AppConfig.String("Db::dns"))
	if err != nil {
		beego.Error("init db failed, err:%v", err)
		return
	}

	beego.Debug("init Db succ")
	return
}

func main() {
	err := initDb()
	if err != nil {
		beego.Error("init db failed, err:%v", err)
		return
	}
	beego.Run()
}

