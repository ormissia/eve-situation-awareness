package middleware

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"

	"aeon/global"
	"aeon/main/esa-admin/model/request"
	"aeon/main/esa-admin/model/response"
	"aeon/utils"
)

// JWT 检查token
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization Header
		token := c.GetHeader("token")

		// 判断是否有token
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			response.ErrorResponse(c, utils.ErrTokenInvalid)
			c.Abort()
			return
		}

		token = token[7:]
		// 验证token格式是否合法
		claims, err := ParseToken(token)
		if err != nil {
			response.ErrorResponse(c, utils.ErrTokenInvalid)
			c.Abort()
			return
		}

		// token格式正确
		// 验证token是否过期
		if claims.ExpiresAt < time.Now().Unix() {
			response.ErrorResponse(c, utils.ErrTokenOverTime)
			c.Abort()
			return
		}

		// TODO 验证userId是否存在

		// TODO 有效期验证
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.ESAConfig.JWT.ExpiresTime
			newToken, _ := UpdateToken(token, *claims)
			newClaims, _ := ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			// 单点登录拦截
			if global.ESAConfig.Server.UseMultipoint {
				// TODO 单点登录逻辑
			}
		}

		// 验证通过将user的信息写入上下文
		c.Set("claims", claims)
		c.Next()
	}
}

var (
	TokenInvalid = errors.New("Token invalid! ")
)

func GetJWTKey() (jwtKey []byte) {
	return []byte(global.ESAConfig.JWT.SigningKey)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (claims *request.CustomClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &request.CustomClaims{},
		func(token *jwt.Token) (interface{}, error) { return GetJWTKey(), nil })
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// RelesaeToken 生成token
func RelesaeToken(claims request.CustomClaims) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(GetJWTKey())
	if err != nil {
		global.ESALog.Error("Token SignedString failed:", zap.Any("err:", err))
	}
	return
}

// UpdateToken 更新token
func UpdateToken(oldToken string, claims request.CustomClaims) (tokenStr string, err error) {
	v, err, _ := global.ESAConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return RelesaeToken(claims)
	})
	return v.(string), err
}
