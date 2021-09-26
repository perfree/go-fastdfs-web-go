package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"go-fastdfs-web-go/models"
)

type BaseController struct {
	beego.Controller
}

// ReturnMsg 定义统一结果集
type ReturnMsg struct {
	Code int
	Msg  string
	Data interface{}
}

// SuccessJson 成功
func (b *BaseController) SuccessJson(data interface{}) {

	res := ReturnMsg{
		200, "success", data,
	}
	b.Data["json"] = res
	b.ServeJSON()
	b.StopRun()
}

// ErrorJson 失败
func (b *BaseController) ErrorJson(code int, msg string, data interface{}) {

	res := ReturnMsg{
		code, msg, data,
	}

	b.Data["json"] = res
	b.ServeJSON()
	b.StopRun()
}

// GetPeersUrl 获取PeersUrl
func (b *BaseController) GetPeersUrl() (string, error) {
	userId := b.GetSession("userId").(int)
	user := models.User{}
	user.Id = userId
	user, err := user.GetById()

	peers := models.Peers{}
	if err != nil {
		return "", err
	}
	peers.Id = user.PeersId
	peers, err = peers.GetById()
	if err != nil {
		return "", err
	}

	if peers.GroupName != "" {
		return peers.ServerAddress + "/" + peers.GroupName, err
	}
	return peers.ServerAddress, err
}

// GetPeers 获取Peers
func (b *BaseController) GetPeers() (models.Peers, error) {
	userId := b.GetSession("userId").(int)
	user := models.User{}
	user.Id = userId
	user, err := user.GetById()

	peers := models.Peers{}
	if err != nil {
		return peers, err
	}
	peers.Id = user.PeersId
	peers, err = peers.GetById()
	if err != nil {
		return peers, err
	}
	return peers, err
}

// GetUser 获取当前用户
func (b *BaseController) GetUser() (models.User, error) {
	userId := b.GetSession("userId").(int)
	user := models.User{}
	user.Id = userId
	return user.GetById()
}

// ValidParam 校验参数
func (b *BaseController) ValidParam(obj interface{}, errMsg string) {
	valid := validation.Validation{}
	e, err := valid.Valid(obj)
	if err != nil {
		logs.Error("DoAdd -> ", err)
		b.ErrorJson(500, errMsg, nil)
	}
	if !e {
		for _, err := range valid.Errors {
			logs.Error("DoAdd -> ", err.Key, err.Message)
			b.ErrorJson(500, err.Message, nil)
		}
	}
}
