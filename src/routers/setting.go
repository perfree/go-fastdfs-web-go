package routers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-fastdfs-web-go/src/commons"
	"go-fastdfs-web-go/src/form"
	"go-fastdfs-web-go/src/models"
	"net/http"
)

type SettingRouter struct {
	BaseRouter
	commons.Response
}

// UserPage 个人资料页
func (s *SettingRouter)UserPage(ctx *gin.Context) {
	user, _ := s.GetUser(ctx)
	ctx.HTML(http.StatusOK, "settings/user.tpl", gin.H{
		"user" : user,
	})
}

// EditUser 编辑用户
func (s *SettingRouter) EditUser(ctx *gin.Context)  {
	var userForm form.UserForm
	if err := ctx.ShouldBind(&userForm); err != nil {
		s.ErrorMsg(ctx, fmt.Sprint(err))
		return
	}

	var user models.User
	user, _ = user.GetById(userForm.Id)
	m5 := md5.New()
	m5.Write([]byte(userForm.Password))
	m5.Write([]byte(user.CredentialsSalt))
	st := m5.Sum(nil)
	userForm.Password = hex.EncodeToString(st)

	if user.Password != userForm.Password {
		s.ErrorMsg(ctx, "原密码错误")
		return
	}

	m5 = md5.New()
	m5.Write([]byte(userForm.NewPassword))
	m5.Write([]byte(user.CredentialsSalt))
	userForm.Password = hex.EncodeToString(m5.Sum(nil))

	newUser := userForm.GetUser()
	newUser.Update(newUser)
	s.Success(ctx)
}