package controllers

import (
	// "fmt"
	// "strconv"
	"tqgin/common"
	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type MicoController struct {
	tqgin.Controller
}

func (this *MicoController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/mico")
	temp.POST("mico_add", this.micoAdd)
	temp.POST("mico_del", this.micoDel)

}

type MicJson struct {
	TarPlayerID int64 `json:"playerid"`
	RoomID      int64 `json:"roomid"`
	MicIndex    int16 `json:"micindex"`
}

//加入mico
func (r *MicoController) micoAdd(c *gin.Context) {

	playerGUID := r.GetPlayerGUID(c)

	var micData MicJson
	err := c.ShouldBindJSON(&micData)
	if err != nil {
		tqgin.ResultFail(c, "参数解析错误")
		return
	}

	if micData.TarPlayerID <= 0 || micData.RoomID <= 0 {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if micData.MicIndex < 0 || micData.MicIndex >= 8 {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	power, bfind := models.GetPowerRoom(playerGUID, micData.RoomID)
	if !bfind || power.RoomPower <= 0 {
		tqgin.ResultFail(c, "没有权限")
		return
	}

	models.MicAdd(micData.RoomID, micData.TarPlayerID, micData.MicIndex)

	mics := models.MicGetAllIndex(micData.RoomID)

	tqgin.ResultOkMsg(c, mics, "添加成功")
}

//移除mico
func (r *MicoController) micoDel(c *gin.Context) {

	//playerGUID := r.GetPlayerGUID(c)

	var micData MicJson
	c.ShouldBindJSON(&micData)

	models.MicDelByPlayerID(micData.TarPlayerID)

	mics := models.MicGetAllIndex(micData.RoomID)
	tqgin.ResultOkMsg(c, mics, "移除成功")
}
