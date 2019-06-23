// tqTestGin project main.go
package main

import (
	"context"
	"fmt"
	"log"

	"tqgin/models"
	"tqgin/routers"

	"tqgin/config"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	tqlog "github.com/sirupsen/logrus"
	//"github.com/unrolled/secure"

	"proto"

	"google.golang.org/grpc"
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

	conn, err := grpc.Dial("localhost:5262", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := proto.NewHelloClient(conn)

	request, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "tq"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(request.Message)

	uid := ksuid.New()

	fmt.Println(uid.String())

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	routers.Router(router)

	fmt.Println("server start.....")
	//router.RunTLS(config.Tqconfig.String("httpIP")+":"+config.Tqconfig.String("httpport"), "miban.pem", "miban.key")

	router.Run(config.Tqconfig.String("httpIP") + ":" + config.Tqconfig.String("httpport"))
	defer models.DB.Close()
}
