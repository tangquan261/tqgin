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
	temp.POST("follow", this.follow)         //关注
	temp.POST("del_follow", this.delfollow)  //取消关注
	temp.POST("get_friend", this.getfriend)  //获取好友列表
	temp.POST("get_follow", this.getfollows) //获取关注列表
	temp.POST("get_fans", this.getFans)      //获取粉丝列表
	temp.POST("add_black", this.addBlack)    //添加黑名单
	temp.POST("del_black", this.removeBlack) //移除黑名单
	temp.POST("get_blacks", this.getBlacks)  //获取黑名单
}

type RelationParam struct {
	PlayerID int64 `json:"playerid"`
}

//关注
func (r *RelationController) follow(c *gin.Context) {

	myPlayerID := r.GetPlayerGUID(c)

	var tarparam RelationParam

	err := c.ShouldBindJSON(&tarparam)

	if err != nil {

		tqgin.ResultFail(c, errorcode.GetMsg(errorcode.ERROR_INVALID_PARAMS))
		return
	}

	err = models.RelationAddFollow(myPlayerID, tarparam.PlayerID)
	if err != nil {
		tqgin.ResultFail(c, err.Error())
	} else {
		tqgin.ResultOk(c, err)
	}
}

//取消关注
func (r *RelationController) delfollow(c *gin.Context) {
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
func (r *RelationController) getfollows(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	user := models.GetFollow(myPlayerID)
	tqgin.ResultOk(c, user)
}

//获取粉丝列表
func (r *RelationController) getFans(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	user := models.GetFans(myPlayerID)
	tqgin.ResultOk(c, user)
}

//获取好友列表
func (r *RelationController) getfriend(c *gin.Context) {
	myPlayerID := r.GetPlayerGUID(c)

	user := models.GetFirend(myPlayerID)
	tqgin.ResultOk(c, user)
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

	tqgin.ResultOk(c, rets)
}
