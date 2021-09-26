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

// AddPage 集群添加页
func (c *PeersController) AddPage() {
	c.TplName = "peers/add.tpl"
}

// DoAdd 添加集群
func (c *PeersController) DoAdd() {
	//var peersForm form.PeersForm
	//err := c.ParseForm(&peersForm)
	//if err != nil {
	//	c.ErrorJson(500, "param error", nil)
	//}
	//valid := validation.Validation{}
	//
	//if valid.HasErrors() {
	//	for _, err := range valid.Errors {
	//		logs.Error(err.Key, err.Message)
	//		c.ErrorJson(500, err.Message, nil)
	//	}
	//}
	//
	//peers, err := peersForm.CheckPeers()
	//if err != nil && peers.ServerAddress != "" {
	//	c.ErrorJson(500, "该集群已存在!", nil)
	//}
	//
	//_, err = peersForm.Save()
	//
	//if err != nil{
	//	c.ErrorJson(500, "添加失败!", nil)
	//}
	//c.SuccessJson(peersForm)
}
