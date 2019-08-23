/*
	声网
*/
package controllers

import (
	"strconv"
	"tqgin/common"
	"tqgin/pkg/Agora"

	"github.com/gin-gonic/gin"
)

type AgoraController struct {
	tqgin.Controller
}

func (this *AgoraController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/agroa")
	temp.POST("get_tocken", this.getTocken) //获取声网tocken
}

func (a *AgoraController) getTocken(c *gin.Context) {

	playerID := a.GetPlayerGUID(c)

	type TokenParam struct {
		RoomID int64 `json:"roomid"`
	}
	var tokenparam TokenParam

	err := c.ShouldBindJSON(&tokenparam)
	if err != nil || tokenparam.RoomID < 0 {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	channel := "channel" + strconv.FormatInt(tokenparam.RoomID, 10)

	token, err := tokenbuilder.BuildTokenWithUID("1f836f0e094446d2858f156ca366313d", "08e1620922bf40ff9ac81517f4219f51", channel, uint32(playerID), tokenbuilder.RolePublisher, 0)
	if err != nil {
		tqgin.ResultFail(c, "获取错误")
	} else {
		tqgin.ResultOkMsg(c, token, "获取成功")
	}
}
