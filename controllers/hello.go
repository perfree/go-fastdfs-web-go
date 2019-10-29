package controllers

import (
	"go-fastdfs-web-go/models"
)

type HelloController struct {
	BaseController
}

// 获取所有用户信息
func (c *HelloController) Get() {
	users := models.GetUsers()
	c.Data["json"] = &JsonData{Code: SUCCESS_CODE, Count: len(users), Msg: "获取数据成功", Data: users}
	//c.ServeJSON()
	c.TplName = "install.html"
}

// 获取所有用户信息
func (c *HelloController) GetUsers() {
	users := models.GetUsers()
	c.Data["json"] = &JsonData{Code: SUCCESS_CODE, Count: len(users), Msg: "获取数据成功", Data: users}
	c.ServeJSON()
}
