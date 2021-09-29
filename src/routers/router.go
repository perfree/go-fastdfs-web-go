package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/filters"
	"go-fastdfs-web-go/src/setting"
)

type Routers struct {
	SystemRouter SystemRouter
	InstallRouter InstallRouter
	HomeRouter HomeRouter
	PeersRouter PeersRouter
	FileRouter FileRouter
	SettingRouter SettingRouter

}

// InitRouter 初始化路由
func (routers *Routers) InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	store := cookie.NewStore([]byte(setting.AppSetting.SessionSecret))
	store.Options(sessions.Options{
		MaxAge: 60 * 60,
		Path:   "/",
	})
	r.Use(sessions.Sessions("session", store))

	r.LoadHTMLGlob("./template/**/*.tpl")
	r.Static("/static", "./static")


	r.Use(filters.InstallFilter()) // 安装拦截
	r.GET("/install", routers.InstallRouter.InstallPage)
	r.GET("/install/checkLocalServer", routers.InstallRouter.CheckLocalServer)
	r.POST("/install/checkServer", routers.InstallRouter.CheckServer)
	r.POST("/install/doInstall", routers.InstallRouter.DoInstall)
	r.GET("/login", routers.SystemRouter.LoginPage)
	r.POST("/doLogin", routers.SystemRouter.DoLogin)
	r.GET("/logout", routers.SystemRouter.LogOut)

	r.Use(filters.LoginFilter()) // 登录拦截
	r.GET("/", routers.SystemRouter.IndexPage)

	r.GET("/home", routers.HomeRouter.Home)
	r.POST("/home/getStatus", routers.HomeRouter.GetStatus)
	r.POST("/home/repair_stat", routers.HomeRouter.RepairStat)
	r.POST("/home/remove_empty_dir", routers.HomeRouter.RemoveEmptyDir)
	r.POST("/home/backup", routers.HomeRouter.Backup)
	r.POST("/home/repair", routers.HomeRouter.Repair)
	r.POST("/home/getAllPeers", routers.HomeRouter.GetAllPeers)
	r.POST("/home/switchPeers", routers.HomeRouter.SwitchPeers)


	r.GET("/peers", routers.PeersRouter.Index)
	r.GET("/peers/page", routers.PeersRouter.PageList)
	r.GET("/peers/add", routers.PeersRouter.AddPage)
	r.POST("/peers/doAdd", routers.PeersRouter.DoAdd)
	r.GET("/peers/edit", routers.PeersRouter.EditPage)
	r.POST("/peers/doEdit", routers.PeersRouter.DoEdit)
	r.POST("/peers/del", routers.PeersRouter.Del)

	r.GET("/file", routers.FileRouter.Index)
	r.POST("/file/getDirFile", routers.FileRouter.GetDirFile)
	r.POST("/file/deleteFile", routers.FileRouter.DeleteFile)
	r.POST("/file/details", routers.FileRouter.Details)
	r.POST("/file/downloadFile", routers.FileRouter.DownloadFile)

	r.GET("/file/upload", routers.FileRouter.UploadPage)
	r.POST("/file/upload/fileUpload", routers.FileRouter.Upload)

	r.GET("/settings/user", routers.SettingRouter.UserPage)
	r.POST("/settings/editUser", routers.SettingRouter.EditUser)

	return r
}