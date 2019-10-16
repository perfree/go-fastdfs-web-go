package routers

import (
	"github.com/astaxie/beego"
	"go-fastdfs-web-go/controllers"
)

// 定义路由
func init() {
	beego.Router("/",&controllers.HelloController{})
	beego.Router("/getUsers",&controllers.HelloController{},"get:GetUsers")
}
