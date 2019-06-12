package routers

import (
	"fmt"
	"tqTestGin/controllers"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {

	fmt.Println("router init...")

	new(controllers.LoginController).RegisterRouter(router)

}
