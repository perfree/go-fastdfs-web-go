package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"go-fastdfs-web-go/models"
	"strings"
)

// LoginFilter 登录拦截
func LoginFilter(ctx *context.Context) {
	_, ok := ctx.Input.Session("userId").(int)
	if !ok && ctx.Request.RequestURI != "/login" && ctx.Request.RequestURI != "/doLogin" && !strings.Contains(ctx.Request.RequestURI, "/install") {
		ctx.Redirect(302, "/login")
	}
}

// InstallFilter 安装拦截
func InstallFilter(ctx *context.Context) {
	user := models.User{}
	s := user.QueryUserCount()
	if s <= 0 && !strings.Contains(ctx.Request.RequestURI, "/install") {
		logs.Info("install filter -> 未查询到用户,判定执行安装操作")
		ctx.Redirect(302, "/install")
	}
	if s > 0 && strings.Contains(ctx.Request.RequestURI, "/install") {
		ctx.Redirect(302, "/")
	}
}
