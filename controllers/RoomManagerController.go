/*
	房间管理
	创建，移除，修改操作
*/

package controllers

import (
	"fmt"
	"strconv"

	// "strconv"
	"tqgin/common"
	// "tqgin/models"

	"github.com/gin-gonic/gin"
)

type RoomManagerController struct {
	tqgin.Controller
}

func (this *RoomManagerController) RegisterRouter(router *gin.Engine) {
	temp := router.Group("/roomManager")
	temp.POST("openRoom", this.OpenRoom)
	temp.POST("closeRoom", this.closeRoom)
	temp.POST("changeRoomName", this.ChangeRoomName)
}

//开启房间
func (r *RoomManagerController) OpenRoom(c *gin.Context) {

	PlayerID, _ := strconv.Atoi(c.PostForm("playerID"))
	name := c.PostForm("roomName")
	tags := c.PostForm("tags")
	disc := c.PostForm("disc")
	fmt.Println(PlayerID, name, tags, disc)
}

//关闭房间
func (r *RoomManagerController) closeRoom(c *gin.Context) {

}

//关闭房间
func (r *RoomManagerController) ChangeRoomName(c *gin.Context) {

}
