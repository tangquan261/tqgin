package routers

import (
	"fmt"
	"net/http"
	"tqgin/controllers"

	//"tqgin/docs"
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

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//不需要验证登录
	authModel(router)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		new(controllers.UploadController).RegisterRouter(router)
		LoginModel(apiv1)
		userInfoModel(apiv1)
		rankInfoModel(apiv1)
		roomInfoModel(apiv1)
	}

	fmt.Println("router init success...")
}

//登录模块
func authModel(router *gin.Engine) {
	new(controllers.AuthController).RegisterRouter(router)
}

func LoginModel(router *gin.RouterGroup) {
	new(controllers.UserinfoController).RegisterRouter(router)
}

//用户信息模块
func userInfoModel(router *gin.RouterGroup) {
	new(controllers.UserInfoMoneyController).RegisterRouter(router)
	new(controllers.RelationController).RegisterRouter(router)
	new(controllers.CycleController).RegisterRouter(router)
	new(controllers.CycleCommetController).RegisterRouter(router)
}

//房间信息模块
func roomInfoModel(router *gin.RouterGroup) {
	new(controllers.RoomManagerController).RegisterRouter(router)
	new(controllers.SquartController).RegisterRouter(router)
	new(controllers.MicoController).RegisterRouter(router)
}

//排行榜
func rankInfoModel(router *gin.RouterGroup) {
	new(controllers.RankController).RegisterRouter(router)
}
