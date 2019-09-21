package controllers

//称号

import (
	"tqgin/common"
	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type TitleController struct {
	tqgin.Controller
}

func (this *TitleController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/title")
	temp.POST("getMyTitleList", this.getMyTitleList)
	temp.POST("modifyTitle", this.modifyTitle)

}

type modifyTitle struct {
	Title_ID int32 `json:"titleID"`
}

func (r *TitleController) modifyTitle(con *gin.Context) {
	playerID := r.GetPlayerGUID(con)

	var param modifyTitle

	err := con.ShouldBindJSON(&param)
	if err != nil {
		tqgin.ResultFail(con, "错误")
		return
	}

	models.ModifyWearTitle(playerID, param.Title_ID)

	tqgin.ResultOk(con, param)
}

type TitleUserinfo struct {
	Title_ID      int32  `json:"titleID"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	IconAnimation int    `json:"iconAnimation"`
	IsWear        int32  `json:"isWear"`
}

func (r *TitleController) getMyTitleList(con *gin.Context) {
	playerID := r.GetPlayerGUID(con)

	var rets []TitleUserinfo

	titles := models.TitlesGetByPlayerID(playerID)

	for _, obj := range titles {

		titleconfig := models.GetTitlemodel(obj.Title_ID)

		if titleconfig != nil {
			var title TitleUserinfo
			title.Title_ID = obj.Title_ID
			title.Icon = titleconfig.Icon
			title.IconAnimation = titleconfig.Icon_animation
			title.Name = titleconfig.Name
			title.IsWear = obj.IsWear
			rets = append(rets, title)
		}
	}

	tqgin.ResultOkMsg(con, rets, "成功")
}
