package wxwork

// 成员对外信息
// 文档地址: https://work.weixin.qq.com/api/doc/90000/90135/92230
type ExternalProfile struct {

	// 企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。
	CorpName string `json:"external_corp_name,omitempty"`

	// 属性列表，目前支持文本、网页、小程序三种类型
	ExternalAttr ExternalAttr `json:"external_attr,omitempty"`
}

type Attrs struct {
	Attrs []ExternalAttr `json:"attrs"`
}

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

type TextAttr struct {
	// 文本属性内容,长度限制12个UTF8字符
	Value string `json:"value,omitempty"`
}

type WebAttr struct {
	// 网页的url,必须包含http或者https头
	Url string `json:"url,omitempty"`

	// 网页的展示标题,长度限制12个UTF8字符
	Title string `json:"title,omitempty"`
}

type MiniprogramAttr struct {
	// 小程序appid，必须是有在本企业安装授权的小程序，否则会被忽略
	Appid int `json:"appid,omitempty"`

	// 小程序的展示标题,长度限制12个UTF8字符
	Title string `json:"title,omitempty"`

	// 小程序的页面路径
	Pagepath string `json:"pagepath,omitempty"`
}
