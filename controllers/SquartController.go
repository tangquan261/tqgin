/*
	房间大厅列表
*/
package controllers

import (
	"tqgin/common"
	"tqgin/models"

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

	//	var retTags login.TagsInfo

	//	for _, tag := range tags {
	//		var onetag login.TagInfo
	//		onetag.ID = tag.ID
	//		onetag.TagName = tag.TagName
	//		retTags.TagInfo = append(retTags.TagInfo, &onetag)
	//	}

	tqgin.ResultOkMsg(con, tags, "成功")
}

func (s *SquartController) applyRoomList(con *gin.Context) {

	type RoomtagParam struct {
		RoomTag string `json:"roomtag"` //房间Tag类型
	}

	var roomTag RoomtagParam

	err := con.ShouldBindJSON(&roomTag)
	if err != nil {
		tqgin.ResultFail(con, "参数错误")
		return
	}

	var data []models.RoomInfo

	if len(roomTag.RoomTag) > 0 {
		data = models.GetHotRoomsByTag(roomTag.RoomTag)
	} else {
		data = models.GetHotAllRooms()
	}

	tqgin.ResultOk(con, data)
}

func (s *SquartController) getBanners(con *gin.Context) {

	banners := models.GetBanners()

	tqgin.ResultOk(con, banners)
}
