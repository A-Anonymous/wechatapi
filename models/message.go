package models

//没用
type Message struct{
	Message string `json:"message"`
}

//介入校验
type Check struct{
	Signature 		string  `json:"signature"`
	Timestamp	 	string  `json:"timestamp"`
	Nonce  			string  `json:"nonce"`
	Echostr  		string  `json:"echostr"`
}

//消息xml
type InfoX struct{
	ToUserName 		string `xml:"ToUserName"`    	// 接收方微信号
	FromUserName 	string `xml:"FromUserName"`		// 发送方微信号，若为普通用户，则是一个OpenID
	CreateTime 		string `xml:"CreateTime"`		// 消息创建时间 （整型）
	MsgType		 	string `xml:"MsgType"`			// 消息类型{image, text, link, location, shortvideo, video, voice}
	Content		 	string `xml:"Content"`			// 文本消息内容
	MsgId		 	string `xml:"MsgId"`			// 消息id，64位整型
	Mediald		 	string `xml:"Mediald"`			// 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。;语音消息媒体id，可以调用多媒体文件下载接口拉取该媒体;视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Format			string `xml:"Format"`			// 语音格式：amr
	Recognition		string `xml:"Recognition"`		// 语音识别结果，UTF8编码
	ThumbMediaId	string `xml:"ThumbMediaId"`		// 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据
	LocationX		string `xml:"Location_X"`		// 地理位置维度
	LocationY		string `xml:"Location_Y"`		// 地理位置经度
	Scale			string `xml:"Scale"`			// 地图缩放大小
	Label			string `xml:"Label"`			// 地理位置信息
	Title			string `xml:"Title"`			// 消息标题
	Description		string `xml:"Description"`		// 消息描述
	Url				string `xml:"Url"`				// 消息链接
	PicUrl			string `xml:"PicUrl"`			// 图片链接（由系统生成）
}