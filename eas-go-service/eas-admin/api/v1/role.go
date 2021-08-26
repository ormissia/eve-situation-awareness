package v1

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"eas-go-service/eas-admin/model"
	"eas-go-service/eas-admin/model/response"
	"eas-go-service/global"
	"eas-go-service/utils"
)

// CreateRole 创建角色
func CreateRole(c *gin.Context) {
	param := new(model.Role)
	_ = c.ShouldBind(&param)

	// TODO 参数校验
	now := time.Now()
	param.CreateTime = now
	param.UpdateTime = now
	if param.ParentRoleId != 0 {
		// 判断该ID是否存在
		roles, err := param.Select([]uint{param.ParentRoleId}, "", 0)
		if err != nil {
			response.ErrorResponseCustom(c, utils.ErrCodeMySQLError, err.Error())
			return
		}
		if len(roles) != 1 {
			response.ErrorResponseCustom(c, utils.ErrCodeParamError, "不存在该ID的父角色")
			return
		}
	}
	if err := param.Creat(); err != nil {
		global.EASLog.Error("Create role failed", zap.Any("err", err))
		response.ErrorResponseCustom(c, utils.ErrCodeMySQLError, "角色代码重复或者其他错误")
		return
	}
	response.SuccessResponse(c, param)
}

// SearchRole 查询角色及子角色
func SearchRole(c *gin.Context) {
	idsStr := c.Query("ids")
	rolename := c.Query("rolename")
	parentRoleIdStr := c.Query("parentRoleId")

	idStrs := make([]string, 0)
	if idsStr != "" {
		idStrs = strings.Split(idsStr, ",")
	}
	ids := make([]uint, 0)
	for _, idStr := range idStrs {
		if id, err := strconv.ParseUint(idStr, 10, 64); err != nil {
			global.EASLog.Error("roleId string convert to uint failed", zap.Any("err", err))
			response.ErrorResponseCustom(c, utils.ErrCodeParamError, "角色代码传参错误")
			return
		} else {
			ids = append(ids, uint(id))
		}
	}
	var parentRoleId uint64
	if parentRoleIdStr != "" {
		parentRoleIdTemp, err := strconv.ParseUint(parentRoleIdStr, 10, 64)
		if err != nil {
			global.EASLog.Error("parentRoleId string convert to uint failed", zap.Any("err", err))
			response.ErrorResponseCustom(c, utils.ErrCodeParamError, "角色代码传参错误")
			return
		}
		parentRoleId = parentRoleIdTemp
	}
	roles, err := new(model.Role).Select(ids, rolename, uint(parentRoleId))
	if err != nil {
		global.EASLog.Error("role search failed", zap.Any("err", err))
		response.ErrorResponseCustom(c, utils.ErrCodeMySQLError, "数据库错误")
		return
	}

	response.SuccessResponse(c, roles)
}

// UpdateRole 创建角色
func UpdateRole(c *gin.Context) {
	param := new(model.Role)
	_ = c.ShouldBind(&param)

	// TODO 参数校验
	param.UpdateTime = time.Now()

	if err := param.Update(); err != nil {
		global.EASLog.Error("Create role failed", zap.Any("err", err))
		response.ErrorResponseCustom(c, utils.ErrCodeMySQLError, "角色代码重复或者其他错误")
		return
	}
	response.SuccessResponse(c, param)
}
