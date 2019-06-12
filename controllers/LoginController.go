/*
	账号信息管理
	注册，登陆，修改密码等
*/

package controllers

import (
	"fmt"
	"tqTestGin/tqgin"

	"github.com/gin-gonic/gin"
)

var nindex int

type LoginController struct {
	tqgin.Controller
}

func (this *LoginController) RegisterRouter(router *gin.Engine) {
	temp := router.Group("/account")
	temp.POST("login", this.login)
	temp.POST("register", this.register)
	temp.POST("changePassWord", this.changePassWord)
}

func (c *LoginController) login(con *gin.Context) {

	account := con.Query("account")
	pwd := con.PostForm("password")

	tqgin.ResultOkMsg(con, gin.H{
		"status":  "ok",
		"account": account,
		"pwd":     pwd,
	}, "lgoin success")
}

func (c *LoginController) register(con *gin.Context) {
	fmt.Println("register")
	tqgin.ResultOk(con, "")

}

func (c *LoginController) changePassWord(con *gin.Context) {
	fmt.Println("register")

	tqgin.ResultOkMsg(con, "", "success")
}
