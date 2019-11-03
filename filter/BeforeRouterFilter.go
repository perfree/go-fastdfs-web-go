package filter

import (
	"github.com/astaxie/beego/context"
	"go-fastdfs-web-go/models"
	"strings"
)

// 检查是否已安装
func CheckInstall(ctx *context.Context) {
	if models.GetUsesTotal() <= 0 && strings.Index(ctx.Request.RequestURI, "/install") < 0 {
		ctx.Redirect(301, "/install")
	}
}

// 检查是否已登录
func CheckLogin(ctx *context.Context) {
	if strings.Index(ctx.Request.RequestURI, "/install") < 0 && strings.Index(ctx.Request.RequestURI, "/login") < 0 {
		ctx.Redirect(301, "/login")
	}
}
