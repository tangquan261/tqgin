// tqTestGin project main.go
package main

import (
	"log"
	"time"
	"tqgin/config"
	"tqgin/models"
	"tqgin/routers"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	tqlog "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	tqlog.SetFormatter(&tqlog.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//tqlog.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	tqlog.SetLevel(tqlog.WarnLevel)
}

var (
	router *gin.Engine
)

func main() {

	//gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	routers.Router(router)

	endless.DefaultMaxHeaderBytes = 1 << 20
	endless.DefaultReadTimeOut = 30 * time.Second
	endless.DefaultWriteTimeOut = 30 * time.Second
	server := endless.NewServer(config.Tqconfig.String("httpIP")+":"+config.Tqconfig.String("httpport"), router)

	err := server.ListenAndServe()
	if err != nil {
		log.Println("err:%v", err)
	}
	defer models.DB.Close()
}
