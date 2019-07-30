package config

import (
	"log"
	"time"

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
	ExportSavePath  = "ExportSavePath"
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
)

func init() {

	var err error
	Tqconfig, err = config.NewConfig("ini", "./conf/app.conf")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("config success init")
}
