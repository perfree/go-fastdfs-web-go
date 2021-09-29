package models

import (
	"time"
)

// User user table
type User struct {
	Id              int       `gorm:"not null;primary key;auto_increment"`
	Account         string    `gorm:"not null;type:varchar(64)"`
	Password        string    `gorm:"not null;type:varchar(64)"`
	Name            string    `gorm:"not null;type:varchar(64)"`
	CredentialsSalt string    `gorm:"not null;type:varchar(64)"`
	Email           string    `gorm:"not null;type:varchar(64)"`
	PeersId         int       `gorm:"not null;type:varchar(64);default:null"`
	CreateTime      time.Time `gorm:"not null;default:null"`
	UpdateTime      time.Time `gorm:"default:null"`
}

// UserCount 获取用户数量
func (u *User) UserCount() int64 {
	var count int64
	db.Model(&User{}).Count(&count)
	return count
}

// Save 保存
func (u *User) Save(user *User) {
	db.Create(user)
}

// GetByAccount 根据账户查询User
func (u *User) GetByAccount(account string) (User, error) {
	var user User
	err := db.Model(&User{}).Where(&User{Account: account}).First(&user).Error
	return user, err
}

// GetById 根据id查询User
func (u *User) GetById(id int) (User, error) {
	var user User
	err := db.Model(&User{}).Where(&User{Id: id}).First(&user).Error
	return user, err
}

// Update 更新
func (u *User) Update(user User) {
	db.Model(&User{}).Where(&User{Id: user.Id}).Updates(&user)
}