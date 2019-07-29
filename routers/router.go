package routers

import (
	"fmt"
	"tqgin/controllers"

	"tqgin/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	fmt.Println("router init begin ...")
	authModel(router)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		loginModel(apiv1)
		userInfoModel(apiv1)
		rankInfoModel(apiv1)
		roomInfoModel(apiv1)
	}
	fmt.Println("router init success...")
}

func authModel(router *gin.Engine) {
	new(controllers.LoginController).RegisterRouter(router)
}

func loginModel(router *gin.RouterGroup) {
	new(controllers.RoomManagerController).RegisterRouter(router)
	new(controllers.SquartController).RegisterRouter(router)
}

func userInfoModel(router *gin.RouterGroup) {

}

func roomInfoModel(router *gin.RouterGroup) {

}

func rankInfoModel(router *gin.RouterGroup) {

}
