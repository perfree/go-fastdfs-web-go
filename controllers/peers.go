package controllers

import (
	"go-fastdfs-web-go/models"
)

type PeersController struct {
	BaseController
}

// Index 集群管理首页
func (c *PeersController) Index() {
	c.TplName = "peers/list.tpl"
}

// PageList 获取集群分页列表
func (c *PeersController) PageList() {
	page := c.Ctx.Input.Query("page")
	limit := c.Ctx.Input.Query("limit")
	peers := models.Peers{}
	pager := peers.PageList(page, limit)
	c.Data["json"] = pager
	c.ServeJSON()
	c.StopRun()
}
