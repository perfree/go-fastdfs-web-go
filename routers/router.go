package routers

import (
	"github.com/astaxie/beego"
	"go-fastdfs-web-go/controllers"
)

func init() {
	beego.Router("/", &controllers.SystemController{}, "get:Home")
	beego.Router("/login", &controllers.SystemController{}, "get:Login")
	beego.Router("/doLogin", &controllers.SystemController{}, "post:DoLogin")
	beego.Router("/logout", &controllers.SystemController{}, "get:LogOut")

	beego.Router("/install", &controllers.InstallController{})
	beego.Router("/install/checkLocalServer", &controllers.InstallController{}, "get:CheckLocalServer")
	beego.Router("/install/checkServer", &controllers.InstallController{}, "post:CheckServer")
	beego.Router("/install/doInstall", &controllers.InstallController{}, "post:DoInstall")

	beego.Router("/home", &controllers.HomeController{}, "get:Home")
	beego.Router("/home/getStatus", &controllers.HomeController{}, "post:GetStatus")
	beego.Router("/home/repair_stat", &controllers.HomeController{}, "post:RepairStat")
	beego.Router("/home/remove_empty_dir", &controllers.HomeController{}, "post:RemoveEmptyDir")
	beego.Router("/home/backup", &controllers.HomeController{}, "post:Backup")
	beego.Router("/home/repair", &controllers.HomeController{}, "post:Repair")
	beego.Router("/home/getAllPeers", &controllers.HomeController{}, "post:GetAllPeers")
	beego.Router("/home/switchPeers", &controllers.HomeController{}, "post:SwitchPeers")

	beego.Router("/peers", &controllers.PeersController{}, "get:Index")
	beego.Router("/peers/page", &controllers.PeersController{}, "get:PageList")
	beego.Router("/peers/add", &controllers.PeersController{}, "get:AddPage")
	beego.Router("/peers/doAdd", &controllers.PeersController{}, "post:DoAdd")
	beego.Router("/peers/edit", &controllers.PeersController{}, "get:EditPage")
	beego.Router("/peers/doEdit", &controllers.PeersController{}, "post:DoEdit")
	beego.Router("/peers/del", &controllers.PeersController{}, "post:Del")

}
