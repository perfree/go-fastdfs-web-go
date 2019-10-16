package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id    int // 主键
	Name  string
	Age   int
	Sex   string
}

func init() {
	orm.RegisterModel(new(User))
}