package main

import (
	"go-fastdfs-web-go/src/models"
	"go-fastdfs-web-go/src/server"
	"go-fastdfs-web-go/src/setting"
)

func init()  {
	setting.LoadSetting()
	models.InitDataBase()
}

func main() {
	server.Run()
}