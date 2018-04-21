 package controller



import (
	//"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/sherlockhua/goproject/seckill/proxy/model"
	"github.com/sherlockhua/goproject/seckill/common"
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

	productId, err := p.GetInt("product_id", 0)
	if err != nil || productId == 0 {
		m = p.pack(1001, fmt.Sprintf("invalid product_id:%d", productId), nil)
		logs.Error("invalid product id:%d", productId)
		return
	}
	userId, err := p.GetInt64("user_id", 0)
	if err != nil || userId == 0 {
		m = p.pack(1001, fmt.Sprintf("invalid useId:%d", userId), nil)
		logs.Error("invalid useId :%d", userId)
		return
	}
	/*
	userIdStr := p.Ctx.GetCookie("UserId")
	if len(userIdStr) == 0 {
		m = p.pack(1003, fmt.Sprintf("invalid user_id:%s", userIdStr), nil)
		logs.Error("invalid user id:%d", product_id)
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil || userId == 0 {
		m = p.pack(1003, fmt.Sprintf("invalid user_id:%s", userIdStr), nil)
		logs.Error("invalid user id:%d", userId)
		return
	}*/

	data, err := model.SecKill(productId, userId, "")
	if err != nil {
		m = p.pack(1002, fmt.Sprintf("service busy, product_id:%d, uid:%d", 
			productId, userId), nil)
		logs.Error("service busy,product id:%d, err:%v user_id:%d", productId, err, userId)
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
		if data.Status == common.ActivitySaleOut {
			activityStatus.Status = common.ActivitySaleOut
		} else {
			activityStatus.Status = common.ActivityStart
		}
	} 

	if now < data.StartTime {
		activityStatus.Status = common.ActivityNotStart
	} 
	if now > data.EndTime {
		activityStatus.Status = common.ActivityEnd
	}

	m = p.pack(0, "success", activityStatus)
	return
}