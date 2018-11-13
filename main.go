package main

import (
	"github.com/astaxie/beego"
	_ "wechatAPI/routers"
)

func main() {

	beego.Info(beego.BConfig.AppName, "V0.1")

	beego.Run()
}
