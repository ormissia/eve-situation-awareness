package model

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"

	"aeon/utils"
)

type BaseParams struct {
	PageNo         int    `json:"page_no" form:"page_no"`
	PageSize       int    `json:"page_size" form:"page_size"`
	StartTimeStamp int64  `json:"start_time_stamp" form:"start_time_stamp"`
	EndTimeStamp   int64  `json:"end_time_stamp" form:"end_time_stamp"`
	TimeType       string `json:"time_type" form:"time_type"`
}

type ChartFormatData struct {
	X []interface{}
	Y map[string][]interface{}
}

var (
	ConvertErrInvalidSource  = errors.New("source is not slice")
	ConvertErrInvalidElement = errors.New("element type is not struct")
	ConvertErrInvalidField   = errors.New("element lack field")
)

func (c *ChartFormatData) Convert(xAxisName string, yAxisNames []string, source interface{}) (err error) {
	sourceSliceValue := reflect.ValueOf(source)

	if sourceSliceValue.Kind() != reflect.Slice {
		return ConvertErrInvalidSource
	}
	sliceLen := sourceSliceValue.Len()

	c.X = make([]interface{}, 0, sliceLen)
	c.Y = make(map[string][]interface{})
	for _, yAxisName := range yAxisNames {
		c.Y[yAxisName] = make([]interface{}, 0, sliceLen)
	}

	for i := 0; i < sliceLen; i++ {
		data := sourceSliceValue.Index(i)
		elementType := reflect.TypeOf(data)
		if elementType.Kind() != reflect.Struct {
			return ConvertErrInvalidElement
		}

		structMap := make(map[string]interface{})
		dataType := data.Type()
		for j := 0; j < data.NumField(); j++ {
			key := dataType.Field(j).Name
			jsonTag, ok := dataType.Field(j).Tag.Lookup("json")
			if ok {
				key = jsonTag
			}
			structMap[key] = data.Field(j).Interface()
		}

		key, ok := structMap[xAxisName]
		if !ok {
			return ConvertErrInvalidField
		}
		c.X = append(c.X, key)
		for _, yAxisName := range yAxisNames {
			y, ok := structMap[yAxisName]
			if !ok {
				return ConvertErrInvalidField
			}
			c.Y[yAxisName] = append(c.Y[yAxisName], y)
		}
	}
	return
}

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
