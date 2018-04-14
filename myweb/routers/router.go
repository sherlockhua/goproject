package routers

import (
	"github.com/sherlockhua/goproject/myweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}, "*:UserInfo")
	beego.Router("/test", &controllers.UserController{}, "*:Test")
}
