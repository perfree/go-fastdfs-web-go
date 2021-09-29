package setting

import (
	"github.com/go-ini/ini"
	"log"
)

// App 定义App配置文件映射
type App struct {
	RunMode 	 string
	HttpPort     int
	SqlFile 	 string
	LogFile 	 string
	SessionSecret string
	AppVer string
	AppVerDate string
}

var AppSetting = &App{}

var cfg *ini.File

// LoadSetting 加载配置
func LoadSetting() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("LoadSetting, fail to parse 'conf/app.ini': %v", err)
	}
	convert("app", AppSetting)
}

// convert 配置文件转换
func convert(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.convert %s err: %v", section, err)
	}
}
