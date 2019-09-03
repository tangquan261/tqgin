package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"

	"time"
	"tqgin/common"
	"tqgin/models"
	"tqgin/pkg/Agora"
	"tqgin/pkg/define"
	"tqgin/pkg/errorcode"
	"tqgin/pkg/util"

	//"tqgin/IRedis"

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
	temp.POST("phone_code", this.getPhoneCode)
}

func (r *AuthController) getPhoneCode(con *gin.Context) {

	type AuthCode struct {
		Account  string               `json:"account"`
		CodeType define.PhoneCodeType `json:"codetype"` //1注册2修改密码3绑定手机
	}

	var auth AuthCode
	err := con.ShouldBindJSON(&auth)
	if err != nil {
		tqgin.ResultFail(con, "参数错误")
		return
	}

	count := models.AuthSaveCount(auth.Account, auth.CodeType)
	if count > 5 {
		tqgin.ResultFail(con, "今日短信超额了")
		return
	}

	authmodel := models.AuthGetCode(auth.Account, auth.CodeType)

	if authmodel == nil ||
		(authmodel != nil && authmodel.UpdatedAt.Add(1*time.Minute).Before(time.Now())) {
		//时间过期了，可以继续申请
		var authNew models.AuthCode
		authNew.Account = auth.Account
		authNew.CodeType = auth.CodeType

		codeMg := CreateCaptcha()
		authNew.CodeText = codeMg

		models.AuthSaveCode(authNew)
		tqgin.ResultOkMsg(con, codeMg, "成功")

	} else {
		tqgin.ResultFail(con, "慢点获取,请稍后")
	}
}

type Authlogin struct {
	Account   string           `json:"account"`
	Password  string           `json:"password"`
	LoginType define.LoginType `json:"logintype"`
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

		data := models.GetUser(account.PlayerID)

		accountstring := strconv.FormatInt(account.PlayerID, 10)

		//IRedis.SetUserAccessTocken(account.PlayerID, 1, tokengen)
		RTMtoken, _ := tokenbuilder.RTMBuildToken("1f836f0e094446d2858f156ca366313d", "08e1620922bf40ff9ac81517f4219f51", accountstring, 1000, 0)

		tqgin.ResultOkMsg(con, gin.H{"token": tokengen, "user": data, "RTMToken": RTMtoken}, "登录成功")
		return
	}

	tqgin.ResultFail(con, msg)
}

type AuthRegister struct {
	Account   string           `json:"account"`
	Password  string           `json:"password"`
	LoginType define.LoginType `json:"logintype"`
	NickName  string           `json:"nickname"`
	SexType   define.SexType   `json:"sextype"`
	PhoneCode string           `json:"phonecode"`
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

	if len(autuparam.PhoneCode) < 4 {
		msg = "验证码错误"
	} else if len(autuparam.Account) < 11 {
		msg = "账号错误"
	} else if len(autuparam.Password) < 6 {
		msg = "密码不能太短"
	} else if len(autuparam.NickName) <= 0 {
		msg = "昵称不能为空"
	} else if autuparam.SexType < define.SexType_Sex_male &&
		autuparam.SexType >= define.SexType_Sex_female {
		msg = "性别错误"
	}

	if status == errorcode.ERROR {
		tqgin.ResultFail(con, msg)
		return
	}

	authCode := models.AuthGetCode(autuparam.Account, 1)

	if authCode == nil || authCode.CodeText != autuparam.PhoneCode ||
		authCode.UpdatedAt.Add(1*time.Minute).Before(time.Now()) {
		tqgin.ResultFail(con, "验证码错误")
		return
	}

	var account models.Account
	account.AccountID = autuparam.Account
	account.Password = autuparam.Password
	account.LoginType = autuparam.LoginType
	account.LoginTime = time.Now()

	newPlayerID, err := models.Register(account)
	if err != nil {
		tqgin.ResultFail(con, err.Error())
	} else {
		user := models.GetDefaultUserinfo(newPlayerID, autuparam.NickName, autuparam.SexType)

		err := models.CreateUser(&user)
		if err == nil {

			Newtoken, _ := util.GenerateTocken(autuparam.Account, autuparam.Password)

			var saveAccount models.Account
			saveAccount.Tocken = Newtoken
			saveAccount.TockenTimeOut = time.Now().Add(24 * 30 * time.Hour)
			models.AccountSave(autuparam.Account, saveAccount)

			msg = "注册成功"
			status = errorcode.SUCCESS
			accountstring := strconv.FormatInt(account.PlayerID, 10)

			RTMtoken, _ := tokenbuilder.RTMBuildToken("1f836f0e094446d2858f156ca366313d", "08e1620922bf40ff9ac81517f4219f51", accountstring, 1000, 0)

			//IRedis.SetUserAccessTocken(account.PlayerID, 1, Newtoken)
			tqgin.ResultOkMsg(con, gin.H{"token": Newtoken, "user": user, "RTMToken": RTMtoken}, msg)

		} else {
			msg = "注册失败"
			tqgin.ResultFail(con, msg)
		}
	}
}

type AuthChangePassword struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	PhoneCode string `json:"phonecode"`
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
		msg = "新的密码太短"
	} else {
		authCode := models.AuthGetCode(autuparam.Account, 2)

		if authCode == nil || authCode.CodeText != autuparam.PhoneCode ||
			authCode.UpdatedAt.Add(1*time.Minute).Before(time.Now()) {
			tqgin.ResultFail(con, "验证码错误")
			return
		}

		account := models.LoginAccount(autuparam.Account)

		if account == nil {
			msg = "账号不存在"
		} else {
			var save models.Account
			save.Password = autuparam.Password
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

func CreateCaptcha() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
