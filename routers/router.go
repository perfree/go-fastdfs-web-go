package routers

import (
	"github.com/astaxie/beego"
	"go-fastdfs-web-go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
