/*
	朋友圈的feeds
*/

package controllers

import (
	"strings"
	"tqgin/common"
	"tqgin/models"

	"tqgin/pkg/define"
	"tqgin/pkg/util"

	"github.com/gin-gonic/gin"
)

type CycleController struct {
	tqgin.Controller
}

func (this *CycleController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/cycle")
	temp.POST("add_feed", this.addFeed)
	temp.POST("del_feed", this.delFeed)
	temp.POST("get_feed", this.getFeed)
	temp.POST("get_feed_follow", this.getFeedFollow)
	//temp.POST("get_feed_fans", this.getFeedFans)
}

//GiftID  int64   `json:"giftid"`
//RoomID  int64   `json:"roomid"`
//NCount  int32   `json:"count"`
//Players []int64 `json:"players"`

type FeedParam struct {
	Cid       string           `json:"cid"`       //创做唯一标识
	FType     define.CycleType `json:"ftype"`     //1,2,3 普通,声音,视频
	SoundRUL  string           `json:"soundurl"`  //声音地址
	PhotoURLs []string         `json:"photourl"`  //图片地址列表
	Content   string           `json:"content"`   //文本
	Ats       []string         `json:"at"`        //at的人列表
	LocX      int64            `json:"locx"`      //x位置
	LocY      int64            `json:"locy"`      //y位置
	LocString string           `json:"locstring"` //位置名称
}

//添加朋友圈
func (r *CycleController) addFeed(c *gin.Context) {
	playerID := r.GetPlayerGUID(c)

	var feed FeedParam
	err := c.ShouldBindJSON(&feed)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	var dbFeed models.CycleModel
	dbFeed.PlayerID = playerID

	dbFeed.Cid = feed.Cid
	dbFeed.FType = feed.FType
	dbFeed.SoundRUL = feed.SoundRUL
	dbFeed.PhotoURLs = strings.Join(feed.PhotoURLs, "@@@")
	dbFeed.Content = util.UnicodeEmojiCode(feed.Content)
	dbFeed.Ats = strings.Join(feed.Ats, "@@@")
	dbFeed.LocX = feed.LocX
	dbFeed.LocY = feed.LocY
	dbFeed.LocString = feed.LocString

	dbFeed.SnowID = util.SnowFlakeUUID()

	err = models.CycleAdd(dbFeed)

	if err != nil {
		tqgin.ResultFail(c, err.Error())
	} else {
		tqgin.ResultOkMsg(c, nil, "成功")
	}
}

type FeedGetParam struct {
	SnowID int64 `json:"snowid"` //请求文章索引，第一次传入0
	FType  int32 `json:"ftype"`  //1,2 声音，普通
}

//删除动态
func (r *CycleController) delFeed(c *gin.Context) {

	var getparam FeedGetParam
	err := c.ShouldBindJSON(&getparam)
	if err != nil || getparam.SnowID <= 0 {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	//删除帖子
	models.CycleDel(getparam.SnowID)
	//删除评论
	models.CycleDelCommetByCycleuuid(getparam.SnowID)
	//删除点赞
	models.CycleDelLikesByCycleuuid(getparam.SnowID)

	tqgin.ResultOk(c, "成功")
}

//获取广场动态，声音
func (r *CycleController) getFeed(c *gin.Context) {

	var getparam FeedGetParam
	err := c.ShouldBindJSON(&getparam)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	if getparam.FType == 2 {
		//声音
		ret := models.CycleGetSound(getparam.SnowID)
		tqgin.ResultOk(c, ret)
	} else if getparam.FType == 3 {
		//视频
		ret := models.CycleGetAudio(getparam.SnowID)
		tqgin.ResultOk(c, ret)
	} else {
		//all
		ret := models.CycleGetFeeds(getparam.SnowID)
		tqgin.ResultOk(c, ret)
	}
}

//我关注的人的动态
func (r *CycleController) getFeedFollow(c *gin.Context) {

	playerID := r.GetPlayerGUID(c)

	var getparam FeedGetParam
	err := c.ShouldBindJSON(&getparam)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	ret := models.CycleGetFeedsFollow(playerID, getparam.SnowID)
	tqgin.ResultOk(c, ret)
}

//我的粉丝的动态
func (r *CycleController) getFeedFans(c *gin.Context) {

	playerID := r.GetPlayerGUID(c)
	var getparam FeedGetParam
	err := c.ShouldBindJSON(&getparam)
	if err != nil {
		tqgin.ResultFail(c, "参数错误")
		return
	}

	ret := models.CycleGetFeedsFans(playerID, getparam.SnowID)
	tqgin.ResultOk(c, ret)
}
