/*
	房间列表
*/

package controllers

import (
	"fmt"
	"tqgin/common"
	"tqgin/models"
	"tqgin/proto"

	"github.com/gin-gonic/gin"
)

type SquartController struct {
	tqgin.Controller
}

func (this *SquartController) RegisterRouter(router *gin.Engine) {
	temp := router.Group("/squart")
	temp.GET("applyTagsList", this.applyTagsList)
	temp.POST("applyRoomList", this.applyRoomList)
	temp.POST("getBanners", this.getBanners)
}

func (s *SquartController) applyTagsList(con *gin.Context) {

	var status int

	tags := models.GetTagList()

	var retTags login.TagsInfo

	for _, tag := range tags {
		var onetag login.TagInfo
		onetag.ID = tag.ID
		onetag.TagName = tag.TagName
		retTags.TagInfo = append(retTags.TagInfo, &onetag)
	}

	fmt.Println(status, tags, "\n", retTags)
	tqgin.ResultOkMsg(con, &retTags, "成功")
}

func (s *SquartController) applyRoomList(con *gin.Context) {

	TagName := con.PostForm("tagName")

	var status int
	var data []models.HotRoomInfo

	var retData login.HotRooms

	if len(TagName) > 0 {
		status = 0
		data = models.GetHotRoomsByTag(TagName)

	} else {
		status = 0
		data = models.GetHotAllRooms()
	}

	for _, room := range data {
		var oneRoom login.HotRoomInfo

		oneRoom.RoomID = room.RoomID
		oneRoom.RoomTagName = room.RoomTagName
		oneRoom.RoomHot = room.RoomHot

		retData.HotRoomInfo = append(retData.HotRoomInfo, &oneRoom)
	}

	fmt.Println(status, data, "\n", retData)
	tqgin.ResultOk(con, &retData)
}

func (s *SquartController) getBanners(con *gin.Context) {

	var retBanner login.Banners

	banners := models.GetBanners()

	for _, banner := range banners {
		var onebanner login.BannerInfo
		onebanner.BannerId = banner.BannerID
		onebanner.TargetType = banner.Target_type

		onebanner.StartTime = banner.Start_time.Unix()
		onebanner.EndTime = banner.End_time.Unix()
		onebanner.BgImg = banner.Bg_img
		onebanner.ClickUrl = banner.Click_url
		retBanner.BannerInfo = append(retBanner.BannerInfo, &onebanner)
	}

	fmt.Println("getBanners", banners, "\t", retBanner)

	tqgin.ResultOk(con, &retBanner)
}
