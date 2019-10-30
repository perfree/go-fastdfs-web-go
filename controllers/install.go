package controllers

import (
	"encoding/json"
	"fmt"
	"go-fastdfs-web-go/common"
	"go-fastdfs-web-go/models"
	"io/ioutil"
	"net/http"
	"regexp"
)

type InstallController struct {
	BaseController
}

// 安装页
func (c *InstallController) Get() {
	// 如存在用户,证明已安装过,直接跳转
	if models.GetUsesTotal() >= 1 {
		c.Redirect("/", 301)
	} else {
		c.TplName = "install.html"
	}
}

// 检查集群配置
func (c *InstallController) CheckServer() {
	peers := models.Peers{}
	if err := c.ParseForm(&peers); err != nil {
		c.Data["json"] = &JsonData{Code: FAIL_CODE, Count: 0, Msg: "参数解析失败", Data: nil}
		c.ServeJSON()
		return
	}
	if models.GetUsesTotal() >= 1 {
		c.Data["json"] = &JsonData{Code: FAIL_CODE, Count: 0, Msg: "您已安装,请直接登录", Data: nil}
		c.ServeJSON()
		return
	}
	if len(peers.ServerName) == 0 || len(peers.ServerName) > 100 {
		c.Data["json"] = &JsonData{Code: FAIL_CODE, Count: 0, Msg: "请正确填写集群名称(100字符以内)", Data: nil}
		c.ServeJSON()
		return
	}
	if len(peers.ServerAddress) == 0 || len(peers.ServerAddress) > 100 {
		c.Data["json"] = &JsonData{Code: FAIL_CODE, Count: 0, Msg: "请正确填写管理地址(100字符以内)", Data: nil}
		c.ServeJSON()
		return
	}
	match, _ := regexp.MatchString("[a-zA-z]+://[^\\s]*", peers.ServerAddress)
	if !match {
		c.Data["json"] = &JsonData{Code: FAIL_CODE, Count: 0, Msg: "管理地址格式不正确", Data: nil}
		c.ServeJSON()
		return
	}
	var urlPath = peers.ServerAddress
	if len(peers.GroupName) > 0 {
		urlPath += "/" + peers.GroupName
	}
	resp, err := http.Get(urlPath + common.STAT)
	if err != nil {
		c.Data["json"] = &JsonData{Code: FAIL_CODE, Count: 0, Msg: "连接go-fastdfs服务失败!请检查管理地址是否已配置白名单!", Data: nil}
		c.ServeJSON()
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	var tmp map[string]interface{}
	_ = json.Unmarshal([]byte(string(body)), &tmp)
	if tmp["status"] != "ok" {
		c.Data["json"] = &JsonData{Code: FAIL_CODE, Count: 0, Msg: "连接go-fastdfs服务失败!请检查管理地址是否已配置白名单!", Data: nil}
		c.ServeJSON()
		return
	}
	c.Data["json"] = &JsonData{Code: SUCCESS_CODE, Count: 0, Msg: "检查通过", Data: nil}
	c.ServeJSON()
}

// 安装
func (c *InstallController) DoInstall() {
	// TODO 安装处理
	u := models.User{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
	}
	fmt.Println(u.Peers.ServerName)
	fmt.Println(u.Password)
	c.ServeJSON()
}
