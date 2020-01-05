/*
	关注，粉丝好友管理
*/

package controllers

import (
	"tqgin/common"
	"tqgin/models"

	"tqgin/pkg/errorcode"

	"github.com/gin-gonic/gin"
)

type RelationController struct {
	tqgin.Controller
}

func (this *RelationController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/relation")
	temp.POST("add_attention", this.addAttention)   //关注
	temp.POST("del_attention", this.delAttention)   //取消关注
	temp.POST("get_friends", this.getfriends)       //获取好友列表
	temp.POST("get_attentions", this.getAttentions) //获取关注列表
	//temp.POST("get_fans", this.getFans)           //获取粉丝列表
	temp.POST("add_black", this.addBlack)    //添加黑名单
	temp.POST("del_black", this.removeBlack) //移除黑名单
	temp.POST("get_blacks", this.getBlacks)  //获取黑名单
}

type RelationParam struct {
	PlayerID int64 `json:"playerid"`
}

//关注
func (r *RelationController) addAttention(c *gin.Context) {

	myPlayerID := r.GetPlayerGUID(c)

	var tarparam RelationParam

	err := c.ShouldBindJSON(&tarparam)

	if err != nil {

		tqgin.ResultFail(c, errorcode.GetMsg(errorcode.ERROR_INVALID_PARAMS))
		return
	}

	err = models.RelationAddFollow(myPlayerID, tarparam.PlayerID)
	if err == nil {
		isFriend := models.RelationIsFans(myPlayerID, tarparam.PlayerID)
		tqgin.ResultOkMsg(c, gin.H{"isFriend": isFriend}, "成功")
	} else {
		tqgin.ResultFail(c, "失败")
	}
}

//取消关注
func (r *RelationController) delAttention(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	var tarparam RelationParam

	err := c.ShouldBindJSON(&tarparam)

	if err != nil {

		tqgin.ResultFail(c, errorcode.GetMsg(errorcode.ERROR_INVALID_PARAMS))
		return
	}

	err = models.RelationDelFollow(myPlayerID, tarparam.PlayerID)

	if err != nil {
		tqgin.ResultFail(c, err.Error())
	} else {
		tqgin.ResultOk(c, err)
	}
}

//获取关注列表
func (r *RelationController) getAttentions(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	user := models.GetFollow(myPlayerID)
	var follows []int64
	for _, obj := range user {
		follows = append(follows, obj.TarplayerID)
	}
	tqgin.ResultOk(c, gin.H{"attentions": follows})
}

//获取粉丝列表
func (r *RelationController) getFans(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	user := models.GetFans(myPlayerID)
	var fans []int64
	for _, obj := range user {
		fans = append(fans, obj.PlayerID)
	}
	tqgin.ResultOk(c, fans)
}

//获取好友列表
func (r *RelationController) getfriends(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	user := models.GetFirend(myPlayerID)
	var friends []int64
	for _, obj := range user {
		friends = append(friends, obj.PlayerID)
	}
	tqgin.ResultOk(c, gin.H{"friends": friends})
}

//添加黑名单
func (r *RelationController) addBlack(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	var tarparam RelationParam

	err := c.ShouldBindJSON(&tarparam)

	if err != nil {

		tqgin.ResultFail(c, errorcode.GetMsg(errorcode.ERROR_INVALID_PARAMS))
		return
	}

	err = models.AddBlack(myPlayerID, tarparam.PlayerID)

	if err != nil {
		tqgin.ResultFail(c, err.Error())
	} else {
		tqgin.ResultOk(c, nil)
	}
}

//添加黑名单
func (r *RelationController) removeBlack(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	var tarparam RelationParam

	err := c.ShouldBindJSON(&tarparam)

	if err != nil {

		tqgin.ResultFail(c, errorcode.GetMsg(errorcode.ERROR_INVALID_PARAMS))
		return
	}

	err = models.RemoveBlack(myPlayerID, tarparam.PlayerID)

	if err != nil {
		tqgin.ResultFail(c, err.Error())
	} else {
		tqgin.ResultOk(c, nil)
	}
}

//获取黑名单
func (r *RelationController) getBlacks(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	rets := models.GetBlacks(myPlayerID)

	var blacks []int64
	for _, obj := range rets {
		blacks = append(blacks, obj.BlackID)
	}
	tqgin.ResultOk(c, gin.H{"blacks": blacks})
}
