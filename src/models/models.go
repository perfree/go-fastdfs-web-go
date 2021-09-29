package models

import (
	"go-fastdfs-web-go/src/setting"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDataBase()  {
	var err error
	db, err = gorm.Open(sqlite.Open(setting.AppSetting.SqlFile), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database")
	}
	dbPool, _ := db.DB()
	dbPool.SetMaxIdleConns(10)
	dbPool.SetMaxOpenConns(100)
	_ = db.AutoMigrate(&User{}, &Peers{})
}