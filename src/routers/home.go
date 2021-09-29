package routers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/commons"
	"go-fastdfs-web-go/src/models"
	"go-fastdfs-web-go/src/setting"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"time"
)

type HomeRouter struct {
	BaseRouter
	commons.Response
}

// Home 首页
func (h *HomeRouter)Home(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "home/home.tpl", gin.H{
		"osName" : runtime.GOOS,
		"osArch" : runtime.GOARCH,
		"version" : setting.AppSetting.AppVer,
		"versionDate" : setting.AppSetting.AppVerDate,
	})
}

// GetStatus 获取状态信息
func (h *HomeRouter) GetStatus(ctx *gin.Context)  {
	path, err := h.GetPeersUrl(ctx)
	if err != nil {
		h.ErrorMsg(ctx, "系统异常")
		return
	}

	result, err := httpUtil.PostForm(path+commons.ApiStatus, nil)
	if err != nil {
		h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
		return
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
		return
	}

	if resultMap["status"] == commons.ApiStatusSuccess {
		data := resultMap["data"].(map[string]interface{})
		result, err := commons.GetStatus(data)
		if err != nil {
			h.ErrorMsg(ctx, "数据解析错误")
			return
		}
		h.SuccessData(ctx, result)
		return
	}
	h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
}

// RepairStat 修正统计信息(30天)
func (h *HomeRouter) RepairStat(ctx *gin.Context) {
	count := 0
	for i := 0; i > -30; i-- {
		beforeDate := time.Now().AddDate(0, 0, i).Format("20060102")
		postValue := url.Values{"date": {beforeDate}}
		data := make(map[string]interface{})
		data["date"] = beforeDate
		path, _ := h.GetPeersUrl(ctx)
		result, err := httpUtil.PostForm(path+commons.ApiRepairStat, postValue)
		if err != nil {
			h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
			return
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(result), &resultMap)
		if err != nil {
			h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
			return
		}
		if resultMap["status"] == commons.ApiStatusSuccess {
			count++
		}
	}
	h.SuccessMsg(ctx, "成功修正" + strconv.Itoa(count) + "天数据")
}


// RemoveEmptyDir 删除空目录
func (h *HomeRouter) RemoveEmptyDir(ctx *gin.Context) {
	path, _ := h.GetPeersUrl(ctx)
	result, err := httpUtil.PostForm(path+commons.ApiRemoveEmptyDir, nil)
	if err != nil {
		h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
		return
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
		return
	}

	if resultMap["status"] == commons.ApiStatusSuccess {
		h.SuccessMsg(ctx, "操作成功,正在后台操作,请勿重复使用此功能")
		return
	}
	h.ErrorMsg(ctx, "操作失败,请稍后再试")
}

// Backup 备份元数据, 30天
func (h *HomeRouter) Backup(ctx *gin.Context) {
	count := 0
	for i := 0; i > -30; i-- {
		beforeDate := time.Now().AddDate(0, 0, i).Format("20060102")
		postValue := url.Values{"date": {beforeDate}}
		data := make(map[string]interface{})
		data["date"] = beforeDate
		path, _ := h.GetPeersUrl(ctx)
		result, err := httpUtil.PostForm(path+commons.ApiBackup, postValue)
		if err != nil {
			h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
			return
		}
		var resultMap map[string]interface{}
		err = json.Unmarshal([]byte(result), &resultMap)
		if err != nil {
			h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
			return
		}
		if resultMap["status"] == commons.ApiStatusSuccess {
			count++
		}
	}
	h.SuccessMsg(ctx, "成功备份" + strconv.Itoa(count) + "天数据")
}

// Repair 同步失败修复
func (h *HomeRouter) Repair(ctx *gin.Context) {
	path, _ := h.GetPeersUrl(ctx)
	result, err := httpUtil.PostForm(path+commons.ApiRepair+"?force=1", nil)
	if err != nil {
		h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
		return
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		h.ErrorMsg(ctx, "调取go-fastDfs接口失败")
		return
	}

	if resultMap["status"] == commons.ApiStatusSuccess {
		h.SuccessMsg(ctx, "操作成功,正在后台操作,请勿重复使用此功能")
		return
	}
	h.ErrorMsg(ctx, "操作失败,请稍后再试")
}

// GetAllPeers 获取所有集群
func (h *HomeRouter) GetAllPeers(ctx *gin.Context) {
	var peers models.Peers
	peersArr, err := peers.GetAllPeers()
	if err != nil {
		h.ErrorMsg(ctx, "获取数据失败")
		return
	}
	h.SuccessData(ctx, peersArr)
}

// SwitchPeers 切换集群
func (h *HomeRouter) SwitchPeers(ctx *gin.Context) {
	peers, err := h.GetPeers(ctx)
	if err != nil {
		h.ErrorMsg(ctx, "系统出错")
		return
	}
	id := ctx.PostForm("id")
	if strconv.Itoa(peers.Id) == id {
		h.ErrorMsg(ctx, "当前正在使用此集群")
		return
	}

	user, _ := h.GetUser(ctx)
	peersId, _ := strconv.Atoi(id)
	user.PeersId = peersId
	user.Update(user)
	h.SuccessMsg(ctx,"更新成功")
}