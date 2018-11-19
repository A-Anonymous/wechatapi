package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"sort"
	"time"
	"wechatapi/models"
	"wechatapi/tools"
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
	//获取默认参数
	token := beego.AppConfig.String("token")
	aesKey := beego.AppConfig.String("EncodingAESKey")
	appId := beego.AppConfig.String("appid")
	nonce := beego.AppConfig.String("nonce")
	timestamp := int(time.Now().Unix())

	//获得body数据
	fmt.Println("body数据:		", string(wcc.Ctx.Input.RequestBody))
	infoX := models.InfoX{}
	err := xml.Unmarshal(wcc.Ctx.Input.RequestBody, &infoX)
	if err != nil {
		//fmt.Printf("error: %v", err)
		return
	}


	//解密
	result, err := tools.DecryptMsg(infoX.Encrypt, token, aesKey,
		appId, nonce, timestamp)
	if err != nil{
		fmt.Println("error:	", err)
		return
	}

	fmt.Println("解密后的明文：	", result)

	re := models.InfoX{}
	err = xml.Unmarshal([]byte(result), &re)
	if err != nil {
		//fmt.Printf("error: %v", err)
		return
	}

	//对内容做处理

	//test

	re.Content = "This is a test!!!"
	re.CreateTime = int(time.Now().Unix())

	re.ToUserName = infoX.FromUserName
	re.FromUserName = infoX.ToUserName
	xmlStr, err := xml.Marshal(re)
	if err != nil{
		fmt.Println(err)
		return
	}

	//加密
	encrypte, err := tools.EncryptMsg(token, aesKey, appId,
		string(xmlStr), nonce, timestamp)
	if err != nil{
		fmt.Println(err)
	}

	res := models.Encrypted{}
	err = xml.Unmarshal([]byte(encrypte), &res)
	if err != nil {
		//fmt.Printf("error: %v", err)
		return
	}
	//fmt.Println("re", encrypte)
	wcc.Data["xml"]=&res
	wcc.ServeXML()





	// 非加密模式，测试
	//fmt.Println(infoX.Content)
	//if infoX.MsgType == "text"{
	//	textX := models.TextX{}
	//	textX.ToUserName = infoX.FromUserName
	//	textX.FromUserName = infoX.ToUserName
	//	textX.Content = "测试功能，无实际意义！！！"
	//	textX.CreateTime = int(time.Now().Unix())
	//	textX.MsgType = infoX.MsgType
	//	wcc.Data["xml"]=&textX
	//	wcc.ServeXML()
	//}else {
	//	wcc.Ctx.WriteString("success")
	//}



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


