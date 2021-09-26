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
func (u *User) Save(user User) (int64, error) {
	return orm.NewOrm().Insert(&user)
}

// GetByAccount 根据账户获取user
func (u *User) GetByAccount(account string) (User, error) {
	var user User
	err := orm.NewOrm().Raw("select * from user where account = ?", account).QueryRow(&user)
	return user, err
}

// GetById 根据id获取user
func (u *User) GetById(id int) (User, error) {
	var user User
	err := orm.NewOrm().Raw("select * from user where id = ?", id).QueryRow(&user)
	return user, err
}

// Update 更新用户
func (u *User) Update(user User) error {
	_, err := orm.NewOrm().Update(&user, "PeersId")
	return err
}
