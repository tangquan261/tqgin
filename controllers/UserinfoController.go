/*
	账号信息管理
	注册，登陆，修改密码等
*/

package controllers

import (
	"log"
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

	temp.POST("get_mine_info", this.getUserMineInfo) //获取我的信息

	temp.POST("get_users_info", this.getUsersInfo)       //获取用户基础ids
	temp.POST("get_user_detail", this.getUserDetailInfo) //获取用户详细信息
	temp.POST("add_photos", this.addPhotos)              //添加图片
	temp.POST("update_info", this.updateInfo)            //更新用户信息

}

//获取自己我的信息
func (c *UserinfoController) getUserMineInfo(con *gin.Context) {

	type mineInfoRet struct {
		Pic         string
		Name        string
		DisPlayerID string
		FollowCount int
		FansCount   int
	}

	playerID := c.GetPlayerGUID(con)

	user := models.GetUser(playerID)

	if user == nil {
		tqgin.ResultFail(con, "失败")
	} else {

		var mineRet mineInfoRet
		mineRet.Name = user.PlayerName
		mineRet.Pic = user.Pic
		mineRet.DisPlayerID = user.DisPlayerID
		mineRet.FollowCount = models.GetFollowCount(playerID)
		mineRet.FansCount = models.GetFansCount(playerID)

		tqgin.Result(con, errorcode.SUCCESS, mineRet, "成功")
	}
}

type UserIDInfo struct {
	PlayerID int64 `json:"playerid"`
}

//获取用户详情信息
func (c *UserinfoController) getUserDetailInfo(con *gin.Context) {

	var user UserIDInfo

	err := con.ShouldBindJSON(&user)
	if err != nil {
		tqgin.ResultFail(con, "参数错误")
		return
	}

	if user.PlayerID < 0 {
		tqgin.ResultFail(con, "error")
		return
	}

	type UserInfoDetail struct {
		DisPlayerID string
		Name        string
		Pic         string
		Gender      int
		FollowCount int
		FansCount   int
		Sign        string //签名
		Photos      []string
		CityName    string //城市
		StarSign    string //星座，由出生日期算出
		Profession  string //职业
		School      string
		MarryState  string                 //婚姻状态
		Dynamic     []models.CycleModel    //动态
		GiftS       []models.GfitUserCount // 礼物列表
	}

	userinfo := models.GetUser(user.PlayerID)

	if userinfo == nil {
		tqgin.ResultFail(con, "失败")
	} else {
		var userinfoRet UserInfoDetail

		userinfoRet.DisPlayerID = userinfo.DisPlayerID
		userinfoRet.Name = userinfo.PlayerName
		userinfoRet.Pic = userinfo.Pic
		userinfoRet.Gender = int(userinfo.Sex)
		userinfoRet.FollowCount = models.GetFollowCount(user.PlayerID)
		userinfoRet.FansCount = models.GetFansCount(user.PlayerID)
		userinfoRet.Sign = userinfo.Sign
		userinfoRet.Photos = strings.Split(userinfo.Photos, ",")
		userinfoRet.CityName = userinfo.CityName
		userinfoRet.StarSign = userinfo.StarSign
		userinfoRet.Profession = userinfo.Profession
		userinfoRet.School = userinfo.School
		userinfoRet.MarryState = userinfo.MarryState
		userinfoRet.Dynamic = models.CycleGetSingleFeeds(user.PlayerID, 0, 4)
		userinfoRet.GiftS = models.GiftGetUserCount(user.PlayerID, 8)

		tqgin.Result(con, errorcode.SUCCESS, userinfoRet, "成功")
	}

}

type playerIDs struct {
	Uids []int64 `json:"uids"`
}

func (c *UserinfoController) getUsersInfo(con *gin.Context) {

	var playerGUIDS playerIDs

	err := con.ShouldBindJSON(&playerGUIDS)
	if err != nil {
		log.Println(err)
		tqgin.ResultFail(con, "param error")
		return
	}

	if len(playerGUIDS.Uids) <= 0 || len(playerGUIDS.Uids) > 10 {
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

	playerGUID := c.GetPlayerGUID(con)

	var photo photos
	err := con.ShouldBindJSON(&photo)

	if err != nil || len(photo.Photos) <= 0 {
		tqgin.ResultFail(con, "error")
		return
	}

	photoString := strings.Join(photo.Photos, ";")
	var usrino models.UserInfo
	usrino.Photos = photoString
	models.SaveUser(playerGUID, usrino)

	tqgin.ResultOk(con, nil)
}

type UserInfoParam struct {
	Name       string        `json:"name"`
	DateString string        `json:"datestring"`
	Gender     login.SexType `json:"gender"`
	Sign       string        `json:"sign"`
	Pic        string        `json:"pic"`
	LocX       float64       `json:"locx"`
	LocY       float64       `json:"locy"`
	Loc        int           `json:"loc"`
}

//修改个人信息
func (c *UserinfoController) updateInfo(con *gin.Context) {

	PlayerGUID := c.GetPlayerGUID(con)

	user := models.GetUser(PlayerGUID)
	if user == nil {
		tqgin.ResultFail(con, "没有用户信息")
		return
	}

	var userinfo UserInfoParam

	err := con.ShouldBindJSON(&userinfo)

	if err != nil {
		tqgin.ResultFail(con, "解析错误")
		return
	}

	user.PlayerName = userinfo.Name

	user.BirthDay = util.DateStringToTime(userinfo.DateString)

	user.Sex = userinfo.Gender
	user.Sign = userinfo.Sign
	user.Pic = userinfo.Pic
	user.Loc = userinfo.Loc
	user.Locx = userinfo.LocX
	user.Locy = userinfo.LocY

	models.SaveUser(PlayerGUID, *user)

	tqgin.Result(con, errorcode.SUCCESS, gin.H{"playerid": PlayerGUID}, "更新数据成功")
}
