/*
	账号信息管理
	注册，登陆，修改密码等
*/

package controllers

import (
	"log"
	"strconv"
	"strings"
	"tqgin/common"
	"tqgin/models"
	"tqgin/pkg/errorcode"
	"tqgin/pkg/util"
	"tqgin/proto"

	//"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type UserinfoController struct {
	tqgin.Controller
}

func (this *UserinfoController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/user")
	temp.POST("login_info", this.loginInfo)
	temp.POST("update_info", this.updateInfo)
	temp.POST("get_user_info", this.getUsersInfo)
	temp.POST("add_photos", this.addPhotos)
}

func (c *UserinfoController) loginInfo(con *gin.Context) {

	playerid := con.PostForm("playerid")

	if len(playerid) <= 0 {
		tqgin.Result(con, errorcode.ERROR, nil, "用户id错误")
		return
	}

	playerID, _ := strconv.ParseInt(playerid, 10, 64)

	account := models.LoginAccountByPlayerID(playerID)

	if account.AccountID == "" {
		tqgin.Result(con, errorcode.ERROR, nil, "不存在，去注册")
	} else {
		data := models.GetUser(account.PlayerID)

		var retLogin login.ReplyLogin
		retLogin.PlayerID = data.PlayerID
		retLogin.PlayerName = data.PlayerName
		retLogin.Diamond = data.Diamond
		retLogin.Gold = data.Gold
		retLogin.Cash = data.Cash
		retLogin.RoomID = data.RoomID
		retLogin.Sex = login.SexType(data.Sex)

		tqgin.Result(con, errorcode.SUCCESS, &retLogin, "获取信息成功")
	}

}

//修改个人信息
func (c *UserinfoController) updateInfo(con *gin.Context) {

	playeridstring, _ := con.Cookie("playerid")

	PlayerGUID, err := strconv.ParseInt(playeridstring, 10, 64)

	if err != nil || PlayerGUID <= 0 {
		tqgin.ResultFail(con, "账号id错误")
		return
	}

	var user *models.UserInfo

	user = models.GetUser(PlayerGUID)
	if user == nil {
		tqgin.ResultFail(con, "没有用户信息")
		return
	}

	user.PlayerName = con.PostForm("name")
	dateString := con.PostForm("date")
	user.BirthDay = util.DateStringToTime(dateString)
	gender, _ := strconv.Atoi(con.PostForm("gender"))
	user.Sex = login.SexType(gender)
	user.Sign = con.PostForm("sign")
	user.Pic = con.PostForm("pic")
	user.Loc, _ = strconv.Atoi(con.PostForm("loc"))
	user.Locx, _ = strconv.ParseFloat(con.PostForm("locx"), 64)
	user.Locy, _ = strconv.ParseFloat(con.PostForm("locy"), 64)

	_ = models.SaveUser(PlayerGUID, user)

	tqgin.Result(con, errorcode.SUCCESS, gin.H{"playerid": PlayerGUID}, "更新数据成功")
}

func (c *UserinfoController) getUsersInfo(con *gin.Context) {

	type playerIDs struct {
		Uids []int64 `json:"uids"`
	}

	var playerGUIDS playerIDs

	err := con.ShouldBindJSON(&playerGUIDS)
	if err != nil {
		log.Println(err)
		tqgin.ResultFail(con, "param error")
		return
	}

	if len(playerGUIDS.Uids) <= 0 {
		tqgin.ResultFail(con, "error")
		return
	}

	var playerinfos []models.UserInfo
	for i := 0; i < len(playerGUIDS.Uids); i++ {
		userinfo := models.GetUser(playerGUIDS.Uids[i])
		if userinfo != nil {
			playerinfos = append(playerinfos, *userinfo)
		}
	}

	tqgin.Result(con, errorcode.SUCCESS, playerinfos, "成功")
}

type photos struct {
	Photos []string `json:"photos"`
}

func (c *UserinfoController) addPhotos(con *gin.Context) {

	playerGUID, _ := con.Cookie("playerid")

	var photo photos
	err := con.ShouldBindJSON(&photo)

	log.Println(err, photo.Photos)
	if err != nil || len(photo.Photos) <= 0 {
		tqgin.ResultFail(con, "error")
		return
	}

	photoString := strings.Join(photo.Photos, "_@_")
	var usrino models.UserInfo
	uPlayerGUID, _ := strconv.ParseInt(playerGUID, 10, 64)
	usrino.Photos = photoString
	models.SaveUser(uPlayerGUID, &usrino)

	tqgin.ResultOk(con, nil)
}
