/*
	房间管理
	创建，移除，修改操作
*/

package controllers

import (
	"fmt"
	"log"
	"strconv"

	"tqgin/common"
	"tqgin/models"
	"tqgin/proto"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
)

type RoomManagerController struct {
	tqgin.Controller
}

func (this *RoomManagerController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/roomManager")
	temp.POST("openRoom", this.OpenRoom)
	temp.POST("closeRoom", this.closeRoom)
	temp.POST("changeRoomName", this.ChangeRoomName)
	temp.POST("applyEnterRoom", this.applyEnterRoom)
}

//开启房间
func (r *RoomManagerController) OpenRoom(con *gin.Context) {

	var createRoominfo login.ApplyCreateRoom

	data := con.PostForm("data")

	err := proto.Unmarshal([]byte(data), &createRoominfo)
	if err != nil {
		log.Fatalln(err)
	}

	playerIDstr, _ := con.Cookie("playerid")
	//token, _ := con.Cookie("token")

	playerID, _ := strconv.ParseInt(playerIDstr, 10, 64)

	var roominfo models.RoomInfo
	roominfo.RoomName = createRoominfo.RoomName
	roominfo.RoomID = playerID
	roominfo.MasterID = playerID
	roominfo.RoomTagName = createRoominfo.RoomTags

	err = models.CreateRoom(&roominfo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("请求创建房间成功", createRoominfo, roominfo)

	tqgin.Result(con, 0, &createRoominfo, "请求创建房间")

}

//关闭房间
func (r *RoomManagerController) closeRoom(con *gin.Context) {

}

//关闭房间
func (r *RoomManagerController) ChangeRoomName(con *gin.Context) {

}

func (r *RoomManagerController) applyEnterRoom(con *gin.Context) {

	roomIDstr := con.PostForm("roomID")

	if len(roomIDstr) > 0 {

		roomID, err := strconv.Atoi(roomIDstr)

		if err != nil {
			fmt.Println("applyEnterRoom atoi", err, roomID)
		}

		//room, err := models.GetRoomInfo(int64(roomID))

	}

	//tqgin.ResultOk(con)
}
