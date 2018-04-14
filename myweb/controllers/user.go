package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) UserInfo() {
	/*
	username := c.GetString("username")
	fmt.Println(username)

	age, _ := c.GetInt("age")
	fmt.Println(age)
	*/
	//需要在app.conf设置： copyrequestbody = true
	//http://192.168.14.200:9999/static/post.exe
	//http://192.168.14.200:9999/static/fiddler.exe
	beego.Debug("client send data:%s\n", string(c.Ctx.Input.RequestBody))
	var m  map[string]interface{} = make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	beego.Debug("username:%v\n", m["username"])
	beego.Debug("passwd:%v\n", m["passwd"])
	beego.Error("this is process failed, err:xxxx")

	c.Data["hello"] = "oldboy.me"
	c.TplName = "user/index.xxx"
}


func (c *UserController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


func (c *UserController) Delete() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type Request struct {
	userId int
	resultChan chan string
}


func (c *UserController) Test () {

	req := &Request{
		userId:1000,
		resultChan: make(chan string, 1),
	}

	go TestProc(req)
	timer := time.NewTicker(5*time.Second)

	select {
	case result := <- req.resultChan:
		c.Data["json"] = result
	case <- timer.C:
		c.Data["json"] = "timeout"
	}

	timer.Stop()
	c.ServeJSON()

}

func TestProc(req* Request) {

	time.Sleep(6 *time.Second)
	req.resultChan <- "process succ"
}