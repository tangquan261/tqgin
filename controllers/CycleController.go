/*
	房间管理
	创建，移除，修改操作
*/

package controllers

import (
	// "fmt"
	// "strconv"
	"tqgin/common"
	// "tqgin/models"

	"github.com/gin-gonic/gin"
)

type CycleController struct {
	tqgin.Controller
}

func (this *CycleController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/cycle")
	temp.POST("create", this.CommitArt)

}

func (r *CycleController) CommitArt(c *gin.Context) {

}
