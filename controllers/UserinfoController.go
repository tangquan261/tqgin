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

	temp.POST("update_info", this.updateInfo)     //更新用户信息
	temp.POST("get_user_info", this.getUsersInfo) //获取用户id
	temp.POST("add_photos", this.addPhotos)       //添加图片
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

	log.Println(err, photo.Photos)
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
