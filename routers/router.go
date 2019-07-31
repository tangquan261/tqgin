package routers

import (
	"fmt"
	"net/http"
	"tqgin/controllers"

	"tqgin/middleware/jwt"

	"tqgin/pkg/qrcode"
	"tqgin/pkg/upload"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Router(router *gin.Engine) {

	fmt.Println("router init begin ...")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	router.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	new(controllers.UploadController).RegisterRouter(router)

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
