package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// User user table
type User struct {
	Id              int       `orm:"AUTO_INCREMENT"`
	Account         string    `orm:"size(64)"`
	Password        string    `orm:"size(64)"`
	Name            string    `orm:"size(64)"`
	CredentialsSalt string    `orm:"size(64)"`
	Email           string    `orm:"size(64);null"`
	PeersId         int       `orm:"null"`
	CreateTime      time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime      time.Time `orm:"type(datetime);null"`
}

func init() {
	orm.RegisterModel(new(User))
}

// QueryUserCount 查询用户表数据量
func (u *User) QueryUserCount() int64 {
	n, _ := orm.NewOrm().QueryTable("user").Count()
	return n
}

// Save 保存用户
func (u *User) Save() (int64, error) {
	return orm.NewOrm().Insert(u)
}

// GetByAccount 根据账户获取user
func (u *User) GetByAccount() (User, error) {
	err := orm.NewOrm().Read(u, "Account")
	return *u, err
}

// GetById 根据id获取user
func (u *User) GetById() (User, error) {
	err := orm.NewOrm().Read(u, "Id")
	return *u, err
}

// Update 更新用户
func (u *User) Update() error {
	_, err := orm.NewOrm().Update(u, "PeersId")
	return err
}
