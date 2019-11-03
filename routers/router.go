package routers

import (
	"github.com/astaxie/beego"
	"go-fastdfs-web-go/controllers"
)

// 定义路由
func init() {
	// 安装
	beego.Router("/install", &controllers.InstallController{})
	beego.Router("/install/doInstall", &controllers.InstallController{}, "post:DoInstall")
	beego.Router("/install/checkServer", &controllers.InstallController{}, "post:CheckServer")

	// 登录
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/getUsers", &controllers.HelloController{}, "get:GetUsers")
}
