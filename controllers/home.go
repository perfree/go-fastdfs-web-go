package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go-fastdfs-web-go/commons"
	"go-fastdfs-web-go/models"
	"net/url"
	"runtime"
	"strconv"
	"time"
)

type HomeController struct {
	BaseController
}

// Home 后台首页
func (c *HomeController) Home() {
	c.Data["osName"] = runtime.GOOS
	c.Data["osArch"] = runtime.GOARCH
	c.Data["version"] = beego.AppConfig.String("AppVer")
	c.Data["versionDate"] = beego.AppConfig.String("AppVerDate")
	c.TplName = "home.tpl"
}

// GetStatus 获取状态信息
func (c *HomeController) GetStatus() {
	path, err := c.GetPeersUrl()
	if err != nil {
		c.ErrorJson(500, "系统异常", nil)
	}

	result, err := httpUtil.PostForm(path+commons.ApiStatus, nil)
	if err != nil {
		logs.Error("GetStatus -> ", err)
		c.ErrorJson(500, "调取go-fastDfs接口失败", nil)
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		logs.Error("GetStatus json exception -> ", err)
		c.ErrorJson(500, "调取go-fastDfs接口失败", nil)
	}
	logs.Info("GetStatus result -> ", resultMap)

	if resultMap["status"] == commons.ApiStatusSuccess {
		data := resultMap["data"].(map[string]interface{})
		result, err := commons.GetStatus(data)
		if err != nil {
			c.ErrorJson(500, "数据解析错误", nil)
		}
		c.SuccessJson(result)
	}

	c.ErrorJson(500, "调取go-fastDfs接口失败", nil)
}

// RepairStat 修正统计信息(30天)
func (c *HomeController) RepairStat() {
	count := 0
	for i := 0; i > -30; i-- {
		beforeDate := time.Now().AddDate(0, 0, i).Format("20060102")
		postValue := url.Values{"date": {beforeDate}}
		data := make(map[string]interface{})
		data["date"] = beforeDate
		path, _ := c.GetPeersUrl()
		result, err := httpUtil.PostForm(path+commons.ApiRepairStat, postValue)
		if err != nil {
			c.ErrorJson(500, "调取go-fastDfs接口失败", nil)
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(result), &resultMap)
		if err != nil {
			logs.Error("RepairStat json exception -> ", err)
			c.ErrorJson(500, "error", nil)
		}
		logs.Info("RepairStat result -> ", resultMap)
		if resultMap["status"] == commons.ApiStatusSuccess {
			count++
		}
	}
	c.SuccessJson("成功修正" + strconv.Itoa(count) + "天数据")
}

// RemoveEmptyDir 删除空目录
func (c *HomeController) RemoveEmptyDir() {
	path, _ := c.GetPeersUrl()
	result, err := httpUtil.PostForm(path+commons.ApiRemoveEmptyDir, nil)
	if err != nil {
		c.ErrorJson(500, "调取go-fastDfs接口失败", nil)
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		logs.Error("RemoveEmptyDir json exception -> ", err)
		c.ErrorJson(500, "error", nil)
	}
	logs.Info("RemoveEmptyDir result -> ", resultMap)

	if resultMap["status"] == commons.ApiStatusSuccess {
		c.SuccessJson("操作成功,正在后台操作,请勿重复使用此功能")
	}
	c.ErrorJson(500, "操作失败,请稍后再试", nil)
}

// Backup 备份元数据, 30天
func (c *HomeController) Backup() {
	count := 0
	for i := 0; i > -30; i-- {
		beforeDate := time.Now().AddDate(0, 0, i).Format("20060102")
		postValue := url.Values{"date": {beforeDate}}
		data := make(map[string]interface{})
		data["date"] = beforeDate
		path, _ := c.GetPeersUrl()
		result, err := httpUtil.PostForm(path+commons.ApiBackup, postValue)
		if err != nil {
			c.ErrorJson(500, "调取go-fastDfs接口失败", nil)
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(result), &resultMap)
		if err != nil {
			logs.Error("RepairStat json exception -> ", err)
			c.ErrorJson(500, "error", nil)
		}
		logs.Info("RepairStat result -> ", resultMap)
		if resultMap["status"] == commons.ApiStatusSuccess {
			count++
		}
	}
	c.SuccessJson("成功备份" + strconv.Itoa(count) + "天数据")
}

// Repair 同步失败修复
func (c *HomeController) Repair() {
	path, _ := c.GetPeersUrl()
	result, err := httpUtil.PostForm(path+commons.ApiRepair+"?force=1", nil)
	if err != nil {
		c.ErrorJson(500, "调取go-fastDfs接口失败", nil)
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		logs.Error("Repair json exception -> ", err)
		c.ErrorJson(500, "error", nil)
	}
	logs.Info("RemoveEmptyDir result -> ", resultMap)

	if resultMap["status"] == commons.ApiStatusSuccess {
		c.SuccessJson("操作成功,正在后台操作,请勿重复使用此功能")
	}
	c.ErrorJson(500, "操作失败,请稍后再试", nil)
}

// GetAllPeers 获取所有集群
func (c *HomeController) GetAllPeers() {
	peers := models.Peers{}
	peersArr, err := peers.GetAllPeers()
	if err != nil {
		logs.Error("获取所有集群出错", err)
		c.ErrorJson(500, "获取数据失败", nil)
	}
	c.SuccessJson(peersArr)
}

// SwitchPeers 切换集群
func (c *HomeController) SwitchPeers() {
	peers, err := c.GetPeers()
	if err != nil {
		c.ErrorJson(500, "系统出错", nil)
	}
	id := c.GetString("id")
	if strconv.Itoa(peers.Id) == id {
		c.ErrorJson(500, "当前正在使用此集群", nil)
	}

	user, _ := c.GetUser()
	peersId, _ := strconv.Atoi(id)
	user.PeersId = peersId
	_ = user.Update()
	c.SuccessJson("更新成功")
}
