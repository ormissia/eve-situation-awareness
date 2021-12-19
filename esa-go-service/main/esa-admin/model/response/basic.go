package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"esa-go-service/utils"
)

type response struct {
	Code utils.ResponseCode `json:"code"`
	Msg  string             `json:"msg"`
	Data interface{}        `json:"data"`
}

type responseCtx struct {
	ctx      *gin.Context
	response response
}

func (r *responseCtx) responseBase() {
	r.ctx.JSON(http.StatusOK, r.response)
}

// SuccessResponse 返回成功的同一封装
func SuccessResponse(c *gin.Context, data interface{}) {
	rc := &responseCtx{
		ctx: c,
		response: response{
			Code: http.StatusOK,
			Msg:  "Successful",
			Data: data,
		},
	}
	rc.responseBase()
}

// ErrorResponse 返回错误的统一封装
func ErrorResponse(c *gin.Context, errCode utils.ResponseCode) {
	rc := &responseCtx{
		ctx: c,
		response: response{
			Code: errCode,
			Msg:  utils.GetResponseMsg(errCode),
			Data: nil,
		},
	}
	rc.responseBase()
}

// ErrorResponseCustom 返回错误的统一封装
func ErrorResponseCustom(c *gin.Context, errCode utils.ResponseCode, msg string) {
	rc := &responseCtx{
		ctx: c,
		response: response{
			Code: errCode,
			Msg:  msg,
			Data: nil,
		},
	}
	rc.responseBase()
}

type PageResult struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}
