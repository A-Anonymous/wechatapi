package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"sort"
	"weChatAPI/models"
	"weChatAPI/tools"
)

type WeChatController struct {
	beego.Controller
}

func (wcc *WeChatController) Get() {
	//mc.Render("/index.html")

	//mc.TplName = "index.html"
	//mc.Ctx.WriteString("hello world")
	check := models.Check{}
	check.Echostr = wcc.GetString("echostr")
	check.Nonce = wcc.GetString("nonce")
	check.Signature = wcc.GetString("signature")
	check.Timestamp = wcc.GetString("timestamp")

	bl, err := checkSignature(check)
	if err != nil{
		wcc.Ctx.WriteString("error")
	}
	if bl{
		wcc.Ctx.WriteString(check.Echostr)
	}else{
		wcc.Ctx.WriteString("check fail")
	}

}


func (wcc *WeChatController) Post() {
	//mc.Render("/index.html")

	//mc.TplName = "index.html"
	//mc.Ctx.WriteString("hello world")
	//fmt.Println(wcc.Ctx.Input.RequestBody)
	//fmt.Println(string(wcc.Ctx.Input.RequestBody))

	infoX := models.InfoX{}
	err := xml.Unmarshal(wcc.Ctx.Input.RequestBody, &infoX)
	if err != nil {
		//fmt.Printf("error: %v", err)
		return
	}


	fmt.Println(infoX.Content)

}


func checkSignature (check models.Check) (bool, error){
    //token、timestamp、nonce
    //token := "123qwe"
	token := beego.AppConfig.String("token")
	var msg []string
	msg = append(msg, token)
	msg = append(msg, check.Timestamp)
	msg = append(msg, check.Nonce)
	sort.Strings(msg)
	str := tools.ArrToString(msg)
	res, err := tools.SHA1(str)
	if err != nil{
		return false, err
	}
	if res == check.Signature{
		return true, nil
	}else{
		return false, nil
	}
}