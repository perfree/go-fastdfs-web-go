package commons

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

// SuccessMsgAndData 成功 带msg和data
func (r *Response) SuccessMsgAndData(ctx *gin.Context, msg string, data interface{}) {
	res := Response{
		200, msg, data,
	}
	ctx.JSON(http.StatusOK, res)
}

// SuccessMsg 成功 带msg
func (r *Response) SuccessMsg(ctx *gin.Context, msg string) {
	res := Response{
		200, msg, "",
	}
	ctx.JSON(http.StatusOK, res)
}

// Success 成功
func (r *Response) Success(ctx *gin.Context) {
	res := Response{
		200, "", "",
	}
	ctx.JSON(http.StatusOK, res)
}

// SuccessData 成功 带数据
func (r *Response) SuccessData(ctx *gin.Context, data interface{}) {
	res := Response{
		200, "", data,
	}
	ctx.JSON(http.StatusOK, res)
}

// Json json数据
func (r *Response) Json(ctx *gin.Context, code int, msg string, data interface{}) {
	res := Response{
		code, msg, data,
	}
	ctx.JSON(http.StatusOK, res)
}

// ErrorMsgAndData 失败 带msg和data
func (r *Response) ErrorMsgAndData(ctx *gin.Context, msg string, data interface{}) {
	res := Response{
		500, "", data,
	}
	ctx.JSON(http.StatusOK, res)
}


// Error 失败
func (r *Response) Error(ctx *gin.Context) {
	res := Response{
		500, "", "",
	}
	ctx.JSON(http.StatusOK, res)
}

// ErrorData 失败 带数据
func (r *Response) ErrorData(ctx *gin.Context, data interface{}) {
	res := Response{
		500, "", data,
	}
	ctx.JSON(http.StatusOK, res)
}

// ErrorMsg 失败 带msg
func (r *Response) ErrorMsg(ctx *gin.Context, msg string) {
	res := Response{
		500, msg, "",
	}
	ctx.JSON(http.StatusOK, res)
}

