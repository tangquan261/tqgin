/*
	房间大厅列表
*/
package controllers

import (
	"tqgin/common"
	"tqgin/models"
	"tqgin/proto"

	"github.com/gin-gonic/gin"
)

type SquartController struct {
	tqgin.Controller
}

func (this *SquartController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/squart")
	temp.GET("applyTagsList", this.applyTagsList)
	temp.POST("applyRoomList", this.applyRoomList)
	temp.POST("getBanners", this.getBanners)
}

func (s *SquartController) applyTagsList(con *gin.Context) {

	tags := models.GetTagList()

	var retTags login.TagsInfo

	for _, tag := range tags {
		var onetag login.TagInfo
		onetag.ID = tag.ID
		onetag.TagName = tag.TagName
		retTags.TagInfo = append(retTags.TagInfo, &onetag)
	}

	tqgin.ResultOkMsg(con, tags, "成功")
}

func (s *SquartController) applyRoomList(con *gin.Context) {

	TagName := con.PostForm("tagName")

	var data []models.RoomInfo

	if len(TagName) > 0 {
		data = models.GetHotRoomsByTag(TagName)
	} else {
		data = models.GetHotAllRooms()
	}

	tqgin.ResultOk(con, data)
}

func (s *SquartController) getBanners(con *gin.Context) {

	banners := models.GetBanners()

	tqgin.ResultOk(con, banners)
}
