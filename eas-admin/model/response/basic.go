package response

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"admin/utils"
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
	r.ctx.JSON(int(r.response.Code), r.response)
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
			Code: http.StatusOK,
			Msg:  utils.GetResponseMsg(errCode),
			Data: nil,
		},
	}
	rc.responseBase()
}

type PageResult struct {
	Total    int64       `json:"total"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	List     interface{} `json:"list"`
}
