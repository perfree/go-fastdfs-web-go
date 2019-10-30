package filter

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"go-fastdfs-web-go/models"
)

// 检查是否已安装
func CheckInstall(ctx *context.Context) {
	if models.GetUsesTotal() <= 0 && ctx.Request.RequestURI != "/install" {
		ctx.Redirect(301, "/install")
	}
}

// 检查是否已登录
func CheckLogin(ctx *context.Context) {
	fmt.Println("123")
}
