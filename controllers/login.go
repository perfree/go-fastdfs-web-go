package controllers

type LoginController struct {
	BaseController
}

// 登录页
func (c *LoginController) Get() {
	c.TplName = "login.html"
}
