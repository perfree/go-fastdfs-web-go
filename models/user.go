package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 获取所有用户信息
func GetUsers() (messages []User) {
	o := orm.NewOrm()
	_,err := o.Raw("SELECT * FROM `user`").QueryRows(&messages)
	if err != nil{
		logs.Error(err.Error())
	}
	return messages
}