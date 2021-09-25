package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"go-fastdfs-web-go/commons"
	"go-fastdfs-web-go/form"
	"go-fastdfs-web-go/models"
	"net/url"
	"regexp"
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
	var peers models.Peers
	err := c.ParseForm(&peers)
	if err != nil {
		c.ErrorJson(500, "param error", nil)
	}
	valid := validation.Validation{}
	valid.Required(peers.Name, "Name").Message("集群名称不能为空且在50字以内")
	valid.MaxSize(peers.Name, 50, "NameMax").Message("集群名称不能为空且在50字以内")

	valid.MaxSize(peers.GroupName, 50, "GroupNameMax").Message("组名称应在50字以内")

	valid.Required(peers.ServerAddress, "ServerAddress").Message("集群服务地址不能为空且在100字以内")
	valid.MaxSize(peers.ServerAddress, 100, "ServerAddressMax").Message("集群服务地址不能为空且在100字以内")

	valid.MaxSize(peers.ShowAddress, 100, "ShowAddressMax").Message("访问域名应在50字以内")

	urlRegexp := regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
	valid.Match(peers.ServerAddress, urlRegexp, "ServerAddressUrl").Message("请正确填写集群服务地址")

	if peers.ShowAddress != "" {
		valid.Match(peers.ShowAddress, urlRegexp, "ShowAddressUrl").Message("请正确填写访问域名")
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
			c.ErrorJson(500, err.Message, nil)
		}
	}

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
	var install = form.Install{}
	err := c.ParseForm(&install)
	if err != nil {
		c.ErrorJson(500, "param error", nil)
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&install)
	if err != nil {
		logs.Error("install -> ", err)
		c.ErrorJson(500, "安装失败", nil)
	}
	if !b {
		for _, err := range valid.Errors {
			logs.Error("install -> ", err.Key, err.Message)
			c.ErrorJson(500, err.Message, nil)
		}
	}

	peers := install.GetPeers()
	_, err = peers.Save()
	if err == nil {
		user := install.GetUser()
		user.PeersId = peers.Id
		_, err = user.Save()
		if err == nil {
			c.SuccessJson("安装成功")
		}
		c.ErrorJson(500, "安装失败", nil)
	}
	c.ErrorJson(500, "安装失败", nil)
}
