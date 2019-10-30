package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 获取当前用户数
func GetUsesTotal() (totalCount int64) {
	o := orm.NewOrm()
	totalCount, _ = o.QueryTable("user").Count()
	return totalCount
}

// 获取所有用户信息
func GetUsers() (messages []User) {
	o := orm.NewOrm()
	/*	TotalCount,_:=o.QueryTable("user").Count()
		fmt.Println(TotalCount)*/
	_, err := o.QueryTable("user").RelatedSel().All(&messages)

	//_,err := o.Raw("SELECT * FROM `t_user` t1 left join `t_peers` t2 on t1.peersId = t2.id").QueryRows(&messages)
	if err != nil {
		logs.Error(err.Error())
	}
	return messages
}
