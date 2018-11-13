package main

import (
	"github.com/astaxie/beego"
	_ "weChatAPI/routers"
)

func main() {

	beego.Info(beego.BConfig.AppName, "V0.1")

	beego.Run()
}
