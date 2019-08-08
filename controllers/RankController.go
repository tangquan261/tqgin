/*

	排行榜相关
*/

package controllers

import (
	"tqgin/common"
	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type RankController struct {
	tqgin.Controller
}

func (this *RankController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/rank")
	temp.POST("rank_info", this.rankinfo)
	temp.POST("room_rank_info", this.roomRankInfo)

}

type rankApply struct {
	Idx     int64              `json:"idx"`
	Type    models.RankType    `json:"type"`    //0-4,表示富豪，魅力，房间，名人，声望
	SubType models.RankSubType `json:"subtype"` //0-2,表示日，周，月
}

//排行榜
func (r *RankController) rankinfo(c *gin.Context) {
	var rank rankApply
	err := c.ShouldBindJSON(&rank)
	if err != nil {
		tqgin.ResultFail(c, "错误")
		return
	}

	ret := models.RankInfoBy(rank.Type, rank.SubType)
	tqgin.ResultOkMsg(c, ret, "成功")
}

type roomRankApply struct {
	RoomID   int64               `json:"roomid"`
	RoomType models.RoomRankType `json:"roomtype"` //0,1表示贡献周榜，魅力周榜
}

//房间内贡献，魅力排行榜
func (r *RankController) roomRankInfo(c *gin.Context) {
	var rank roomRankApply
	err := c.ShouldBindJSON(&rank)
	if err != nil {
		tqgin.ResultFail(c, "错误")
		return
	}

	models.RoomRankInfoBy(rank.RoomID, rank.RoomType)
	tqgin.ResultOkMsg(c, rank, "成功")
}
