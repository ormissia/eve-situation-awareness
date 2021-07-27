package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"

	"admin/global"
	"admin/middleware"
	"admin/model"
	"admin/model/request"
	"admin/model/response"
	"admin/utils"
)

func GetUserInfo(c *gin.Context) {
	response.SuccessResponse(c, "user")
}

func Login(c *gin.Context) {
	var l = request.Login{}
	err := c.ShouldBind(&l)
	if err != nil {
		response.ErrorResponse(c, utils.ErrPermissionDenied)
		return
	}
	//验证
	u := &model.User{
		Username: l.Username,
		Password: l.Password,
	}
	u.Login()
	claims := request.CustomClaims{
		ID:         0,
		Username:   u.Username,
		BufferTime: global.EASConfig.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                             // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.EASConfig.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "ormissia",                                           // 签名的发行者
		},
	}
	token, err := middleware.ReleaseToken(claims)

	data := response.Login{
		User:      *u,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}

	response.SuccessResponse(c, data)
}
