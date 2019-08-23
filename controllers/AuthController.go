package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"time"
	"tqgin/common"
	"tqgin/models"
	"tqgin/pkg/Agora"
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

type Authlogin struct {
	Account   string          `json:"account"`
	Password  string          `json:"passowrd"`
	LoginType login.LoginType `json:"logintype"`
	//NickName  string          `json:"nickname"`
	//SexType   login.SexType   `json:"sextype"`
}

func (r *AuthController) login(con *gin.Context) {

	var autuparam Authlogin

	err := con.ShouldBindJSON(&autuparam)

	if err != nil {
		tqgin.ResultFail(con, "登录失败")
		return
	}

	if len(autuparam.Account) <= 0 || len(autuparam.Password) <= 0 {
		tqgin.ResultFail(con, "登录失败")
		return
	}

	account := models.LoginAccount(autuparam.Account)

	if account == nil {
		tqgin.ResultFail(con, "账号错误")
		return
	}

	var msg string

	if account.AccountID == "" {
		//不存在，去注册
		msg = "账号错误"
	} else if account.Password != autuparam.Password {
		//密码错误
		msg = "账号错误"
	} else if account.ForbidTime > 0 && account.ForbidTime < time.Now().Unix() {
		msg = "被封禁时间，请联系客服"
	} else {
		//密码正确
		tokengen, _ := util.GenerateTocken(autuparam.Account, autuparam.Password)

		var saveAccount models.Account
		saveAccount.Tocken = tokengen
		models.AccountSave(account.AccountID, saveAccount)

		msg = "登录成功"
		data := models.GetUser(account.PlayerID)

		accountstring := strconv.FormatInt(account.PlayerID, 10)

		RTMtoken, _ := tokenbuilder.RTMBuildToken("1f836f0e094446d2858f156ca366313d", "08e1620922bf40ff9ac81517f4219f51", accountstring, 1000, 0)

		tqgin.ResultOkMsg(con, gin.H{"token": tokengen, "user": data, "RTMTocken": RTMtoken}, "登录成功")
		return
	}

	tqgin.ResultFail(con, msg)
}

type AuthRegister struct {
	Account   string          `json:"account"`
	Password  string          `json:"passowrd"`
	LoginType login.LoginType `json:"logintype"`
	NickName  string          `json:"nickname"`
	SexType   login.SexType   `json:"sextype"`
}

func (r *AuthController) register(con *gin.Context) {

	var autuparam AuthRegister

	err := con.ShouldBindJSON(&autuparam)

	if err != nil {
		tqgin.ResultFail(con, "注册失败")
		return
	}

	var msg string
	var status int
	if len(autuparam.Account) < 11 {
		status = errorcode.ERROR
		msg = "账号错误"
	} else if len(autuparam.Password) < 6 {
		status = errorcode.ERROR
		msg = "密码不能太短"
	} else if len(autuparam.NickName) <= 0 {
		status = errorcode.ERROR
		msg = "昵称不能为空"
	} else if autuparam.SexType < login.SexType_Sex_male && autuparam.SexType >= login.SexType_Sex_female {
		status = errorcode.ERROR
		msg = "性别错误"
	}
	if status == errorcode.ERROR {
		tqgin.ResultFail(con, msg)
		return
	}

	var account models.Account
	account.AccountID = autuparam.Account
	account.Password = autuparam.Password
	account.LoginType = autuparam.LoginType
	account.LoginTime = time.Now()

	err = models.Register(account)
	if err != nil {
		msg = "注册失败"
		tqgin.ResultFail(con, msg)
	} else {
		user := models.GetDefaultUserinfo(account.PlayerID, autuparam.NickName, autuparam.SexType)

		err := models.CreateUser(&user)
		if err == nil {
			Newtoken, _ := util.GenerateTocken(autuparam.Account, autuparam.Password)
			msg = "注册成功"
			status = errorcode.SUCCESS
			accountstring := strconv.FormatInt(account.PlayerID, 10)

			RTMtoken, _ := tokenbuilder.RTMBuildToken("1f836f0e094446d2858f156ca366313d", "08e1620922bf40ff9ac81517f4219f51", accountstring, 1000, 0)

			tqgin.ResultOkMsg(con, gin.H{"token": Newtoken, "user": user, "RTMTocken": RTMtoken}, msg)

		} else {
			msg = "注册失败"
			tqgin.ResultFail(con, msg)
		}
	}
}

type AuthChangePassword struct {
	Account      string `json:"account"`
	Password     string `json:"passowrd"`
	NewPassoword string `json:"newpassword"`
}

func (c *AuthController) changePassWord(con *gin.Context) {

	var autuparam AuthChangePassword

	err := con.ShouldBindJSON(&autuparam)

	if err != nil {
		tqgin.ResultFail(con, "注册失败")
		return
	}

	var msg string

	if len(autuparam.Account) < 11 {
		msg = "账号错误"
	} else if len(autuparam.Password) < 6 {
		msg = "当前密码太短"
	} else if len(autuparam.NewPassoword) < 6 {
		msg = "新的密码太短"
	} else if autuparam.NewPassoword == autuparam.Password {
		msg = "新旧密码不能相同"
	} else {
		account := models.LoginAccount(autuparam.Account)

		if account == nil {
			msg = "账号不存在"
		} else if account.Password != autuparam.Password {
			msg = "密码错误"
		} else {
			var save models.Account
			save.Password = autuparam.NewPassoword
			models.AccountSave(autuparam.Account, save)
			tqgin.ResultOk(con, "修改密码成功")
			return
		}
	}

	tqgin.ResultFail(con, msg)
}

func createloginTocken(account, password string) string {
	h := md5.New()
	h.Write([]byte(account + time.Now().String() + password))
	md5string := hex.EncodeToString(h.Sum(nil))
	return md5string

}
