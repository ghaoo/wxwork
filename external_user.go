package wxwork

// ExternalProfile 成员对外信息
// 文档地址: https://work.weixin.qq.com/api/doc/90000/90135/92230
type ExternalProfile struct {

	// 企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。
	CorpName string `json:"external_corp_name,omitempty"`

	// 视频号属性
	WechatChannels WechatChannels `json:"wechat_channels"`

	// 属性列表，目前支持文本、网页、小程序三种类型
	ExternalAttr ExternalAttr `json:"external_attr,omitempty"`
}

// Attrs 自定义字段
type Attrs struct {
	Attrs []ExternalAttr `json:"attrs"`
}

// ExternalAttr 自定义字段内容
type ExternalAttr struct {
	// 属性类型: 0-文本 1-网页 2-小程序
	Type int `json:"type,omitempty"`

	// 属性名称： 需要先确保在管理端有创建该属性，否则会忽略
	Name string `json:"name,omitempty"`

	// 文本类型的属性  type为0时必填
	Text TextAttr `json:"text,omitempty"`

	// 网页类型的属性，url和title字段要么同时为空表示清除该属性，要么同时不为空	type为1时必填
	Web WebAttr `json:"web,omitempty"`

	// 小程序类型的属性，appid和title字段要么同时为空表示清除改属性，要么同时不为空	type为2时必填
	Miniprogram MiniprogramAttr `json:"miniprogram,omitempty"`
}

// TextAttr 文本属性
type TextAttr struct {
	// 文本属性内容,长度限制12个UTF8字符
	Value string `json:"value,omitempty"`
}

// WebAttr 网页属性
type WebAttr struct {
	// 网页的url,必须包含http或者https头
	Url string `json:"url,omitempty"`

	// 网页的展示标题,长度限制12个UTF8字符
	Title string `json:"title,omitempty"`
}

// MiniprogramAttr 小程序属性
type MiniprogramAttr struct {
	// 小程序appid，必须是有在本企业安装授权的小程序，否则会被忽略
	Appid int `json:"appid,omitempty"`

	// 小程序的展示标题,长度限制12个UTF8字符
	Title string `json:"title,omitempty"`

	// 小程序的页面路径
	PagePath string `json:"pagepath,omitempty"`
}

// WechatChannels 视频号属性。须从企业绑定到企业微信的视频号中选择，可在“我的企业”页中查看绑定的视频号。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取。注意：externalcontact/get不返回该字段
type WechatChannels struct {
	// 视频号名字（设置后，成员将对外展示该视频号）
	NickName string `json:"nickname"`

	// 对外展示视频号状态。0表示企业视频号已被确认，可正常使用，1表示企业视频号待确认
	Status int `json:"status"`
}
