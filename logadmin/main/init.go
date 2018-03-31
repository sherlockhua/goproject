package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
	"github.com/sherlockhua/goproject/logadmin/model"
)

func initDb() (err error) {

	database, err := sqlx.Open("mysql", "root:@tcp(10.0.0.200:3306)/logagent")
	if err != nil {
		logs.Warn("open mysql failed,", err)
		return
	}

	err = database.Ping()
	if err != nil {
		logs.Error("connect to mysql failed, err:%v", err)
		return
	}
	
	model.InitDb(database)
	return
}
