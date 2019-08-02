/*
	用户货币管理
*/

package controllers

import (
	//"errors"
	//"log"

	// "fmt"
	// "strconv"
	"tqgin/common"
	"tqgin/models"

	"github.com/gin-gonic/gin"
)

type UserInfoMoneyController struct {
	tqgin.Controller
}

func (this *UserInfoMoneyController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/money")
	temp.POST("add_diamond", this.add_diamond)

}

type MoneyJsonAdd struct {
	Diamond int64 `json:"diamond"`
	Gold    int64 `json:"gold"`
	Cash    int64 `json:"cash"`
}

func (r *UserInfoMoneyController) add_diamond(c *gin.Context) {
	playerGUID := r.GetPlayerGUID(c)
	if playerGUID <= 0 {
		tqgin.ResultFail(c, "not find")
		return
	}

	var money MoneyJsonAdd

	err := c.ShouldBindJSON(&money)
	if err != nil {
		tqgin.ResultFail(c, "解析错误")
		return
	}

	err = AddPlayerDiamond(playerGUID, money.Diamond)
	if err != nil {
		tqgin.ResultFail(c, "添加错误")
		return
	}
	tqgin.ResultOkMsg(c, money, "添加成功")
}

func AddPlayerDiamond(playerGUID, diamond int64) error {

	diamond = 1
	for i := 0; i < 1; i++ {
		go func() {

			models.ModifyDinamondUser(playerGUID, -1)
		}()
	}
	return nil
}
