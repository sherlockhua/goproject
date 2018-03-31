package router

import (
	"github.com/sherlockhua/goproject/logadmin/controller/log"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/app/list", &log.LogController{}, "*:ListApp")

	beego.Router("/log/list", &log.LogController{}, "*:ListLog")
	beego.Router("/app/create", &log.LogController{}, "*:CreateApp")
	beego.Router("/app/submit", &log.LogController{}, "*:SubmitApp")

	beego.Router("/etcd/submit", &log.LogController{}, "*:SubmitEtcd")
	beego.Router("/", &log.LogController{}, "*:ListApp")
}
