/*
	账号信息管理
	注册，登陆，修改密码等
*/

package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"time"

	//"encoding/json"
	//"fmt"

	//"net/http"
	//"strconv"
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
	temp.POST("loginInfo", this.loginInfo)
	temp.POST("changePassWord", this.changePassWord)
}

func (c *LoginController) login(con *gin.Context) {

	var logindata login.ApplyLogin

	data := con.PostForm("data")

	err := proto.Unmarshal([]byte(data), &logindata)
	if err != nil {
		log.Fatalln(err)
	}

	if len(logindata.Account) <= 0 || len(logindata.Password) <= 0 {

		tqgin.Result(con, 1, nil, "账号，密码不能为空")
		return
	}

	account := models.LoginAccount(logindata.Account)

	var status int
	var msg string

	var retLogin login.ReplyLogin

	if account.AccountID == "" {
		//不存在，去注册
		status = 1
		msg = "不存在，去注册"
	} else {
		fmt.Println("passwod", account.Password, "----", logindata.Password, logindata.Account)
		if account.Password != logindata.Password {
			//密码错误
			status = 2
			msg = "密码错误"
		} else {
			//密码正确
			//account.ForbidTime
			md5string := createloginTocken(account.AccountID, account.Password)
			account.Tocken = md5string
			account.TockenTimeOut = time.Now().Add(time.Duration(30) * time.Hour * 24)

			models.AccountSaveTocken(account)

			status = 0
			msg = "登录成功"
			data := models.GetUser(account.PlayerID)
			retLogin.PlayerID = data.PlayerID
			retLogin.PlayerName = data.PlayerName
			retLogin.Diamond = data.Diamond
			retLogin.Gold = data.Gold
			retLogin.Cash = data.Cash
			retLogin.RoomID = data.RoomID
			retLogin.Token = md5string
			retLogin.Sex = login.SexType(data.Sex)
		}
	}

	tqgin.Result(con, status, &retLogin, msg)
}

func (c *LoginController) register(con *gin.Context) {

	var registerData login.RegisterInfo

	data := con.PostForm("data")

	err := proto.Unmarshal([]byte(data), &registerData)
	if err != nil {
		log.Fatalln(err)
	}

	var status int
	var msg string

	var retLogin login.ReplyLogin
	fmt.Println(registerData)
	if len(registerData.Account) < 11 {
		status = 1
		msg = "账号错误"
	} else if len(registerData.Password) < 6 {
		status = 2
		msg = "密码不能太短"
	} else if len(registerData.NickNmae) <= 0 {
		status = 3
		msg = "昵称不能为空"
	} else if registerData.SexType < login.SexType_Sex_male && registerData.SexType >= login.SexType_Sex_female {
		status = 4
		msg = "性别错误"
	} else {

		var account models.Account

		account.AccountID = registerData.Account
		account.Password = registerData.Password
		account.LoginType = registerData.Type
		account.LoginTime = time.Now()
		md5string := createloginTocken(account.AccountID, account.Password)
		account.Tocken = md5string
		account.TockenTimeOut = time.Now().Add(time.Duration(30) * time.Hour * 24)

		status = models.Register(&account)

		if status == 0 {

			user := models.GetDefaultUserinfo(account.PlayerID, registerData.NickNmae, registerData.SexType)

			bret := models.CreateUser(&user)
			if bret {
				msg = "注册成功"
				status = 0

				retLogin.PlayerID = user.PlayerID
				retLogin.PlayerName = user.PlayerName
				retLogin.Diamond = user.Diamond
				retLogin.Gold = user.Gold
				retLogin.Cash = user.Cash
				retLogin.RoomID = user.RoomID
				retLogin.Token = md5string
				retLogin.Sex = login.SexType(user.Sex)

			} else {
				msg = "创建用户信息失败"
				status = 1
			}
		} else {
			msg = "注册失败，已经注册过"
			status = 1
		}
	}

	tqgin.Result(con, status, &retLogin, msg)
}

func (c *LoginController) loginInfo(con *gin.Context) {

	playerIDstr, _ := con.Cookie("playerid")
	token, _ := con.Cookie("token")

	if len(playerIDstr) <= 0 || len(token) <= 0 {
		tqgin.Result(con, 1, nil, "登录失败，请重新登录")
		return
	}

	playerID, _ := strconv.ParseInt(playerIDstr, 10, 64)

	account := models.LoginAccountByPlayerID(playerID)

	var status int
	var msg string

	var retLogin login.ReplyLogin

	if account.AccountID == "" {
		//不存在，去注册
		status = 1
		msg = "不存在，去注册"
	} else {
		fmt.Println(account.Tocken, "current:", token)
		if account.Tocken != token {
			//密码错误
			status = 2
			msg = "密码错误"
		} else {
			//密码正确
			//account.ForbidTime
			status = 0
			msg = "登录成功"
			data := models.GetUser(account.PlayerID)
			retLogin.PlayerID = data.PlayerID
			retLogin.PlayerName = data.PlayerName
			retLogin.Diamond = data.Diamond
			retLogin.Gold = data.Gold
			retLogin.Cash = data.Cash
			retLogin.RoomID = data.RoomID
			retLogin.Token = token
			retLogin.Sex = login.SexType(data.Sex)
		}
	}

	tqgin.Result(con, status, &retLogin, msg)
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

func createloginTocken(account, password string) string {

	h := md5.New()
	h.Write([]byte(account + time.Now().String() + password))
	md5string := hex.EncodeToString(h.Sum(nil))
	return md5string
}
