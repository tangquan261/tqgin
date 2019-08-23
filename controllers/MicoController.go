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

	temp.POST("mico_apply_up", this.applyMicoUp)        //上麦
	temp.POST("mico_apply_down", this.applyMicoDown)    //下麦
	temp.POST("mico_hold_up", this.holdUpMic)           //抱上麦
	temp.POST("mic_hold_down", this.holdDownMic)        //抱下麦
	temp.POST("mic_index_modify", this.modifyMicStatus) // 麦序状态的关闭与开启
	temp.POST("mic_heattime", this.heatTime)            //心跳
}

type MicJson struct {
	RoomID   int64 `json:"roomid"`
	MicIndex int16 `json:"micindex"`
}

//主动自己上麦
func (r *MicoController) applyMicoUp(c *gin.Context) {

	playerGUID := r.GetPlayerGUID(c)

	var micData MicJson
	err := c.ShouldBindJSON(&micData)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if micData.RoomID <= 0 {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if micData.MicIndex < 0 || micData.MicIndex >= 8 {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	roominfo := models.GetRoomById(micData.RoomID)

	if roominfo.RoomOrderMic == 1 {
		//自由上麦
		//直接加入麦序
		models.MicAdd(micData.RoomID, playerGUID, micData.MicIndex)
	} else {
		//排序上麦
		models.MicAdd(micData.RoomID, playerGUID, micData.MicIndex)
		//models.MicQueueAdd(micData.RoomID, playerGUID)
	}

	mics := models.MicGetAllIndex(micData.RoomID)
	micsQueue := models.MicQueueInfo(micData.RoomID)

	tqgin.ResultOkMsg(c, gin.H{"mics": mics, "micqueue": micsQueue}, "成功")
}

//主动自己下麦，或者取消排麦
func (r *MicoController) applyMicoDown(c *gin.Context) {

	playerGUID := r.GetPlayerGUID(c)

	var micData MicJson
	err := c.ShouldBindJSON(&micData)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}
	//从麦序移除
	models.MicDelByPlayerID(playerGUID)
	//从排麦移除
	models.MicQueueDel(playerGUID)

	mics := models.MicGetAllIndex(micData.RoomID)
	micsQueue := models.MicQueueInfo(micData.RoomID)

	tqgin.ResultOkMsg(c, gin.H{"mics": mics, "micqueue": micsQueue}, "成功")
}

type MicHoldJson struct {
	RoomID      int64 `json:"roomid"`
	TarPlayerID int64 `json:"tarPlayerID"`
	MicIndex    int16 `json:"micindex"`
}

//抱上麦
func (r *MicoController) holdUpMic(c *gin.Context) {
	playerGUID := r.GetPlayerGUID(c)

	var micData MicHoldJson
	err := c.ShouldBindJSON(&micData)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if !hasJurisdiction(micData.RoomID, playerGUID) {
		tqgin.ResultFail(c, "没有权限")
		return
	}

	models.MicAdd(micData.RoomID, playerGUID, micData.MicIndex)

	mics := models.MicGetAllIndex(micData.RoomID)
	micsQueue := models.MicQueueInfo(micData.RoomID)

	tqgin.ResultOkMsg(c, gin.H{"mics": mics, "micqueue": micsQueue}, "成功")
}

//抱下麦
func (r *MicoController) holdDownMic(c *gin.Context) {
	playerGUID := r.GetPlayerGUID(c)

	var micData MicHoldJson
	err := c.ShouldBindJSON(&micData)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if !hasJurisdiction(micData.RoomID, playerGUID) {
		tqgin.ResultFail(c, "没有权限")
		return
	}

	err = models.MicDelByPlayerID(micData.TarPlayerID)
	if err != nil {
		tqgin.ResultFail(c, "失败")
		return
	}

	mics := models.MicGetAllIndex(micData.RoomID)
	micsQueue := models.MicQueueInfo(micData.RoomID)

	tqgin.ResultOkMsg(c, gin.H{"mics": mics, "micqueue": micsQueue}, "成功")

}

type MicStatusJson struct {
	RoomID    int64 `json:"roomid"`
	MicIndex  int16 `json:"micindex"`
	MicStatus int16 `json:"micstatus"`
}

//麦序的关闭开启
func (r *MicoController) modifyMicStatus(c *gin.Context) {
	playerGUID := r.GetPlayerGUID(c)

	var micData MicStatusJson
	err := c.ShouldBindJSON(&micData)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if !hasJurisdiction(micData.RoomID, playerGUID) {
		tqgin.ResultFail(c, "没有权限")
		return
	}

	err = models.MicUpdateStatus(micData.RoomID, micData.MicIndex, micData.MicStatus)

	if err != nil {
		tqgin.ResultFail(c, "操作失败")
		return
	}

	mics := models.MicGetAllIndex(micData.RoomID)
	micsQueue := models.MicQueueInfo(micData.RoomID)

	tqgin.ResultOkMsg(c, gin.H{"mics": mics, "micqueue": micsQueue}, "成功")
}

//麦序心跳30秒心跳一次
func (r *MicoController) heatTime(c *gin.Context) {

	playerID := r.GetPlayerGUID(c)
	type MicHeatTime struct {
		RoomID   int64 `json:"roomid"`
		MicIndex int16 `json:"micindex"`
	}

	var heattime MicHeatTime

	err := c.ShouldBindJSON(&heattime)
	if err != nil {
		tqgin.ResultFail(c, "错误")
		return
	}

	mics := models.MicUpdate(heattime.RoomID, playerID, heattime.MicIndex)

	tqgin.ResultOkMsg(c, mics, "成功")
}

//判断是否有权限操作
func hasJurisdiction(roomid, playerID int64) bool {

	if roomid == playerID {
		return true
	}

	_, hasPower := models.GetPowerRoom(playerID, roomid)
	if hasPower {
		return true
	}
	return false
}
