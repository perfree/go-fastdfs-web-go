package models

import (
	"go-fastdfs-web-go/src/commons"
	"strconv"
	"time"
)

// Peers peers table
type Peers struct {
	Id            int       `gorm:"not null;primary key;auto_increment"`
	Name          string    `gorm:"not null;type:varchar(64)"`
	GroupName     string    `gorm:"not null;type:varchar(64)"`
	ServerAddress string    `gorm:"not null;type:varchar(256)"`
	CreateTime    time.Time `gorm:"not null;default:null"`
	ShowAddress   string    `gorm:"type:varchar(256)"`
}

// Save 保存
func (p *Peers) Save(peers *Peers)  {
	db.Create(peers)
}

// GetById 根据id查询Peers
func (p *Peers) GetById(id int) (Peers, error) {
	var peers Peers
	err := db.Model(&Peers{}).Where(&Peers{Id: id}).First(&peers).Error
	return peers, err
}

// GetAllPeers 获取所有Peers
func (p *Peers) GetAllPeers() ([]Peers, error) {
	var peers []Peers
	err := db.Model(&Peers{}).Find(&peers).Error
	return peers, err
}

// PageList 分页
func (p *Peers) PageList(page string, limit string) commons.Pager {
	peersDb := db.Model(&Peers{})
	var count int64
	peersDb.Count(&count)

	pageIndex,_ := strconv.Atoi(page)
	pageSize,_  :=  strconv.Atoi(limit)

	var peersList []Peers
	peersDb.Offset((pageIndex-1)*pageSize).Limit(pageSize).Find(&peersList)
	pager := commons.Pager{}
	pager.Msg = "success"
	pager.Data = peersList
	pager.Total = count
	pager.State = 200
	return pager
}

// CheckPeers 检查address是否存在
func (p *Peers) CheckPeers(address string) (Peers, error) {
	var peers Peers
	err := db.Model(&Peers{}).Where(&Peers{ServerAddress: address}).First(&peers).Error
	return peers, err
}

// Update 更新
func (p *Peers) Update(peers Peers) {
	db.Model(&Peers{}).Where(&Peers{Id: peers.Id}).Updates(&peers)
}

// Del 删除
func (p *Peers) Del(id int) {
	db.Delete(&Peers{Id: id})
}