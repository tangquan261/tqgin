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

	needGold := int64(giftInfo.CashNum) * int64(len(gG.Players)) * int64(gG.NCount)

	if UserInfo.Gold < needGold {
		tqgin.ResultFail(c, "钻石不足")
		return
	}

	err = models.ModifyGoldUser(playerGUID, -needGold)
	if err != nil {
		tqgin.ResultFail(c, "钻石不足")
		return
	}
	//添加送礼记录
	models.AddGiveGiftLog(gG.GiftID, playerGUID, gG.RoomID, gG.Players, gG.NCount)
	//增加自己的财富值
	models.ModifyRichUser(playerGUID, needGold*10)

	//增加给送礼人的魅力值
	for i := 0; i < len(gG.Players); i++ {
		models.ModifyCharmUser(gG.Players[i], needGold*10)
	}

	if gG.RoomID != 0 {
		//增加自己在该房间财富
		var roomRank models.RoomRankInfo
		roomRank.PlayerID = playerGUID
		roomRank.Rich = needGold * 10
		roomRank.RoomID = gG.RoomID

		models.RoomRankinfoSave(roomRank)

		//增加别人在该魅力
		for i := 0; i < len(gG.Players); i++ {
			var roomRank models.RoomRankInfo
			roomRank.PlayerID = gG.Players[i]
			roomRank.Charm = needGold * 10
			roomRank.RoomID = gG.RoomID
			models.RoomRankinfoSave(roomRank)
		}
	}

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
