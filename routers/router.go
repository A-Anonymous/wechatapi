package routers

import (
	"github.com/astaxie/beego"
	"wechatapi/controllers"
)

func init(){
	beego.SetStaticPath("/static", "static")
	beego.Router("/wechat", &controllers.WeChatController{})
	beego.Router("/test", &controllers.TestController{})
	//beego.Router("/search", &controllers.SearchController{})
}
