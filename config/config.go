package config

import (
	"log"

	"github.com/astaxie/beego/config"
)

var Tqconfig config.Configer

const (
	JwtSecret       = "JwtSecret"
	PrefixUrl       = "PrefixUrl"
	RuntimeRootPath = "RuntimeRootPath"
	ImageSavePath   = "ImageSavePath"
	ImageMaxSize    = "ImageMaxSize"
	ImageAllowExts  = "ImageAllowExts"
	QrCodeSavePath  = "QrCodeSavePath"
	FontSavePath    = "FontSavePath"

	LogSavePath = "LogSavePath"
	LogSaveName = "LogSaveName"
	LogFileExt  = "LogFileExt"
	TimeFormat  = "TimeFormat"

	//server
	RunMode      = "RunMode"
	HttpPort     = "HttpPort"
	ReadTimeout  = "ReadTimeout"
	WriteTimeout = "WriteTimeout"

	//redis
	Redis_Host        = "redis_Host"
	Redis_Password    = "redis_Password"
	REdis_DBIndex     = "redis_DBIndex"
	Redis_MaxIdle     = "redis_MaxIdle"
	Redis_MaxActive   = "redis_MaxActive"
	Redis_IdleTimeout = "redis_IdleTimeout"

	Snow_Work_Id = "snow_Work_Id"
)

func init() {
	var err error
	Tqconfig, err = config.NewConfig("ini", "/Users/tq/go/src/tqgin/conf/app.conf")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("config success init")
}

func GetConfigString(conf string) string {
	return Tqconfig.String(conf)
}

func GetConfigInt(conf string) int {
	Value, err := Tqconfig.Int(conf)

	if err != nil {
		log.Fatal(err)
	}

	return Value
}
