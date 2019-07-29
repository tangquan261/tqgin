package jwt

import (
	"net/http"
	"tqgin/pkg/errorcode"
	"tqgin/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		tocken := c.PostForm("tocken")
		if tocken == "" {
			code = errorcode.ERROR_INVALID_PARAMS
		} else {
			_, err := util.ParseToken(tocken)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errorcode.ERROR_AUTH_TOKEN_TIMEOUT
				default:
					code = errorcode.ERROR_AUTH_TOKEN_CHECK_FAIL
				}
			}
		}

		if code != errorcode.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
