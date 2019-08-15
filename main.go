// tqTestGin project main.go
package main

import (
	"time"
	"tqgin/config"
	"tqgin/models"
	"tqgin/pkg/tqlog"
	"tqgin/routers"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func main() {
	tqlog.ConfigLog()
	models.ConfigDB()

	//gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	routers.Router(router)

	endless.DefaultMaxHeaderBytes = 1 << 20
	endless.DefaultReadTimeOut = 30 * time.Second
	endless.DefaultWriteTimeOut = 30 * time.Second
	server := endless.NewServer(config.Tqconfig.String("httpIP")+":"+config.Tqconfig.String("httpport"), router)

	err := server.ListenAndServe()
	if err != nil {
		tqlog.TQSysLog.Warn("init server error")
	}
	defer models.DB.Close()
}
