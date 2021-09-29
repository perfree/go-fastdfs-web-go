package filters

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/models"
	"net/http"
	"strings"
)

// InstallFilter 安装拦截
func InstallFilter() gin.HandlerFunc {
	return func(c *gin.Context){
		user := models.User{}
		count := user.UserCount()
		if count <= 0 && !strings.Contains(c.FullPath(), "/install") {
			c.Redirect(http.StatusFound, "/install")
		}
		if count > 0 && strings.Contains(c.FullPath(), "/install") {
			c.Redirect(http.StatusFound, "/")
		}
	}
}

// LoginFilter 登录拦截
func LoginFilter() gin.HandlerFunc {
	return func(c *gin.Context){
		session := sessions.Default(c)
		userId := session.Get("UserId")
		if userId == nil && c.FullPath() != "/login" && c.FullPath() != "/doLogin" && !strings.Contains(c.FullPath(), "/install") {
			c.Redirect(http.StatusFound, "/login")
		}
	}
}