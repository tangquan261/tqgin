/*
	房间管理
	创建，移除，修改操作
*/

package controllers

import (
	//"fmt"
	"strconv"
	"tqgin/common"
	"tqgin/models"
	"tqgin/pkg/Agora"

	"github.com/gin-gonic/gin"
)

type RoomManagerController struct {
	tqgin.Controller
}

func (this *RoomManagerController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/room_manager")
	temp.POST("config_room", this.applyConfigRoom)     //请求创建房间,获取创建房间配置信息
	temp.POST("open_room", this.OpenRoom)              //请求创建
	temp.POST("close_room", this.closeRoom)            //关闭房间
	temp.POST("change_roominfo", this.ChangeRoomInfo)  //修改房间信息
	temp.POST("apply_roominfo", this.applyRoomInfo)    //请求房间信息
	temp.POST("apply_add_admin", this.applyAddAdmin)   //添加房间管理
	temp.POST("apply_del_admin", this.applyDelAdmin)   //移除房间管理
	temp.POST("apply_enter_Room", this.applyEnterRoom) //进入房间
	//添加或者移除房间管理员
	//添加或者移除房间黑名单
	//获取房间管理员和黑名单的列表
}

func (r *RoomManagerController) applyEnterRoom(con *gin.Context) {

	type EnterRoomParam struct {
		RoomID int64 `json:"roomid"` //房间id
	}

	var enterRoom EnterRoomParam
	err := con.ShouldBindJSON(&enterRoom)
	if err != nil {
		tqgin.ResultFailData(con, nil, "错误")
		return
	}

	playerID := r.GetPlayerGUID(con)
	channel := "channel" + strconv.FormatInt(enterRoom.RoomID, 10)

	token, _ := tokenbuilder.BuildTokenWithUID("1f836f0e094446d2858f156ca366313d",
		"08e1620922bf40ff9ac81517f4219f51", channel, uint32(playerID),
		tokenbuilder.RolePublisher, 0)

	room := models.GetRoomById(enterRoom.RoomID)
	micsqueue := models.MicQueueInfo(enterRoom.RoomID)
	mics := models.MicGetAllIndex(enterRoom.RoomID)

	tqgin.ResultOkMsg(con, gin.H{"token": token, "room": room,
		"micqueue": micsqueue, "mics": mics}, "获取成功")
}

func (r *RoomManagerController) applyConfigRoom(con *gin.Context) {

	playerID := r.GetPlayerGUID(con)

	myRoom := models.GetRoomById(playerID)

	tags := models.GetTagList()

	tqgin.ResultOkMsg(con, gin.H{"myroom": myRoom, "tags": tags}, "成功")
}

type OpenRoomParam struct {
	RoomName      string `json:"roomname"`      //房间名字
	RoomIntro     string `json:"roomintro"`     //房间介绍
	RoomNotice    string `json:"roomnotice"`    //介绍公告
	RoomTag       string `json:"roomtag"`       //房间Tag类型
	RoomPic       string `json:"roompic"`       //房间头像
	RoomAudioType int32  `json:"roomaudiotype"` //房间声音类型
}

//开启房间
func (r *RoomManagerController) OpenRoom(con *gin.Context) {

	var createRoominfo OpenRoomParam

	err := con.ShouldBindJSON(&createRoominfo)
	if err != nil {
		tqgin.ResultFail(con, "参数错误")
		return
	}

	playerID := r.GetPlayerGUID(con)

	var roominfo models.RoomInfo
	roominfo.RoomID = playerID

	roominfo.RoomName = createRoominfo.RoomName
	roominfo.RoomIntro = createRoominfo.RoomIntro
	roominfo.RoomNotice = createRoominfo.RoomNotice
	roominfo.RoomTag = createRoominfo.RoomTag
	roominfo.RoomPic = createRoominfo.RoomPic
	roominfo.RoomAudioType = createRoominfo.RoomAudioType

	err = models.CreateRoom(roominfo)
	if err != nil {
		tqgin.ResultFail(con, "创建失败")
	} else {
		tqgin.ResultOkMsg(con, createRoominfo, "创建成功")
	}
}

//关闭房间
func (r *RoomManagerController) closeRoom(con *gin.Context) {

	playerID := r.GetPlayerGUID(con)

	models.CloseHotRoom(playerID)

	tqgin.ResultOkMsg(con, nil, "关闭成功")
}

type RoomIDParam struct {
	roomid int64 `json:"roomid"`
}

//根据id获取房间信息
func (r *RoomManagerController) applyRoomInfo(con *gin.Context) {
	var roomIDInfo RoomIDParam

	err := con.ShouldBindJSON(&roomIDInfo)
	if err != nil {
		tqgin.ResultFail(con, "参数错误")
		return
	}

	retmodel := models.GetRoomById(roomIDInfo.roomid)

	tqgin.ResultOkMsg(con, retmodel, "成功")
}

//修改房间名字
func (r *RoomManagerController) ChangeRoomInfo(con *gin.Context) {

	playerID := r.GetPlayerGUID(con)

	type RoomInfoParam struct {
		RoomID     int64  `json:"roomid"`
		RoomName   string `json:"roomname"`   //房间名字
		RoomIntro  string `json:"roomintro"`  //房间介绍
		RoomNotice string `json:"roomnotice"` //介绍公告
		RoomPic    string `json:"roompic"`    //房间头像
		RoomBGPic  string `json:"roombgpic"`  //房间背景
	}
	var roomparam RoomInfoParam

	err := con.ShouldBindJSON(&roomparam)
	if err != nil || playerID != roomparam.RoomID {
		tqgin.ResultFail(con, "参数错误")
		return
	}

	retmodel := models.GetRoomById(roomparam.RoomID)

	if retmodel == nil {
		tqgin.ResultFail(con, "房间不存在")
		return
	}

	var room models.RoomInfo
	if len(roomparam.RoomName) > 0 {
		room.RoomName = roomparam.RoomName
	}

	if len(roomparam.RoomIntro) > 0 {
		room.RoomIntro = roomparam.RoomIntro
	}

	if len(roomparam.RoomNotice) > 0 {
		room.RoomNotice = roomparam.RoomNotice
	}
	if len(roomparam.RoomPic) > 0 {
		room.RoomPic = roomparam.RoomPic
	}
	if len(roomparam.RoomBGPic) > 0 {
		room.RoomBGPic = roomparam.RoomBGPic
	}

	err = models.SaveRoominfo(playerID, room)
	if err != nil {
		tqgin.ResultFail(con, "修改房间失败")
	} else {
		tqgin.ResultOkMsg(con, room, "成功")
	}
}

func (r *RoomManagerController) applyAddAdmin(con *gin.Context) {

	playerID := r.GetPlayerGUID(con)

	type RoomInfoParam struct {
		RoomID      int64 `json:"roomid"`
		TarPlayerID int64 `json:"tarplayerid"`
	}
	var roomparam RoomInfoParam

	err := con.ShouldBindJSON(&roomparam)
	if err != nil || playerID != roomparam.RoomID {
		tqgin.ResultFail(con, "参数错误")
		return
	}
}
func (r *RoomManagerController) applyDelAdmin(con *gin.Context) {
	playerID := r.GetPlayerGUID(con)

	type RoomInfoParam struct {
		RoomID      int64 `json:"roomid"`
		TarPlayerID int64 `json:"tarplayerid"`
	}
	var roomparam RoomInfoParam

	err := con.ShouldBindJSON(&roomparam)
	if err != nil || playerID != roomparam.RoomID {
		tqgin.ResultFail(con, "参数错误")
		return
	}
}
