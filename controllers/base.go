package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

// 结果响应状态码:成功
const SUCCESS_CODE  = 0

// 结果响应状态码:失败
const SUCCESS_FAIL  = 1

// 返回json列表 数据格式
type JsonData struct {
	Code  int         `json:"code"`  //错误代码
	Msg   string      `json:"msg"`   //输出信息
	Count int         `json:"total"` // 数据数量
	Data  interface{} `json:"data"`  //数据
}