package jwt

import (
	"net/http"
	"strconv"
	"tqgin/pkg/errorcode"
	"tqgin/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = errorcode.SUCCESS

		token, _ := c.Cookie("token")

		if token == "" {
			code = errorcode.ERROR_INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errorcode.ERROR_AUTH_TOKEN_TIMEOUT
				default:
					code = errorcode.ERROR_AUTH_TOKEN_CHECK_FAIL
				}
			}
		}

		playerstr, _ := c.Cookie("playerid")
		playerid, err := strconv.ParseInt(playerstr, 10, 64)
		if err != nil || playerid <= 0 {
			code = errorcode.ERROR_INVALID_PARAMS
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
