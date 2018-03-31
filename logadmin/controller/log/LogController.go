package log

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/sherlockhua/goproject/logadmin/model"
	"strings"
)

type LogController struct {
	beego.Controller
}

func (p *LogController) ListApp() {

	logs.Debug("enter list app")

	appInfoList, err := model.GetAllAppInfo()
	if err != nil {
		logs.Warn("get app info failed, err:%v", err)
	}

	logs.Debug("get all log info:%v", appInfoList)
	p.Data["applist"] = appInfoList
	p.Layout = "layout/layout.html"
	p.TplName = "log/applist.html"
}


func (p *LogController) ListLog() {
	
		logs.Debug("enter list app")
	
		appInfoList, err := model.GetAllAppInfo()
		if err != nil {
			logs.Warn("get app info failed, err:%v", err)
		}
	
		var loglist []model.AppInfo
		for _, v := range appInfoList {
			if len(v.LogPath) == 0 {
				continue
			}

			loglist = append(loglist, v)
		}
		logs.Debug("get all log info:%v", appInfoList)
		p.Data["loglist"] = loglist
		p.Layout = "layout/layout.html"
		p.TplName = "log/log_list.html"
	}

func (p *LogController) CreateApp() {

	logs.Debug("enter list app")

	p.Layout = "layout/layout.html"
	p.TplName = "log/app.html"
}

func (p *LogController) SubmitApp() {

	logs.Debug("enter list app")

	appName := p.GetString("app_name")
	appType := p.GetString("app_type")
	appIP := p.GetString("app_ip")
	logPath := p.GetString("log_path")

	p.Layout = "layout/layout.html"
	p.TplName = "log/app.html"

	if len(appName) == 0 || len(appType) == 0 || len(appIP) == 0 || len(logPath) == 0 {
		p.Data["Error"] = "appName或appType或appIP参数不正确"
		p.TplName = "log/app_error.html"
		return
	}

	appInfo := &model.AppInfo{}
	appInfo.AppName = appName
	appInfo.AppType = appType
	appInfo.LogPath = logPath
	ips := strings.Split(appIP, ",")
	for _, v := range ips {
		ipInfo := model.IPInfo{}
		ipInfo.IP = v
		appInfo.IPInfo = append(appInfo.IPInfo, ipInfo)
	}

	_, err := model.InsertAppInfo(appInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("insert app info failed, err:%v", err)
		p.TplName = "log/app_error.html"
		return
	}

	p.Redirect("/app/list", 302)
}


func (p *LogController) SubmitEtcd() {
	
		logs.Debug("enter list app")
	
		p.Layout = "layout/layout.html"
		appId, err := p.GetInt("app_id")
		if appId == 0  || err != nil{
			logs.Error("app id:%d err:%v", appId, err)
			p.Data["Error"] = "appName或appType或appIP参数不正确"
			p.TplName = "log/app_error.html"
			return
		}
	
		appInfo, err := model.GetAppInfo(appId)
		if err != nil {
			p.Data["Error"] = fmt.Sprintf("insert app info failed, err:%v", err)
			p.TplName = "log/app_error.html"
			return
		}
	
		logs.Debug("appinfo is:%#v", appInfo)
		
		p.Redirect("/app/list", 302)
	}
	