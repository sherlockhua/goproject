package routers

import (
	"beego_short_url/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	//http.HandleFunc("/trans/long2short", Long2Short)
	//http.HandleFunc("/trans/short2long", Short2Long)
	beego.Router("/trans/long2short", &controllers.ShortUrlController{}, "post:Long2Short")
	beego.Router("/trans/short2long", &controllers.ShortUrlController{}, "post:Short2Long")
	beego.Router("/shorturl", &controllers.ShortUrlController{}, "get:ShortUrlList")
	beego.Router("/jump", &controllers.ShortUrlController{}, "get:Jump")
}
