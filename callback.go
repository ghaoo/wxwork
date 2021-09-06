package wxwork

// Callback 应用回调配置，需加密
type Callback struct {
	// 企业应用接收企业微信推送请求的访问协议和地址，支持http或https协议
	URL string `json:"url,omitempty" xml:"url,omitempty"`
	// 用于生成签名
	Token string `json:"token" xml:"token"`
	// 用于消息体的加密，是AES密钥的Base64编码
	EncodingAESKey string `json:"encodingaeskey" xml:"encodingaeskey"`
}
