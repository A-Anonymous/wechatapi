package routers

import (
	"github.com/astaxie/beego"
	"weChatAPI/controllers"
)

func init(){
	beego.SetStaticPath("/static", "static")
	beego.Router("/wechat", &controllers.WeChatController{})
	//beego.Router("/search", &controllers.SearchController{})
}
