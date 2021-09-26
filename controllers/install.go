package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"go-fastdfs-web-go/commons"
	"go-fastdfs-web-go/form"
	"net/url"
)

var httpUtil = commons.HttpUtil{}

type InstallController struct {
	BaseController
}

// Get 安装页
func (c *InstallController) Get() {
	c.TplName = "install.tpl"
}

// CheckLocalServer 检查本机是否安装goFastDfs
func (c *InstallController) CheckLocalServer() {
	postValue := url.Values{"action": {"get"}}
	result, err := httpUtil.PostForm("http://127.0.0.1:8080/group1"+commons.ApiReload, postValue)
	if err != nil {
		result, err = httpUtil.PostForm("http://127.0.0.1:8080"+commons.ApiReload, postValue)
	}

	if err != nil {
		logs.Error("check -> ", err)
		c.ErrorJson(500, "error", nil)
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		logs.Error("check json exception -> ", err)
		c.ErrorJson(500, "error", nil)
	}
	logs.Info("check result -> ", resultMap)

	if resultMap["status"] == commons.ApiStatusSuccess {
		c.SuccessJson(resultMap)
	}
	c.ErrorJson(500, "error", nil)
}

// CheckServer 校验Server
func (c *InstallController) CheckServer() {
	var peers form.PeersForm
	err := c.ParseForm(&peers)
	if err != nil {
		c.ErrorJson(500, "param error", nil)
	}
	c.ValidParam(&peers, "安装失败")

	// 拼装url
	path := peers.ServerAddress
	if peers.GroupName != "" {
		path += "/" + peers.GroupName
	}
	path += commons.ApiStatus
	logs.Info("CheckServer url -> ", path)

	// 测试连接GoFastDfs
	result, err := httpUtil.PostForm(path, nil)
	if err != nil {
		logs.Error("CheckServer -> ", err)
		c.ErrorJson(500, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!", nil)
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		logs.Error("CheckServer json exception -> ", err)
		c.ErrorJson(500, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!", nil)
	}
	logs.Info("CheckServer result -> ", resultMap)

	if resultMap["status"] == commons.ApiStatusSuccess {
		c.SuccessJson(resultMap)
	}
	c.ErrorJson(500, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!", nil)
}

// DoInstall 安装
func (c *InstallController) DoInstall() {
	var install = form.InstallForm{}
	err := c.ParseForm(&install)
	if err != nil {
		c.ErrorJson(500, "param error", nil)
	}
	c.ValidParam(&install, "安装失败")

	peers := install.GetPeers()
	_, err = peersDao.Save(peers)
	if err == nil {
		user := install.GetUser()
		user.PeersId = peers.Id
		_, err = userDao.Save(user)
		if err == nil {
			c.SuccessJson("安装成功")
		}
		c.ErrorJson(500, "安装失败", nil)
	}
	c.ErrorJson(500, "安装失败", nil)
}
