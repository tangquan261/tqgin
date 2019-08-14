/*
	朋友圈的feeds,点赞，评论
*/

package controllers

import (
	"tqgin/common"
	"tqgin/models"

	"tqgin/pkg/util"

	"github.com/gin-gonic/gin"
)

type CycleCommetController struct {
	tqgin.Controller
}

func (this *CycleCommetController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/cycle")
	temp.POST("add_commet", this.addCommet)
	temp.POST("del_commet", this.delCommet)
	temp.POST("add_likes", this.addLikes)
}

type CommetParam struct {
	Uuid    string `json:"uuid"`    //文章id
	FromID  int64  `json:"fromID"`  //前一个评论id，1级评论是0
	Content string `json:"content"` //评论内容
	Likes   int32  `json:"likes"`   //1点赞0 取消点赞
}

//添加评论
func (r *CycleCommetController) addCommet(c *gin.Context) {
	playerID := r.GetPlayerGUID(c)

	var feed CommetParam
	err := c.ShouldBindJSON(&feed)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if len(feed.Uuid) <= 0 {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	cycle := models.CycleGetModel(feed.Uuid)
	if cycle == nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	var commet models.CycleCommet
	commet.Uuid = feed.Uuid
	commet.PlayerID = playerID
	commet.FromID = feed.FromID
	commet.Conent = util.UnicodeEmojiCode(feed.Content)

	err = models.CycleAddCommet(commet)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
	} else {
		tqgin.ResultOk(c, nil)
	}
}

//删除评论
func (r *CycleCommetController) delCommet(c *gin.Context) {
	//playerID := r.GetPlayerGUID(c)

	var feed CommetParam
	err := c.ShouldBindJSON(&feed)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	err = models.CycleDelCommet(feed.FromID)

	if err != nil {
		tqgin.ResultFail(c, "参数错误")
	} else {
		tqgin.ResultOk(c, nil)
	}
}

//点赞，取消点赞
func (r *CycleCommetController) addLikes(c *gin.Context) {
	playerID := r.GetPlayerGUID(c)

	var feed CommetParam
	err := c.ShouldBindJSON(&feed)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if feed.Likes == 1 {
		var data models.CycleLike
		data.PlayerID = playerID
		data.UID = feed.FromID
		data.Uuid = feed.Uuid

		models.CycleAddLikeCommet(data)
		tqgin.ResultOk(c, nil)
	} else {
		models.CycleDelLikeCommet(feed.FromID)
		tqgin.ResultOk(c, nil)
	}

}
