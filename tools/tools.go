package tools

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const ASCII_LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
const ASCII_UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ASCII_LETTERS = ASCII_LOWERCASE + ASCII_UPPERCASE
const DIGITS = "0123456789"


//将字符串数组拼接成一个字符串
func ArrToString(arr []string) string{
	str := strings.Trim(fmt.Sprint(arr), "[]")
	str = strings.Replace(str, " ", "", -1)
	return str
}


//将[]byte 转换成字符  例子：\x34\xaa
func BytesToHexString(bt []byte, sign bool)(string){
	var re string
	if sign{
		for k := range bt{
			//fmt.Println(k)
			re += `\x` + hex.EncodeToString(bt[k:k+1])
		}
	}else {
		for k := range bt{
			//fmt.Println(k)
			re += hex.EncodeToString(bt[k:k+1])
		}
	}
	return re
}


func HexStringToBytes(hs string, sign bool)([]byte, error){
	var str string
	if sign{
		//fmt.Println(hs)
		str = strings.Replace(hs, `\x`, "", -1)
		//fmt.Println(str)

	}else{
		str = hs
	}
	re, err := hex.DecodeString(str)
	if err != nil{
		//fmt.Println(err)
		return nil, err
	}
	return re, nil
}

//？？？表示将32位的主机字节顺序转化为32位的网络字节顺序
//int to byte
func SockHtonl(input int)[]byte{

	reBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(reBytes, uint32(input))

	return reBytes
}

//byte to int
func SockNonhl(inputBytes []byte) uint32{

	//reBytes := make([]byte, 4)
	reInt := binary.BigEndian.Uint32(inputBytes)

	return reInt
}





//生成随机字符串范围包括大小写字母，数字
func GetRandomStr(num int, lower bool, upper bool, digits bool)string{
	var re string
	var list string
	if lower{
		list += ASCII_LOWERCASE
	}
	if upper{
		list += ASCII_UPPERCASE
	}
	if digits{
		list += DIGITS
	}
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<num; i++{
		index := rand.Intn(len(list)-1)
		re += string(list[index])
	}

	return re
}


//补位
func Pkcs7Encode(text []byte, blockSize int)[]byte{
	//计算需要填充的位数
	needSize := blockSize - (len(text) % blockSize)
	if needSize == 0 {
		needSize = blockSize
	}
	//获得补位字符
	//Repeat()函数的功能是把切片b复制count个,然后合成一个新的字节切片返回.
	//func Repeat(b[]byte,count int) []byte
	pad := bytes.Repeat([]byte{byte(needSize)}, needSize)
	return append(text, pad...)
}


//
func Pkcs7Decode(text []byte)[]byte{

	length := len(text)
	unpadding := int(text[length-1])

	return text[:(length - unpadding)]
}


//AES 加密
func AesEncrypt(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	encodeBytes = Pkcs7Encode(encodeBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(key[:16]))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)
	// 使用BASE64对加密后的字符串进行编码
	fmt.Println(crypted)
	base64.StdEncoding.EncodeToString(crypted)
	return base64.StdEncoding.EncodeToString(crypted), nil
}


func AesDecrypt(crypted, key []byte) ([]byte, error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)

	//fmt.Println("削减")
	//fmt.Println(origData)
	//fmt.Println(string(origData))
	origData = Pkcs7Decode(origData)
	return origData, nil

}



