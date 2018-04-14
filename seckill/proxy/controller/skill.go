 package controller



import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/sherlockhua/goproject/seckill/proxy/model"
	"fmt"
)

type ProxyController struct {
	beego.Controller
}

func (p *ProxyController) SecKill() {
	logs.Debug("enter sec kill app")
	p.Data["json"] = "this is a SecKill bug"
	p.ServeJSON()
}

func (p *ProxyController) pack(errno int, message string,
	 data interface{}) map[string]interface{}{

	m := make(map[string]interface{},  16)
	m["errno"] = errno
	m["message"] = message
	m["data"] = data
	return m
}

func (p *ProxyController) SecInfo() {
	logs.Debug("enter sec info app")

	var m map[string]interface{} = make(map[string]interface{}, 16)
	defer func() {
		p.Data["json"] = m
		p.ServeJSON()
	}()

	product_id, err := p.GetInt("product_id", 0)
	if err != nil || product_id == 0 {
		m = p.pack(1001, fmt.Sprintf("invalid product_id:%d", product_id), nil)
		logs.Error("invalid product id:%d", product_id)
		return
	}
	
	data, err := model.SecInfo(product_id)
	if err != nil {
		m = p.pack(1002, fmt.Sprintf("service busy, product_id:%d", product_id), nil)
		logs.Error("service busy,product id:%d, err:%v", product_id, err)
		return 
	}

	m = p.pack(0, "success", data)
	return
}