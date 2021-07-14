package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResult struct {
	Total    int64       `json:"total"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	List     interface{} `json:"list"`
}

func (r *Response) response(c *gin.Context, code int) {
	c.JSON(code, r)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	r := Response{
		Code: http.StatusOK,
		Msg:  "Successful",
		Data: data,
	}
	r.response(c, http.StatusOK)
}
