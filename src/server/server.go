package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/routers"
	"go-fastdfs-web-go/src/setting"
	"io"
	"os"
)


// Run 运行服务
func Run()  {
	gin.SetMode(setting.AppSetting.RunMode)
	gin.DisableConsoleColor()
	f, _ := os.Create(setting.AppSetting.LogFile)
	gin.DefaultWriter = io.MultiWriter(f)
	routersInit := new(routers.Routers).InitRouter()

	fmt.Println(setting.AppSetting.HttpPort)
	_ = routersInit.Run(fmt.Sprintf(":%d", setting.AppSetting.HttpPort))

}