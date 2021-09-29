package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/commons"
	"go-fastdfs-web-go/src/form"
	"net/http"
	"net/url"
	"time"
)

var httpUtil = commons.HttpUtil{}

// InstallRouter 安装Router
type InstallRouter struct {
	BaseRouter
	commons.Response
}

// InstallPage 安装页
func (i *InstallRouter) InstallPage(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "install/install.tpl", nil)
}

// CheckLocalServer 检查本机是否安装goFastDfs
func (i *InstallRouter)CheckLocalServer(ctx *gin.Context)  {
	postValue := url.Values{"action": {"get"}}
	result, err := httpUtil.PostForm("http://127.0.0.1:8080/group1" + commons.ApiReload, postValue)
	if err != nil || result == ""{
		result, err = httpUtil.PostForm("http://127.0.0.1:8080" + commons.ApiReload, postValue)
	}

	if err != nil || result == "" {
		i.Error(ctx)
		return
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		i.Error(ctx)
		return
	}

	if resultMap["status"] == commons.ApiStatusSuccess {
		i.SuccessData(ctx, resultMap)
		return
	}
	i.Error(ctx)
}

// CheckServer 校验Server
func (i *InstallRouter) CheckServer(ctx *gin.Context)  {
	var peers form.PeersForm
	if err := ctx.ShouldBind(&peers); err != nil {
		i.ErrorMsg(ctx, fmt.Sprint(err))
		return
	}
	// 拼装url
	path := peers.ServerAddress
	if peers.GroupName != "" {
		path += "/" + peers.GroupName
	}
	path += commons.ApiStatus

	// 测试连接GoFastDfs
	result, err := httpUtil.PostForm(path, nil)
	if err != nil {
		i.ErrorMsg(ctx, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!")
		return
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil {
		i.ErrorMsg(ctx, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!")
		return
	}


	if resultMap["status"] == commons.ApiStatusSuccess {
		i.SuccessData(ctx, resultMap)
		return
	}

	i.ErrorMsg(ctx, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!")
}

// DoInstall 安装
func (i *InstallRouter) DoInstall(ctx *gin.Context) {
	var install = form.InstallForm{}
	if err := ctx.ShouldBind(&install); err != nil {
		i.ErrorMsg(ctx, fmt.Sprint(err))
		return
	}

	peers := install.GetPeers()
	peers.CreateTime = time.Now()
	peers.Save(&peers)
	user := install.GetUser()
	user.CreateTime = time.Now()
	user.PeersId = peers.Id
	user.Save(&user)
	i.Success(ctx)
}