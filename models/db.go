package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 用户表
type User struct {
	Id              int64
	Account         string    `orm:"size(64)"`
	Password        string    `orm:"size(64)"`
	Name            string    `orm:"size(64)"`
	CredentialsSalt string    `orm:"size(64)"`
	Email           string    `orm:"size(64)"`
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime      time.Time `orm:"null"`
	Peers           *Peers    `orm:"rel(one)"`
}

// 集群表
type Peers struct {
	Id            int64
	ServerName    string    `orm:"size(255)"`
	GroupName     string    `orm:"size(255)"`
	ServerAddress string    `orm:"size(255)"`
	ShowAddress   string    `orm:"null;size(255)"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User), new(Peers))
}
