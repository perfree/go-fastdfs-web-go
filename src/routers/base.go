package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/models"
)

type BaseRouter struct {
}

// GetPeersUrl 获取PeersUrl
func (b *BaseRouter) GetPeersUrl(ctx *gin.Context) (string, error) {
	user, err := b.GetUser(ctx)
	if err != nil {
		return "", err
	}

	var peers models.Peers
	peers, err = peers.GetById(user.PeersId)
	if err != nil {
		return "", err
	}

	if peers.GroupName != "" {
		return peers.ServerAddress + "/" + peers.GroupName, err
	}
	return peers.ServerAddress, err
}

// GetPeers 获取Peers
func (b *BaseRouter) GetPeers(ctx *gin.Context) (models.Peers, error) {
	user, err := b.GetUser(ctx)
	if err != nil {
		return models.Peers{}, err
	}

	var peers models.Peers
	peers, err = peers.GetById(user.PeersId)
	if err != nil {
		return peers, err
	}
	return peers, err
}

// GetUser 获取当前用户
func (b *BaseRouter) GetUser(ctx *gin.Context) (models.User, error) {
	session := sessions.Default(ctx)
	userId := session.Get("UserId").(int)
	var user models.User
	return user.GetById(userId)
}

// GetShowUrl 获取ShowUrl
func (b *BaseRouter) GetShowUrl(ctx *gin.Context) string {
	peers, _ := b.GetPeers(ctx)
	showUrl := ""
	if peers.ShowAddress == "" {
		if peers.GroupName == "" {
			showUrl = peers.ServerAddress
		} else {
			showUrl = peers.ServerAddress + "/" + peers.GroupName
		}
	} else {
		if peers.GroupName == "" {
			showUrl = peers.ShowAddress
		} else {
			showUrl = peers.ShowAddress + "/" + peers.GroupName
		}
	}
	return showUrl
}

// GetShowUrlNotGroup 获取ShowUrl
func (b *BaseRouter) GetShowUrlNotGroup(ctx *gin.Context) string {
	peers, _ := b.GetPeers(ctx)
	showUrl := ""
	if peers.ShowAddress == "" {
		showUrl = peers.ServerAddress
	} else {
		showUrl = peers.ShowAddress
	}
	return showUrl
}