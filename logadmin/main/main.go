package main

import (
	"github.com/astaxie/beego"
	_ "github.com/sherlockhua/goproject/logadmin/router"
)

func initBeego() {
	//	beego.SetStaticPath("/app", "/static/media")
	//beego.BConfig.WebConfig.StaticDir["/app/media"] = "app/media"
	//beego.SetStaticPath("/app/media/js", "views/media/js")
}

func main() {
	initBeego()

	err := initDb()
	if err != nil {
		return
	}

	beego.Run()
}
