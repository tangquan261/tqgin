package routers

import (
	"fmt"
	"tqgin/controllers"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	fmt.Println("router init begin ...")

	new(controllers.LoginController).RegisterRouter(router)

	fmt.Println("router init success...")

}
