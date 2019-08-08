/*
	好友，关注管理
*/

package controllers

import (
	"log"
	//"log"
	"strconv"

	// "fmt"
	// "strconv"
	"tqgin/common"
	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type AttentionController struct {
	tqgin.Controller
}

func (this *AttentionController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/attention")

	temp.POST("attention_list", this.attentionList)
	temp.POST("add_friend", this.addFriend)
	temp.POST("remove_friend", this.removeFriend)
	temp.POST("add_black", this.addBlack)
	temp.POST("remove_black", this.removeBlack)
	temp.POST("add_attention", this.addAttention)
	temp.POST("remove_attention", this.removeAttention)

}

type playerJson struct {
	PlayerID int64 `json:"playerID"`
}

type PlayerRet struct {
	PlayerID int64 `json:"playerID"`
	Name     string
}

//关系列表
func (r *AttentionController) attentionList(c *gin.Context) {

	myID, _ := c.Cookie("playerid")
	nPlayerID, _ := strconv.ParseInt(myID, 10, 64)

	uids := models.GetFriends(nPlayerID)

	var friendList []PlayerRet
	for _, friend := range uids {
		userinfo := models.GetUser(friend.PlayerID)
		if userinfo.PlayerID != 0 {
			friendList = append(friendList, PlayerRet{PlayerID: userinfo.PlayerID, Name: userinfo.PlayerName})
		}
	}

	uidsBlack := models.GetBlacks(nPlayerID)

	var blackList []PlayerRet
	for _, friend := range uidsBlack {
		userinfo := models.GetUser(friend.PlayerID)
		if userinfo.PlayerID != 0 {
			blackList = append(blackList, PlayerRet{PlayerID: userinfo.PlayerID, Name: userinfo.PlayerName})
		}
	}

	uidsAttention := models.GetAttentions(nPlayerID)

	var attentionList []PlayerRet
	for _, friend := range uidsAttention {
		userinfo := models.GetUser(friend.PlayerID)
		if userinfo.PlayerID != 0 {
			attentionList = append(attentionList, PlayerRet{PlayerID: userinfo.PlayerID, Name: userinfo.PlayerName})
		}
	}

	tqgin.ResultOkMsg(c, gin.H{"friends": friendList, "black": blackList, "attention": attentionList}, "成功")
}

//添加好友
func (r *AttentionController) addFriend(c *gin.Context) {
	myID, _ := c.Cookie("playerid")
	nPlayerID, _ := strconv.ParseInt(myID, 10, 64)

	var friend playerJson

	err := c.ShouldBindJSON(&friend)

	if err != nil {
		tqgin.ResultFail(c, "解析json错误")
		return
	}

	err = models.AddFriend(nPlayerID, friend.PlayerID)
	if err != nil {
		tqgin.ResultFail(c, "数据错误")
		return
	}
	tqgin.ResultOkMsg(c, friend, "添加成功")
}

//删除好友
func (r *AttentionController) removeFriend(c *gin.Context) {
	myID, _ := c.Cookie("playerid")
	nPlayerID, _ := strconv.ParseInt(myID, 10, 64)

	var friend playerJson

	err := c.ShouldBindJSON(&friend)

	if err != nil {
		tqgin.ResultFail(c, "解析json错误")
		return
	}

	err = models.RemoveFriend(nPlayerID, friend.PlayerID)
	if err != nil {
		log.Println(err)
		tqgin.ResultFail(c, "数据错误")
		return
	}
	tqgin.ResultOkMsg(c, friend, "删除成功")

}

//添加黑名单
func (r *AttentionController) addBlack(c *gin.Context) {
	myID, _ := c.Cookie("playerid")
	nPlayerID, _ := strconv.ParseInt(myID, 10, 64)

	var friend playerJson

	err := c.ShouldBindJSON(&friend)

	if err != nil {
		tqgin.ResultFail(c, "解析json错误")
		return
	}

	err = models.AddBlack(nPlayerID, friend.PlayerID)
	if err != nil {
		tqgin.ResultFail(c, "数据错误")
		return
	}
	tqgin.ResultOkMsg(c, friend, "添加成功")
}

//移除黑名单
func (r *AttentionController) removeBlack(c *gin.Context) {
	myID, _ := c.Cookie("playerid")
	nPlayerID, _ := strconv.ParseInt(myID, 10, 64)

	var friend playerJson

	err := c.ShouldBindJSON(&friend)

	if err != nil {
		tqgin.ResultFail(c, "解析json错误")
		return
	}

	err = models.RemoveBlack(nPlayerID, friend.PlayerID)
	if err != nil {
		log.Println(err)
		tqgin.ResultFail(c, "数据错误")
		return
	}
	tqgin.ResultOkMsg(c, friend, "删除成功")
}

//添加关注房间
func (r *AttentionController) addAttention(c *gin.Context) {
	myID, _ := c.Cookie("playerid")
	nPlayerID, _ := strconv.ParseInt(myID, 10, 64)

	var friend playerJson

	err := c.ShouldBindJSON(&friend)

	if err != nil {
		tqgin.ResultFail(c, "解析json错误")
		return
	}

	err = models.AddAttention(nPlayerID, friend.PlayerID)
	if err != nil {
		tqgin.ResultFail(c, "数据错误")
		return
	}
	tqgin.ResultOkMsg(c, friend, "添加成功")
}

//取消关注
func (r *AttentionController) removeAttention(c *gin.Context) {
	myID, _ := c.Cookie("playerid")
	nPlayerID, _ := strconv.ParseInt(myID, 10, 64)

	var friend playerJson

	err := c.ShouldBindJSON(&friend)

	if err != nil {
		tqgin.ResultFail(c, "解析json错误")
		return
	}

	err = models.RemoveAttention(nPlayerID, friend.PlayerID)
	if err != nil {
		log.Println(err)
		tqgin.ResultFail(c, "数据错误")
		return
	}
	tqgin.ResultOkMsg(c, friend, "删除成功")
}
