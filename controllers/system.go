package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"go-fastdfs-web-go/models"
)

type SystemController struct {
	BaseController
}

// Home 首页
func (c *SystemController) Home() {
	c.TplName = "index.tpl"
}

// Login Login页面
func (c *SystemController) Login() {
	c.TplName = "login.tpl"
}

// DoLogin 登录
func (c *SystemController) DoLogin() {
	var user = models.User{}
	err := c.ParseForm(&user)
	if err != nil {
		c.ErrorJson(500, "请求参数错误", nil)
	}
	if user.Account == "" {
		c.ErrorJson(500, "账户不能为空", nil)
	}
	if user.Password == "" {
		c.ErrorJson(500, "密码不能为空", nil)
	}
	password := user.Password
	queryUser, err := user.GetByAccount()

	if err != nil {
		c.ErrorJson(500, "账户不存在", nil)
	}

	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(queryUser.CredentialsSalt))
	st := m5.Sum(nil)
	user.Password = hex.EncodeToString(st)

	if user.Password == queryUser.Password {
		c.SetSession("userId", user.Id)
		c.SuccessJson("success")
	}

	c.ErrorJson(500, "密码错误", nil)
}
