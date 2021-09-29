package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/commons"
	"go-fastdfs-web-go/src/form"
	"go-fastdfs-web-go/src/models"
	"net/http"
	"strconv"
	"time"
)

type PeersRouter struct {
	BaseRouter
	commons.Response
}

// Index 集群管理首页
func (p *PeersRouter) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "peers/list.tpl" , nil)
}

// PageList 获取集群分页列表
func (p *PeersRouter) PageList(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	var peers models.Peers
	pager := peers.PageList(page, limit)
	ctx.JSON(http.StatusOK, pager)
}

// AddPage 集群添加页
func (p *PeersRouter) AddPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "peers/add.tpl" , nil)
}

// DoAdd 添加集群
func (p *PeersRouter) DoAdd(ctx *gin.Context) {
	var peersForm form.PeersForm
	if err := ctx.ShouldBind(&peersForm); err != nil {
		p.ErrorMsg(ctx, fmt.Sprint(err))
		return
	}

	var peers models.Peers
	_, err := peers.CheckPeers(peersForm.ServerAddress)
	if err == nil {
		p.ErrorMsg(ctx, "该集群已存在")
		return
	}

	// 拼装url
	path := peersForm.ServerAddress
	if peersForm.GroupName != "" {
		path += "/" + peersForm.GroupName
	}
	path += commons.ApiStatus

	// 测试连接GoFastDfs
	result, err := httpUtil.PostForm(path, nil)
	if err != nil {
		p.ErrorMsg(ctx, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!")
		return
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil || resultMap["status"] != commons.ApiStatusSuccess {
		p.ErrorMsg(ctx, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!")
		return
	}

	peers = peersForm.GetPeers()
	peers.CreateTime = time.Now()
	peers.Save(&peers)

	p.SuccessData(ctx, peers)
}

// EditPage 编辑页面
func (p *PeersRouter) EditPage(ctx *gin.Context) {
	id := ctx.Query("id")
	peers := models.Peers{}
	peers.Id, _ = strconv.Atoi(id)
	peers, _ = peers.GetById(peers.Id)
	ctx.HTML(http.StatusOK, "peers/edit.tpl", gin.H{
		"peers": peers,
	})
}

// DoEdit 编辑
func (p *PeersRouter) DoEdit(ctx *gin.Context) {
	var peersForm form.PeersForm
	if err := ctx.ShouldBind(&peersForm); err != nil {
		p.ErrorMsg(ctx, fmt.Sprint(err))
		return
	}

	var peers models.Peers

	_, err := peers.CheckPeers(peersForm.ServerAddress)
	oldPeers, _ := peers.GetById(peersForm.Id)
	if oldPeers.ServerAddress != peersForm.ServerAddress && err == nil {
		p.ErrorMsg(ctx, "该集群已存在!")
		return
	}

	// 拼装url
	path := peersForm.ServerAddress
	if peersForm.GroupName != "" {
		path += "/" + peersForm.GroupName
	}
	path += commons.ApiStatus

	// 测试连接GoFastDfs
	result, err := httpUtil.PostForm(path, nil)
	if err != nil {
		p.ErrorMsg(ctx, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!")
		return
	}

	var resultMap map[string]interface{}
	err = json.Unmarshal([]byte(result), &resultMap)
	if err != nil || resultMap["status"] != commons.ApiStatusSuccess {
		p.ErrorMsg(ctx, "连接GoFastDfs服务失败!请检查服务地址是否正确,以及是否配置白名单!")
		return
	}

	peers.Update(peersForm.GetPeers())
	p.SuccessData(ctx, peersForm.GetPeers())
}

// Del 删除
func (p *PeersRouter) Del(ctx *gin.Context) {
	idStr := ctx.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	var peers models.Peers
	peers.Del(id)
	p.Success(ctx)
}