/*
	账号信息管理
	注册，登陆，修改密码等
*/

package controllers

import (
	"strconv"
	"tqgin/common"
	"tqgin/models"

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

	accountID := con.PostForm("accountid")

	pwd := con.PostForm("password")

	account := models.LoginAccount(accountID)

	var status int
	var msg string

	if account.AccountID == "" {
		//不存在，去注册
		status = 1
		msg = "不存在，去注册"
	} else {

		if account.Password != pwd {
			//密码错误
			status = 2
			msg = "密码错误"
		} else {
			//密码正确
			status = 0
			msg = "登录成功"
		}
	}

	tqgin.Result(con, status, nil, msg)
}

func (c *LoginController) register(con *gin.Context) {

	var account models.Account
	var retAccount *models.Account
	var status int
	var msg string

	account.AccountID = con.PostForm("accountid")
	account.Password = con.PostForm("password")
	loginType, _ := strconv.Atoi(con.PostForm("loginType"))

	account.LoginType = int8(loginType)

	if len(account.AccountID) < 11 {
		status = 1
		msg = "账号错误"
	} else if len(account.Password) < 6 {
		status = 2
		msg = "密码不能太短"
	} else {

		status, retAccount = models.Register(&account)

		if status == 0 {
			msg = "注册成功"
		} else {
			msg = "注册失败，已经注册过"
		}
	}

	tqgin.Result(con, status, retAccount, msg)
}

func (c *LoginController) changePassWord(con *gin.Context) {

	status := 1
	var msg string

	accountID := con.PostForm("accountid")
	oldPassword := con.PostForm("oldpassword")
	newPassword := con.PostForm("newpassword")

	if len(accountID) < 11 {
		msg = "账号错误"
	} else if len(oldPassword) < 6 {
		msg = "当前密码错误"
	} else if len(newPassword) < 6 {
		msg = "新的密码太短"
	} else if newPassword == oldPassword {
		msg = "新旧密码不能相同"
	} else {

		account := models.LoginAccount(accountID)

		if len(account.AccountID) < 11 {
			msg = "账号不存在"
		} else if account.Password != oldPassword {
			msg = "密码错误"
		} else {
			account.Password = newPassword
			status = models.AccountChangePwd(account, newPassword)
			if status == 0 {
				msg = "密码修改成功"
			} else {
				msg = "密码修改失败，打印错误日志"
			}
		}
	}

	tqgin.Result(con, status, nil, msg)
}
