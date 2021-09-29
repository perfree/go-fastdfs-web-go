package form

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
	"go-fastdfs-web-go/src/models"
)

type InstallForm struct {
	Name          string    `binding:"required,max=50"`
	GroupName     string    `binding:"required,max=50"`
	ServerAddress string    `binding:"required,max=100,url"`
	ShowAddress   string    `binding:"omitempty,max=100,url"`
	Account       string	`binding:"required,max=30"`
	Password      string	`binding:"required,max=30"`
	UserName      string	`binding:"required,max=30"`
	Email         string	`binding:"required,max=30,email"`
}


// GetUser 获取user
func (install *InstallForm) GetUser() models.User {
	user := models.User{}
	user.Email = install.Email
	user.Account = install.Account
	user.Name = install.UserName
	user.Password = install.Password
	user.CredentialsSalt = uuid.NewV4().String()

	m5 := md5.New()
	m5.Write([]byte(user.Password))
	m5.Write([]byte(user.CredentialsSalt))
	st := m5.Sum(nil)
	user.Password = hex.EncodeToString(st)
	return user
}

// GetPeers 获取Peers
func (install *InstallForm) GetPeers() models.Peers {
	peers := models.Peers{}
	peers.GroupName = install.GroupName
	peers.ServerAddress = install.ServerAddress
	peers.ShowAddress = install.ShowAddress
	peers.Name = install.Name
	return peers
}