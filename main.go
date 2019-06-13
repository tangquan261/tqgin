// tqTestGin project main.go
package main

import (
	"fmt"
	"log"

	_ "tqgin/models"
	"tqgin/routers"

	"github.com/astaxie/beego/config"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func main() {

	tqConfig, err := config.NewConfig("ini", "./conf/app.conf")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tqConfig)

	router = gin.Default()

	routers.Router(router)

	router.Run(":" + tqConfig.String("httpport"))
}
