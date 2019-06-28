/*
	房间列表
*/

package controllers

import (
	"fmt"
	//"fmt"
	//	"strconv"
	"tqgin/common"

	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type SquartController struct {
	tqgin.Controller
}

func (this *SquartController) RegisterRouter(router *gin.Engine) {
	temp := router.Group("/squart")
	temp.GET("applyTagsList", this.applyTagsList)
	temp.POST("applyRoomList", this.applyRoomList)
}

func (s *SquartController) applyTagsList(con *gin.Context) {

	var status int

	tags := models.GetTagList()

	fmt.Println(status, tags)
	//tqgin.Result(con, status, gin.H{"tags": tags}, "")
}

func (s *SquartController) applyRoomList(con *gin.Context) {
	TagName := con.PostForm("tagName")
	var status int
	var data []*models.HotRoomInfo
	if len(TagName) > 0 {
		status = 0
		//data = models.GetHotRoomsByTag(TagName)
	} else {
		status = 1
	}
	fmt.Println(status, data)
	//tqgin.Result(con, status, gin.H{"rooms": data}, "")
}
