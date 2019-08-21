package controllers

import (
	"strconv"

	// "fmt"
	// "strconv"
	"tqgin/common"
	//"tqgin/models"
	"tqgin/pkg/Agora"

	"github.com/gin-gonic/gin"
)

type AgoraController struct {
	tqgin.Controller
}

func (this *AgoraController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/agroa")
	temp.POST("get_tocken", this.getTocken)
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
	//account := "account" + strconv.FormatInt(playerID, 10)

	//token, err := rtctokenbuilder.BuildTokenWithUserAccount("1f836f0e094446d2858f156ca366313d", "08e1620922bf40ff9ac81517f4219f51", channel, account, rtctokenbuilder.RolePublisher, 0)

	token, err := rtctokenbuilder.BuildTokenWithUID("1f836f0e094446d2858f156ca366313d", "08e1620922bf40ff9ac81517f4219f51", channel, uint32(playerID), rtctokenbuilder.RolePublisher, 0)
	if err != nil {
		tqgin.ResultFail(c, "获取错误")
	} else {
		tqgin.ResultOkMsg(c, token, "获取成功")
	}
}
