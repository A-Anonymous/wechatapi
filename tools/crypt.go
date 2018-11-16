package tools

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"sort"
	"strconv"
	"wechatapi/models"
)


//wechat 加密
func EncryptMsg(token string, aesKey string, appId string,
	xmlStr string, nonce string, timestamp int)(string, error){
	decodeBytes, err := base64.StdEncoding.DecodeString(aesKey + "=")
	if err != nil {
	     fmt.Println(err)
	}
	fmt.Println("a", string(decodeBytes))

	//明文长度
	lenStr := BytesToHexString(SockHtonl(len(xmlStr)),true)
	randomstr := GetRandomStr(16, true, true, true)
	text := randomstr + lenStr + xmlStr + appId
	//fmt.Println("len", len(xmlStr))

	//加密
	cryptograph, err := AesEncrypt(text, decodeBytes)
	if err != nil{
		return "", err
	}
	fmt.Println("cryptograph")
	fmt.Println(cryptograph)

	//生成安全签名
	var msg []string
	msg = append(msg, token)
	msg = append(msg, strconv.Itoa(timestamp))
	msg = append(msg, nonce)
	msg = append(msg, cryptograph)
	sort.Strings(msg)
	str := ArrToString(msg)
	signature, err := SHA1(str)
	if err != nil{
		return "", err
	}


	fmt.Println("加密时生成的签名：", signature)
	var encrypt models.Encrypt
	encrypt.Signature = signature
	encrypt.Nonce = nonce
	encrypt.Timestamp = timestamp
	encrypt.Encrypt = cryptograph

	data, err := xml.MarshalIndent(&encrypt, "", "\t")
	if err != nil {
		//fmt.Println(err)
		return "", nil
	}
	//text = string(data)
	//fmt.Println("test")
	//fmt.Println(string(data))
	//fmt.Println()

	return string(data), nil
}

//wechat 解密
func DecryptMsg(postMsg string, token string, aesKey string,
	appId string, nonce string, timestamp int)(string, error){

	var decrypt models.Encrypt


	//解码xml
	xml.Unmarshal([]byte(postMsg), &decrypt)
	fmt.Println(decrypt)

	//验签
	var msg []string
	msg = append(msg, token)
	msg = append(msg, strconv.Itoa(timestamp))
	msg = append(msg, nonce)
	msg = append(msg, decrypt.Encrypt)
	sort.Strings(msg)
	str := ArrToString(msg)
	signature, err := SHA1(str)
	fmt.Println("解密时生成的签名：", signature)
	fmt.Println(signature)
	fmt.Println(decrypt.Signature)
	if err != nil{
		return "", err
	}
	if signature != decrypt.Signature{
		return "验证失败", nil
	}

	//AES解密
	decodeBytes, err := base64.StdEncoding.DecodeString(aesKey + "=")
	if err != nil {
		fmt.Println(err)
	}
	encrypt, err := base64.StdEncoding.DecodeString(decrypt.Encrypt)
	if err != nil {
		fmt.Println(err)
	}
	content, err := AesDecrypt([]byte(encrypt), decodeBytes)

	fmt.Println("解密后的明文")
	fmt.Println(decrypt.Encrypt)
	fmt.Println(string(content))


	//去除16随机字符
	content = content[16:]
	fmt.Println("去除16随机字符")
	fmt.Println(string(content))

	fmt.Println(string(content[:4]))

	byteS, err := HexStringToBytes(string(content[:16]), true)
	if err != nil{
		return "", err
	}
	lenXml := SockNonhl(byteS)
	fmt.Println(lenXml)
	fmt.Println(string(content[lenXml + 16:]))
	if string(content[lenXml + 16:]) != appId{
		fmt.Println("cuole")
	}
	re := string(content[16:lenXml+16])
	return re, nil

}


