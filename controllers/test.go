package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)
var globalSessions *session.Manager

func init(){
	sessionConfig := &session.ManagerConfig{
		CookieName:"gosessionid",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 7200,
		Secure: false,
		CookieLifeTime: 3600,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}


type TestController struct {
	beego.Controller
}

func (tc *TestController) Get() {
	test := tc.GetString("test")

	sess, _ := globalSessions.SessionStart(tc.Ctx.ResponseWriter, tc.Ctx.Request)
	defer sess.SessionRelease(tc.Ctx.ResponseWriter)

	//（5）根据当前请求对象，设置一个session
	res := sess.Get(test)
	if res != nil{
		tc.Ctx.WriteString(res.(string))
	}else{
		tc.Ctx.WriteString("no message")
	}


}

func (tc *TestController) Put() {

	test := tc.GetString("test")

	sess, _ := globalSessions.SessionStart(tc.Ctx.ResponseWriter, tc.Ctx.Request)
	defer sess.SessionRelease(tc.Ctx.ResponseWriter)

	//（5）根据当前请求对象，设置一个session
	sess.Set("test", test)
	tc.Ctx.WriteString(test)
}