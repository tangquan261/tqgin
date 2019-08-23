package jwt

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"tqgin/pkg/errorcode"

	"tqgin/pkg/tqlog"
	"tqgin/pkg/util"

	"tqgin/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

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
		} else {
			account := models.LoginAccountByPlayerID(playerid)
			if account == nil || account.Tocken != token {
				code = errorcode.ERROR_AUTH_TOKEN_CHECK_FAIL
			}
		}

		//打印请求日志
		rawData, _ := c.GetRawData()

		tqlog.TQRequest.Info("playerid:", playerid, "code:", code,
			"method:", c.Request.Method, "conentType:", c.ContentType(), "body:", string(rawData))

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData)) // 关键点

		if code != errorcode.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
				"data": rawData,
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
