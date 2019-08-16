/*
	房间管理
	创建，移除，修改操作
*/

package controllers

import (
	"tqgin/common"
	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type RoomManagerController struct {
	tqgin.Controller
}

func (this *RoomManagerController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/room_manager")
	temp.POST("app_open_room", this.applyOpenRoom)     //请求创建房间,获取创建房间配置信息
	temp.POST("open_room", this.OpenRoom)              //请求创建
	temp.POST("close_room", this.closeRoom)            //关闭房间
	temp.POST("change_roomName", this.ChangeRoomName)  //修改房间信息
	temp.POST("apply_enter_Room", this.applyEnterRoom) //进入房间
	temp.POST("apply_roominfo", this.applyRoomInfo)    //请求房间信息
}

func (r *RoomManagerController) applyOpenRoom(con *gin.Context) {

}

//开启房间
func (r *RoomManagerController) OpenRoom(con *gin.Context) {

	type OpenRoomParam struct {
		RoomName      string `json:"roomname"`      //房间名字
		RoomIntro     string `json:"roomdetail"`    //房间介绍
		RoomNotice    string `json:"roomnotice"`    //介绍公告
		RoomTag       string `json:"roomtags"`      //房间类型
		RoomPic       string `json:"roompic"`       //房间头像
		RoomAudioType int32  `json:"roomaudiotype"` //房间声音类型
	}

	var createRoominfo OpenRoomParam

	err := con.ShouldBindJSON(&createRoominfo)
	if err != nil {
		tqgin.ResultFail(con, "创建失败")
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

}

//关闭房间
func (r *RoomManagerController) ChangeRoomName(con *gin.Context) {

}

func (r *RoomManagerController) applyEnterRoom(con *gin.Context) {

}

//根据id获取房间信息
func (r *RoomManagerController) applyRoomInfo(con *gin.Context) {

}
