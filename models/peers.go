package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 保存集群信息
func SavePeer(peers Peers) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&peers)
	if err != nil {
		fmt.Println(err.Error())
	}
	return id
}
