// tqTestGin project main.go
package main

import (
	"fmt"

	"tqgin/models"
	"tqgin/routers"

	"tqgin/config"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	tqlog "github.com/sirupsen/logrus"
	//"github.com/unrolled/secure"
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
	uid := ksuid.New()

	fmt.Println(uid.String())

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	routers.Router(router)

	router.Use(TlsHandler())
	fmt.Println("server start.....")
	//router.RunTLS(config.Tqconfig.String("httpIP")+":"+config.Tqconfig.String("httpport"), "miban.pem", "miban.key")

	router.Run(config.Tqconfig.String("httpIP") + ":" + config.Tqconfig.String("httpport"))
	defer models.DB.Close()
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// secureMiddleware := secure.New(secure.Options{
		// 	SSLRedirect: true,
		// 	SSLHost:     "localhost:8089",
		// })
		// err := secureMiddleware.Process(c.Writer, c.Request)

		// // If there was an error, do not continue.
		// if err != nil {
		// 	return
		// }

		c.Next()
	}
}
