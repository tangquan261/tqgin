/*
	账号信息管理
	注册，登陆，修改密码等
*/

package controllers

import (
	//"encoding/json"
	"fmt"

	//"net/http"
	"strconv"
	"tqgin/common"
	"tqgin/models"
	"tqgin/proto"

	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	"github.com/golang/protobuf/proto"
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

	var logindata login.ApplyLogin

	data := con.PostForm("data")

	err := proto.Unmarshal([]byte(data), &logindata)

	fmt.Println("apply login: %v", logindata)

	cookie, err := con.Cookie("token")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("cookie:", cookie)

	if len(logindata.Account) <= 0 || len(logindata.Password) <= 0 {

		tqgin.Result(con, 1, nil, "账号，密码不能为空")
		return
	}

	account := models.LoginAccount(logindata.Account)

	var status int
	var msg string

	var retLogin login.ReplyLogin

	retLogin.Code = 1
	retLogin.Errinfo = "没有暖用"

	//var data *models.UserInfo
	if account.AccountID == "" {
		//不存在，去注册
		status = 1
		msg = "不存在，去注册"
	} else {

		if account.Password != logindata.Password {
			//密码错误
			status = 2
			msg = "密码错误"
		} else {
			//密码正确
			status = 0
			msg = "登录成功"
			//data = models.GetUser(account.PlayerID)
		}
	}

	tqgin.Result(con, status, &retLogin, msg)
}

func (c *LoginController) register(con *gin.Context) {

	var account models.Account
	var retAccount *models.Account
	var status int
	var msg string

	account.AccountID = con.PostForm("accountid")
	account.Password = con.PostForm("password")
	loginType, _ := strconv.Atoi(con.PostForm("loginType"))

	nickName := con.PostForm("nickname")
	sex, _ := strconv.Atoi(con.PostForm("sex"))

	account.LoginType = int8(loginType)

	fmt.Println(account.AccountID, account.Password, account.LoginType)
	if len(account.AccountID) < 11 {
		status = 1
		msg = "账号错误"
	} else if len(account.Password) < 6 {
		status = 2
		msg = "密码不能太短"
	} else if len(nickName) <= 0 {
		status = 2
		msg = "昵称不能为空"
	} else {

		status, retAccount = models.Register(&account)

		if status == 0 {

			user := models.GetDefaultUserinfo(retAccount.PlayerID, nickName, int8(sex))

			bret := models.CreateUser(&user)
			if bret {
				msg = "注册成功"
				status = 0
			} else {
				msg = "创建用户信息失败"
			}
		} else {
			msg = "注册失败，已经注册过"
		}
	}
	fmt.Println(msg, status)
	//tqgin.Result(con, status, retAccount, msg)
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
