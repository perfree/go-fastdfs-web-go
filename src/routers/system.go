package routers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/commons"
	"go-fastdfs-web-go/src/models"
	"net/http"
)

// SystemRouter 系统级router,如首页,登录,退出登陆等
type SystemRouter struct {
	BaseRouter
	commons.Response
}

// IndexPage 首页
func (s *SystemRouter) IndexPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "index/index.tpl", nil)
}

// LoginPage 登录页
func (s *SystemRouter) LoginPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "login/login.tpl", nil)
}

// DoLogin 登录
func (s *SystemRouter) DoLogin(c *gin.Context) {
	var user =  models.User{}
	err := c.Bind(&user)
	if err != nil {
		s.ErrorMsg(c, "请求参数错误")
		return
	}
	if user.Account == "" {
		s.ErrorMsg(c, "账户不能为空")
		return
	}
	if user.Password == "" {
		s.ErrorMsg(c, "密码不能为空")
		return
	}

	password := user.Password
	queryUser, err := user.GetByAccount(user.Account)

	if err != nil {
		s.ErrorMsg(c, "账户不存在")
		return
	}

	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(queryUser.CredentialsSalt))
	st := m5.Sum(nil)
	user.Password = hex.EncodeToString(st)

	if user.Password == queryUser.Password {
		session := sessions.Default(c)
		session.Set("UserId", queryUser.Id)
		_ = session.Save()
		s.Success(c)
		return
	}
	s.ErrorMsg(c, "密码错误")
}

// LogOut 退出
func (s *SystemRouter) LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("UserId")
	_ = session.Save()
	c.Redirect(http.StatusFound, "/login")
}