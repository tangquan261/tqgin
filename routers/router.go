package routers

import (
	"fmt"
	"tqgin/controllers"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	fmt.Println("router init begin ...")
	loginModel(router)
	UserInfoModel(router)
	rankInfoModel(router)
	RoomInfoModel(router)
	fmt.Println("router init success...")
}

func loginModel(router *gin.Engine) {
	new(controllers.LoginController).RegisterRouter(router)
	new(controllers.RoomManagerController).RegisterRouter(router)
	new(controllers.SquartController).RegisterRouter(router)
}

func UserInfoModel(router *gin.Engine) {

}

func RoomInfoModel(router *gin.Engine) {

}

func rankInfoModel(router *gin.Engine) {

}
