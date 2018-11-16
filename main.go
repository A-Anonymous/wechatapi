package main

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"time"
	_ "wechatapi/routers"
	"wechatapi/tools"
)

func main() {

	//t, err := tools.EncryptMsg("1", "1", "1", "1", "1", 0)
	//fmt.Println(t)
	//fmt.Println(err)

	//test()

	//hontl()


	token := "spamtest"
	aesKey := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG"
	xmlStr := ` <xml><ToUserName><![CDATA[oia2TjjewbmiOUlr6X-1crbLOvLw]]></ToUserName><FromUserName><![CDATA[gh_7f083739789a]]></FromUserName><CreateTime>1407743423</CreateTime><MsgType>  <![CDATA[video]]></MsgType><Video><MediaId><![CDATA[eYJ1MbwPRJtOvIEabaxHs7TX2D-HV71s79GUxqdUkjm6Gs2Ed1KF3ulAOA9H1xG0]]></MediaId><Title><![CDATA[testCallBackReplyVideo]]></Title><Descript  ion><![CDATA[testCallBackReplyVideo]]></Description></Video></xml>`
	appId := "wx2c2769f8efd9abc2"
	nonce := "1320562132"
	timestamp := int(time.Now().Unix())

	fmt.Println(timestamp)

	re, err := tools.EncryptMsg(token, aesKey, appId,
		xmlStr, nonce, timestamp)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(re)

		fmt.Println("加密结束，开始验证解密")
		result, err := tools.DecryptMsg(re, token, aesKey,
			appId, nonce, timestamp)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}


	beego.Info(beego.BConfig.AppName, "V0.1")

	beego.Run()
}

func test (){
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"xml"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "   ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}


func hontl(){
	to_xml := ` <xml><ToUserName><![CDATA[oia2TjjewbmiOUlr6X-1crbLOvLw]]></ToUserName><FromUserName><![CDATA[gh_7f083739789a]]></FromUserName><CreateTime>1407743423</CreateTime><MsgType>  <![CDATA[video]]></MsgType><Video><MediaId><![CDATA[eYJ1MbwPRJtOvIEabaxHs7TX2D-HV71s79GUxqdUkjm6Gs2Ed1KF3ulAOA9H1xG0]]></MediaId><Title><![CDATA[testCallBackReplyVideo]]></Title><Descript  ion><![CDATA[testCallBackReplyVideo]]></Description></Video></xml>`
	tools.Pkcs7Encode([]byte(to_xml), 32)

}


