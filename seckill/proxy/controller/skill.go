 package controller



import (
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/sherlockhua/goproject/seckill/proxy/model"
	"fmt"
	"time"
)

type ActivityStatus struct {
	ProductId int
	StartTime int64
	EndTime int64
	Status int
}

type ProxyController struct {
	beego.Controller
}

func (p *ProxyController) SecKill() {
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
	
	userIdStr := p.Ctx.GetCookie("UserId")
	if len(userIdStr) == 0 {
		m = p.pack(1003, fmt.Sprintf("invalid user_id:%s", userIdStr), nil)
		logs.Error("invalid user id:%d", product_id)
		return
	}

	user_id, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil || user_id == 0 {
		m = p.pack(1003, fmt.Sprintf("invalid user_id:%s", userIdStr), nil)
		logs.Error("invalid user id:%d", product_id)
		return
	}

	data, err := model.SecInfo(product_id, user_id)
	if err != nil {
		m = p.pack(1002, fmt.Sprintf("service busy, product_id:%d, uid:%d", 
			product_id, user_id), nil)
		logs.Error("service busy,product id:%d, err:%v user_id:%d", product_id, err, user_id)
		return 
	}

	m = p.pack(0, "success", data)
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

	var activityStatus ActivityStatus
	activityStatus.StartTime = data.StartTime
	activityStatus.EndTime = data.EndTime
	activityStatus.ProductId = data.ProductId

	now := time.Now().Unix()
	if now >= data.StartTime && now < data.EndTime {
		if data.Status == model.ActivitySaleOut {
			activityStatus.Status = model.ActivitySaleOut
		} else {
			activityStatus.Status = model.ActivityStart
		}
	} 

	if now < data.StartTime {
		activityStatus.Status = model.ActivityNotStart
	} 
	if now > data.EndTime {
		activityStatus.Status = model.ActivityEnd
	}

	m = p.pack(0, "success", activityStatus)
	return
}