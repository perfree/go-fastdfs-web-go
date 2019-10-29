package models

import "github.com/astaxie/beego/orm"

// 用户表
type User struct {
	Id              int
	Account         string
	Password        string
	Name            string
	CredentialsSalt string
	Email           string
	CreateTime      string
	UpdateTime      string
	Peers           *Peers `orm:"rel(one)"`
}

// 集群表
type Peers struct {
	Id            int
	Name          string
	GroupName     string
	ServerAddress string
	ShowAddress   string
	CreateTime    string
}

func init() {
	orm.RegisterModel(new(User), new(Peers))
}
