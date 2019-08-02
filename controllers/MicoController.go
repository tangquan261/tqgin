package controllers

import (
	// "fmt"
	// "strconv"
	"tqgin/common"
	// "tqgin/models"

	"github.com/gin-gonic/gin"
)

type MicoController struct {
	tqgin.Controller
}

func (this *MicoController) RegisterRouter(router *gin.RouterGroup) {
	temp := router.Group("/mico")
	temp.POST("mico_add", this.micoAdd)
	temp.POST("mico_del", this.micoDel)

}

//加入mico
func (r *MicoController) micoAdd(c *gin.Context) {

}

//移除mico
func (r *MicoController) micoDel(c *gin.Context) {

}
