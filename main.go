package main

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	_ "wechatapi/routers"
	"wechatapi/tools"
)

func main() {

	//t, err := tools.EncryptMsg("1", "1", "1", "1", "1", 0)
	//fmt.Println(t)
	//fmt.Println(err)

	//test()

	//hontl()

//
//	token := "123qwe"
//	aesKey := "4NkW5VLJFj6RhSuaqtw6oiA0cm1Cgm0YPTQJT3jyzp0"
//	xmlStr := ` <xml><ToUserName><![CDATA[oia2TjjewbmiOUlr6X-1crbLOvLw]]></ToUserName><FromUserName><![CDATA[gh_7f083739789a]]></FromUserName><CreateTime>1407743423</CreateTime><MsgType>  <![CDATA[video]]></MsgType><Video><MediaId><![CDATA[eYJ1MbwPRJtOvIEabaxHs7TX2D-HV71s79GUxqdUkjm6Gs2Ed1KF3ulAOA9H1xG0]]></MediaId><Title><![CDATA[testCallBackReplyVideo]]></Title><Description><![CDATA[testCallBackReplyVideo]]></Description></Video></xml>`
//	appId := "wxfcc7077466b8eb8e"
//	nonce := "1320562132"
//	timestamp := 12345678
//
//	fmt.Println(timestamp)
//
//	re, err := tools.EncryptMsg(token, aesKey, appId,
//		xmlStr, nonce, timestamp)
//	if err != nil{
//		fmt.Println(err)
//	}else{
//		fmt.Println(re)
//
//		fmt.Println("加密结束，开始验证解密")
////		en := `<xml>
////   <ToUserName><![CDATA[gh_f0cbc1e91308]]></ToUserName>
////   <FromUserName><![CDATA[oNFs_xFslFEqq3EcuHTNGvuPTVTE]]></FromUserName>
////   <CreateTime>1542355081</CreateTime>
////   <MsgType><![CDATA[text]]></MsgType>
////   <Content><![CDATA[te]]></Content>
////   <MsgId>6624364632441363610</MsgId>
////   <Encrypt><![CDATA[UN83+NiuiBdqgHI4MvF+FlBIOKm5XGE26nzj1Gql9ZPvpjr17GHhJ/jsIB5bB3yVhZtNJNu29k2NJnXoDCZeNqyycc/L+x6i2TryHLQ4FSh5RiqF/0E69IzhHqlOcgTR8oEnhnCk+tL0rV5HFqzRUTplb8kd1JXu5zFWTG8dsIdJ0+1nYYPFbExFVOW4GUMrRfzTNrMxRTnUtJFAWw16A8FpOROBO8XRDYr/OrJSLxbS0HPSOZSobo6Tmiv6WU1gSq7qIPehvDWkGTOZP752iOAwKgS6kD7I9b1vf92EI1K/zSJAs8GEmkhR1xGxYiXVEqkKeY5+rneh7Ed/OKpUcT4Yy81q3OQ1izalBu+VxGwWLELQWX1uIDkpPYRhHNCxlliEtqGOSO1BtI8qoGcs8nDt/HYSseCbuz5Py7mCKL4=]]></Encrypt>
////</xml>`
//		result, err := tools.DecryptMsg(re, token, aesKey,
//			appId, nonce, timestamp)
//		if err != nil{
//			fmt.Println("error:	", err)
//			return
//		}
//		fmt.Println("result:	", result)
//		var t models.InfoX
//		xml.Unmarshal([]byte(``), &t)
//		fmt.Println(t)
//	}
//
//

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


