package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
	"tqgin/common"
	"tqgin/models"
	"tqgin/pkg/errorcode"
	"tqgin/pkg/util"
	"tqgin/proto"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	tqgin.Controller
}

func (this *AuthController) RegisterRouter(router *gin.Engine) {
	temp := router.Group("/auth")
	temp.POST("login", this.login)
	temp.POST("register", this.register)
	temp.POST("change_pass_word", this.changePassWord)
}

func (r *AuthController) login(con *gin.Context) {
	Account := con.PostForm("Account")
	Password := con.PostForm("Password")
	Type, _ := strconv.Atoi(con.PostForm("Type"))

	fmt.Println("login data:", Account, Password, Type)

	if len(Account) <= 0 || len(Password) <= 0 {
		tqgin.Result(con, errorcode.ERROR, nil, "账号，密码不能为空")
		return
	}

	account := models.LoginAccount(Account)

	var status int
	var msg string

	var retLogin login.ReplyLogin

	if account.AccountID == "" {
		//不存在，去注册
		status = errorcode.ERROR
		msg = "不存在，去注册"
	} else {
		if account.Password != Password {
			//密码错误
			status = errorcode.ERROR
			msg = "密码错误"
		} else {
			//密码正确
			tokengen, _ := util.GenerateTocken(Account, Password)

			models.AccountSaveTocken(Account, tokengen)

			status = errorcode.SUCCESS
			msg = "登录成功"
			data := models.GetUser(account.PlayerID)
			retLogin.PlayerID = data.PlayerID
			retLogin.PlayerName = data.PlayerName
			retLogin.Diamond = data.Diamond
			retLogin.Gold = data.Gold
			retLogin.Cash = data.Cash
			retLogin.RoomID = data.RoomID
			retLogin.Token = tokengen
			retLogin.Sex = login.SexType(data.Sex)
		}
	}

	tqgin.Result(con, status, &retLogin, msg)
}

func (r *AuthController) register(con *gin.Context) {
	Account := con.PostForm("Account")
	Password := con.PostForm("Password")
	Type, _ := strconv.Atoi(con.PostForm("Type"))
	NickName := con.PostForm("NickName")
	SexType, _ := strconv.Atoi(con.PostForm("SexType"))

	fmt.Println("register data:", Account, Password, Type, NickName, SexType)

	var status int
	var msg string

	var retLogin login.ReplyLogin

	if len(Account) < 11 {
		status = errorcode.ERROR
		msg = "账号错误"
	} else if len(Password) < 6 {
		status = errorcode.ERROR
		msg = "密码不能太短"
	} else if len(NickName) <= 0 {
		status = errorcode.ERROR
		msg = "昵称不能为空"
	} else if login.SexType(SexType) < login.SexType_Sex_male && login.SexType(SexType) >= login.SexType_Sex_female {
		status = errorcode.ERROR
		msg = "性别错误"
	} else {

		var account models.Account
		account.AccountID = Account
		account.Password = Password
		account.LoginType = login.LoginType(Type)
		account.LoginTime = time.Now()

		status = models.Register(&account)

		if status == 0 {
			user := models.GetDefaultUserinfo(account.PlayerID, NickName, login.SexType(SexType))

			err := models.CreateUser(&user)
			if err == nil {
				tokengen, _ := util.GenerateTocken(Account, Password)
				msg = "注册成功"
				status = errorcode.SUCCESS

				retLogin.PlayerID = user.PlayerID
				retLogin.PlayerName = user.PlayerName
				retLogin.Diamond = user.Diamond
				retLogin.Gold = user.Gold
				retLogin.Cash = user.Cash
				retLogin.RoomID = user.RoomID
				retLogin.Token = tokengen
				retLogin.Sex = login.SexType(user.Sex)

			} else {
				msg = "创建用户信息失败"
				status = errorcode.ERROR
			}
		} else {
			msg = "注册失败，已经注册过"
			status = errorcode.ERROR
		}
	}

	tqgin.Result(con, status, &retLogin, msg)
}

func (c *AuthController) changePassWord(con *gin.Context) {

	status := errorcode.ERROR
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
			if status == errorcode.SUCCESS {
				msg = "密码修改成功"
			} else {
				msg = "密码修改失败，打印错误日志"
			}
		}
	}

	tqgin.Result(con, status, nil, msg)
}

func createloginTocken(account, password string) string {
	h := md5.New()
	h.Write([]byte(account + time.Now().String() + password))
	md5string := hex.EncodeToString(h.Sum(nil))
	return md5string
}
