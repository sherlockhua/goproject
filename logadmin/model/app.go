package model

import (
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type AppInfo struct {
	AppId      int    `db:"app_id"`
	AppName    string `db:"app_name"`
	AppType    string `db:"app_type"`
	LogPath    string `db:"log_path"`
	CreateTime string `db:"create_time"`
	IPInfo     []IPInfo
}

type IPInfo struct {
	IP string `db:"ip"`
}

var Db *sqlx.DB

func InitDb(db *sqlx.DB) {

	Db = db
}

func GetAllAppInfo() (appInfo []AppInfo, err error) {

	err = Db.Select(&appInfo, "select app_id, app_name, log_path, app_type , create_time from tbl_app")
	if err != nil {
		logs.Warn("exec failed, ", err)
		return
	}

	return
}

func InsertAppInfo(appInfo *AppInfo) (id int64, err error) {

	conn, err := Db.Begin()
	if err != nil {
		logs.Warn("begin failed, err:%v", err)
		return
	}

	defer func() {
		if err != nil {
			conn.Rollback()
			return
		}

		conn.Commit()
	}()

	r, err := conn.Exec("insert into tbl_app(app_name, app_type, log_path)values(?, ?, ?)",
		appInfo.AppName, appInfo.AppType, appInfo.LogPath)
	if err != nil {
		logs.Warn("insert [%v] failed, err:%v", appInfo, err)
		return
	}

	id, err = r.LastInsertId()
	if err != nil {
		logs.Warn("get last id failed, err:%v", err)
		return
	}

	for i := 0; i < len(appInfo.IPInfo); i++ {
		_, err = conn.Exec("insert into tbl_app_ip(app_id, ip)values(?, ?)", id, appInfo.IPInfo[i].IP)
	}

	return
}

func GetAppInfo(appId int) (appInfo *AppInfo, err error) {

	appInfo = &AppInfo{}

	err = Db.Get(appInfo, "select app_id, app_name, app_type, log_path from tbl_app where app_id=?", appId)
	if err != nil {
		logs.Warn("get app info failed, id:%v err:%v", appId, err)
		return
	}

	err = Db.Select(&appInfo.IPInfo, "select ip from tbl_app_ip where app_id=?", appId)
	if err != nil {
		logs.Warn("get ip list failed, err:%v", err)
		return
	}

	return
}
