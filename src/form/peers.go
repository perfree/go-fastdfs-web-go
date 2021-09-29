package form

import (
	"go-fastdfs-web-go/src/models"
	"time"
)

// PeersForm peers form
type PeersForm struct {
	Id            int
	Name          string    `binding:"required,max=50"`
	GroupName     string    `binding:"omitempty,max=50"`
	ServerAddress string    `binding:"required,max=100,url"`
	CreateTime    time.Time
	ShowAddress   string    `binding:"omitempty,max=100,url"`
}

// GetPeers 获取Peers
func (p *PeersForm)GetPeers() models.Peers {
	peers := models.Peers{
		Id : p.Id,
		Name: p.Name,
		GroupName: p.GroupName,
		ServerAddress: p.ServerAddress,
		CreateTime: p.CreateTime,
		ShowAddress: p.ShowAddress,
	}
	return peers
}