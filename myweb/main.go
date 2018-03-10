package main

import (
	"encoding/json"
	_ "github.com/sherlockhua/goproject/myweb/routers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	beego.AddTemplateExt("xxx")
	beego.SetStaticPath("/code", "code")

	beego.BConfig.WebConfig.DirectoryIndex = true
	host := beego.AppConfig.String("dbconfig::mysqlhost")
	port, _ := beego.AppConfig.Int("dbconfig::mysqlport")
	//beego.AppConfig.
	redisPort := beego.AppConfig.DefaultInt("dbconfig::redisport", 10000)
	fmt.Println(host, port, redisPort)

	//日志配置
	m := make(map[string]interface{})
	m["filename"] = "./logs/test.log"
	config, _ := json.Marshal(m)

	//file代表打印到文件中
	//console代表打印到终端
	beego.SetLogger("file", string(config))
	beego.Debug("service init success, ready to running")
	beego.BeeLogger.DelLogger("console")
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	beego.Run()
	
}

