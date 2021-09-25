package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"go-fastdfs-web-go/filters"
	_ "go-fastdfs-web-go/routers"
)

func init() {
	_ = orm.RegisterDriver("sqlite", orm.DRSqlite)
	_ = orm.RegisterDataBase("default", "sqlite3", "DataBase.db")
	_ = orm.RunSyncdb("default", false, true)
}

func main() {
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"logs/go-fastDfs-web.log","level":6}`)
	logs.EnableFuncCallDepth(true)
	beego.InsertFilter("/*", beego.BeforeRouter, filters.InstallFilter)
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LoginFilter)
	beego.Run()
}
