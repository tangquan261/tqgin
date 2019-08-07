/*
	用户货币管理
*/

package controllers

import (
	"tqgin/common"
	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type UserInfoMoneyController struct {
	tqgin.Controller
}

func (this *UserInfoMoneyController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/money")
	temp.POST("add_diamond", this.add_diamond)
	temp.POST("give_gift", this.giveGift)
	temp.POST("all_gift", this.allGift)
}

type MoneyJsonAdd struct {
	Diamond int64 `json:"diamond"`
	Gold    int64 `json:"gold"`
	Cash    int64 `json:"cash"`
}

func (r *UserInfoMoneyController) add_diamond(c *gin.Context) {
	playerGUID := r.GetPlayerGUID(c)
	if playerGUID <= 0 {
		tqgin.ResultFail(c, "not find")
		return
	}

	var money MoneyJsonAdd

	err := c.ShouldBindJSON(&money)
	if err != nil {
		tqgin.ResultFail(c, "解析错误")
		return
	}

	err = AddPlayerDiamond(playerGUID, money.Diamond)
	if err != nil {
		tqgin.ResultFail(c, "添加错误")
		return
	}
	tqgin.ResultOkMsg(c, money, "添加成功")
}

type giveGift struct {
	GiftID  int64   `json:"giftid"`
	RoomID  int64   `json:"roomid"`
	NCount  int32   `json:"count"`
	Players []int64 `json:"players"`
}

func (r *UserInfoMoneyController) giveGift(c *gin.Context) {
	playerGUID := r.GetPlayerGUID(c)
	if playerGUID <= 0 {
		tqgin.ResultFail(c, "not find")
		return
	}
	var gG giveGift

	err := c.ShouldBindJSON(&gG)

	if err != nil {
		tqgin.ResultFail(c, "解析错误")
		return
	}

	giftInfo := models.GetGiftByID(gG.GiftID)
	if giftInfo.GiftID == 0 {
		tqgin.ResultFail(c, "礼物不存在")
		return
	}
	UserInfo := models.GetUser(playerGUID)

	needDiamond := int64(giftInfo.CashNum) * int64(len(gG.Players)) * int64(gG.NCount)

	if UserInfo.Diamond < needDiamond {
		tqgin.ResultFail(c, "钻石不足")
		return
	}

	err = models.ModifyDinamondUser(playerGUID, -needDiamond)
	if err != nil {
		tqgin.ResultFail(c, "钻石不足")
		return
	}

	models.AddGiveGiftLog(gG.GiftID, playerGUID, gG.RoomID, gG.Players, gG.NCount)

	tqgin.ResultOkMsg(c, nil, "送礼成功")
}

func (r *UserInfoMoneyController) allGift(c *gin.Context) {

	gifts := models.GetAllGift()

	tqgin.ResultOk(c, gifts)
}

func AddPlayerDiamond(playerGUID, diamond int64) error {

	diamond = 1
	for i := 0; i < 1; i++ {
		go func() {

			models.ModifyDinamondUser(playerGUID, -1)
		}()
	}
	return nil
}
