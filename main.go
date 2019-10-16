package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "go-fastdfs-web-go/routers"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	_ = orm.RegisterDriver("sqlite3", orm.DRSqlite)
	_ = orm.RegisterDataBase("default", "sqlite3", beego.AppConfig.String("data_sources"))
	_ = orm.RunSyncdb("default", false, false)
}

// log配置
func initLogger()(err error) {
	config := make(map[string]interface{})
	config["filename"] = beego.AppConfig.String("log_path")
	// map 转 json
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("initLogger failed, marshal err:", err)
		return
	}
	_ = logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.SetLogFuncCall(true)
	logs.SetLevel(5)
	return
}

func main() {
	_ = initLogger()
	o := orm.NewOrm()
	_ = o.Using("default")
	beego.Run()
}

