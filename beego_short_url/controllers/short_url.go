package controllers

import (
	///"net/http"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego"
	 "github.com/sherlockhua/goproject/short_url/logic"
	"github.com/sherlockhua/goproject/short_url/model"
)

type ShortUrlController struct {
	beego.Controller
}


func (c *ShortUrlController) Jump() {
	shortUrl := c.GetString("short_url")
	if len(shortUrl) == 0 {
		///http.Redirect(c.Ctx.Inpu)
		return
	}

	var req model.Short2LongRequest
	var resp  = &model.Short2LongResponse{}

	defer func() {
		if err := recover(); err != nil {
			beego.Error("panic err:%v", err)
			return
		}
	}()

	req.ShortUrl = shortUrl
	resp, err := logic.Short2Long(&req)
	if err != nil {
		beego.Error("Short2Long failed, err:%v", err)
		return
	}

	fmt.Println(resp)
	beego.Debug("origin url:%s short url:%s", resp.OriginUrl, shortUrl)
	c.Redirect(resp.OriginUrl, 301)
}

func (c *ShortUrlController) ShortUrlList() {
	limit, err := c.GetInt("limit")
	if err != nil {
		beego.Warn("not have limit params, use default 10")
		limit = 10
	}

	data, err := logic.GetLastShortUrl(limit)
	if err != nil {
		beego.Error("get url list failed, err:%v", err)

	}

	for i, v := range data {
		v.ShortUrl = fmt.Sprintf("/jump/?short_url=%s", v.ShortUrl)
		data[i] = v
	}
	c.Data["url_list"] = data
	c.TplName = "index.tpl"
}

func (c *ShortUrlController) Long2Short() {
	var req model.Long2ShortRequest
	var resp *model.Long2ShortResponse = &model.Long2ShortResponse{}

	defer func() {
		if err := recover(); err != nil {
			beego.Error("panic err:%v", err)
			resp.Code = 1005
			resp.Message = "server is busy"
			c.Data["json"] = resp
			c.ServeJSON()
		}
	}()

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		beego.Error("unmarshal failed, err:%v", err)
		resp.Code = 1001
		resp.Message = "json unmarshal failed"
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	resp, err = logic.Long2Short(&req)
	if err != nil {
		beego.Error("Long2Short failed, err:%v", err)
		resp.Code = 1002
		resp.Message = "json unmarshal failed"
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}


func (c *ShortUrlController) Short2Long() {
	var req model.Short2LongRequest
	var resp  = &model.Short2LongResponse{}

	defer func() {
		if err := recover(); err != nil {
			beego.Error("panic err:%v", err)
			resp.Code = 1005
			resp.Message = "server is busy"
			c.Data["json"] = resp
			c.ServeJSON()
		}
	}()

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		beego.Error("unmarshal failed, err:%v", err)
		resp.Code = 1001
		resp.Message = "json unmarshal failed"
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	resp, err = logic.Short2Long(&req)
	if err != nil {
		beego.Error("Short2Long failed, err:%v", err)
		resp.Code = 1002
		resp.Message = "json unmarshal failed"
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
}
