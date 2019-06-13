// tqTestGin project main.go
package main

import (
	"fmt"

	"tqgin/models"
	"tqgin/routers"

	"github.com/astaxie/beego/config"
	"github.com/gin-gonic/gin"
	//tqlog "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// tqlog.SetFormatter(&tqlog.JSONFormatter{})

	// // Output to stdout instead of the default stderr
	// // Can be any io.Writer, see below for File example
	// tqlog.SetOutput(os.Stdout)

	// // Only log the warning severity or above.
	// tqlog.SetLevel(tqlog.WarnLevel)
}

var (
	router   *gin.Engine
	Tqconfig config.Configer
)

func main() {

	var err error
	Tqconfig, err = config.NewConfig("ini", "./conf/app.conf")

	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println(Tqconfig)

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	routers.Router(router)

	fmt.Println("server start.....")
	router.Run(Tqconfig.String("httpIP") + ":" + Tqconfig.String("httpport"))

	defer models.DB.Close()
}
