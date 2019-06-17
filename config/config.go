package config

import (
	"log"

	"github.com/astaxie/beego/config"
)

var Tqconfig config.Configer

func init() {

	var err error
	Tqconfig, err = config.NewConfig("ini", "./conf/app.conf")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("config success init")
}
