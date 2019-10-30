package routers

import (
	"github.com/astaxie/beego"
	"go-fastdfs-web-go/controllers"
)

// 定义路由
func init() {
	beego.Router("/install", &controllers.InstallController{})
	beego.Router("/install/doInstall", &controllers.InstallController{}, "post:DoInstall")
	beego.Router("/install/checkServer", &controllers.InstallController{}, "post:CheckServer")
	beego.Router("/getUsers", &controllers.HelloController{}, "get:GetUsers")
}
