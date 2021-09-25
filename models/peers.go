package models

import (
	"github.com/astaxie/beego/orm"
	"go-fastdfs-web-go/commons"
	"strconv"
	"time"
)

// Peers peers table
type Peers struct {
	Id            int
	Name          string    `orm:"size(64);null"`
	GroupName     string    `orm:"size(64);null"`
	ServerAddress string    `orm:"size(64);null"`
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
	ShowAddress   string    `orm:"size(64);null"`
}

func init() {
	orm.RegisterModel(new(Peers))
}

// Save 保存
func (peers *Peers) Save() (int64, error) {
	return orm.NewOrm().Insert(peers)
}

// GetById 根据id获取Peers
func (peers *Peers) GetById() (Peers, error) {
	err := orm.NewOrm().Read(peers, "Id")
	return *peers, err
}

// GetAllPeers 获取所有集群
func (peers *Peers) GetAllPeers() ([]Peers, error) {
	var peersArr []Peers
	_, err := orm.NewOrm().QueryTable("peers").All(&peersArr)
	return peersArr, err
}

// PageList 分页
func (peers *Peers) PageList(page string, limit string) commons.Pager {
	o := orm.NewOrm()
	qs := o.QueryTable("peers")
	pageNum, _ := strconv.Atoi(page)
	limitNum, _ := strconv.Atoi(limit)
	// 总条数
	count, _ := qs.Count()
	//存储分页数据的切片
	peersArr := new([]Peers)
	//获取分页数据
	_, _ = qs.Limit(limitNum, limitNum*(pageNum-1)).All(peersArr)

	pager := commons.Pager{}
	pager.Msg = "success"
	pager.Data = peersArr
	pager.Total = count
	pager.State = 200
	return pager
}
