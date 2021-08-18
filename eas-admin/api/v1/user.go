package v1

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"admin/global"
	"admin/middleware"
	"admin/model"
	"admin/model/request"
	"admin/model/response"
	"admin/utils"
)

// GetUserInfo 获取任一用户的信息
func GetUserInfo(c *gin.Context) {
	uuid := c.Query("uuid")

	user, err := new(model.User).SelectUserByUUID(uuid)
	if err != nil {
		response.ErrorResponseCustom(c, utils.ErrCodeParamError, "未找到用户")
		return
	}

	response.SuccessResponse(c, user)
}

func Login(c *gin.Context) {
	var param = request.Login{}
	err := c.ShouldBind(&param)
	if err != nil {
		response.ErrorResponse(c, utils.ErrCodeMissingParamError)
		return
	}
	// TODO 验证
	u := &model.User{
		Username: param.Username,
		Password: param.Password,
	}
	user, err := u.Login()
	if err != nil {
		response.ErrorResponse(c, utils.ErrUserNotFoundOrErr)
		return
	}

	// 签发token
	claims := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		BufferTime:  global.EASConfig.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                             // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.EASConfig.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "eas-admin",                                          // 签名的发行者
		},
	}
	token, err := middleware.ReleaseToken(claims)

	data := response.Login{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}

	response.SuccessResponse(c, data)
}
