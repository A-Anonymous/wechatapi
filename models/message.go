package models

//没用
type Message struct{
	Message string `json:"message"`
}

//介入校验
type Check struct{
	Signature string  `json:"signature"`
	Timestamp	 string  `json:"timestamp"`
	Nonce  string  `json:"nonce"`
	Echostr  string  `json:"echostr"`
}
