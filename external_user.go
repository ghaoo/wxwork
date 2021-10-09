package wxwork

import (
	"bytes"
	"encoding/json"
)

// ExternalProfile 成员对外信息
// 文档地址: https://work.weixin.qq.com/api/doc/90000/90135/92230
type ExternalProfile struct {

	// 企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。
	CorpName string `json:"external_corp_name,omitempty"`

	// 视频号属性
	WechatChannels WechatChannels `json:"wechat_channels,omitempty"`

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

// FollowUser 企业服务人员
type FollowUser struct {
	// 企业成员userid
	Userid string `json:"userid"`

	// 企业成员对此外部联系人的备注
	Remark string `json:"remark,omitempty"`

	// 企业成员对此外部联系人的描述
	Description string `json:"description,omitempty"`

	// 企业成员添加此外部联系人的时间
	CreateTime int64 `json:"createtime,omitempty"`

	Tags
}

// ExternalContact 外部联系人详情
type ExternalContact struct {
	// 外部联系人的userid
	ExternalUserid string `json:"external_userid,omitempty"`

	// 外部联系人的名称
	Name string `json:"name,omitempty"`

	// 外部联系人头像，代开发自建应用需要管理员授权才可以获取，第三方不可获取
	Avatar string `json:"avatar,omitempty"`

	// 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Type int `json:"type,omitempty"`

	// 外部联系人性别 0-未知 1-男性 2-女性
	Gender int `json:"gender,omitempty"`

	// 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段
	Unionid string `json:"unionid,omitempty"`

	// 外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，仅当联系人类型是企业微信用户时有此字段
	Position string `json:"position,omitempty"`

	// 外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
	CorpName string `json:"corp_name,omitempty"`

	// 外部联系人所在企业的主体名称，仅当联系人类型是企业微信用户时有此字段
	CorpFullName string `json:"corp_full_name,omitempty"`

	// 外部联系人的自定义展示信息，可以有多个字段和多种类型，包括文本，网页和小程序，仅当联系人类型是企业微信用户时有此字段
	ExternalProfile ExternalProfile `json:"external_profile,omitempty"`

	// 添加了此外部联系人的企业成员
	FollowUser User
}

// ExternalCustomerList 获取客户列表
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/92113
func (a *Agent) ExternalCustomerList(userid string) ([]string, error) {
	param := map[string]string{
		"userid": userid,
	}
	body, _ := json.Marshal(param)

	var caller struct {
		baseCaller
		ExternalUserid []string `json:"external_userid"`
	}

	err := a.ExecuteWithToken("POST", "externalcontact/list", nil, bytes.NewReader(body), &caller)

	return caller.ExternalUserid, err
}

// ExternalUserGet 获取客户详情
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/92114
func (a *Agent) ExternalUserGet(external_userid, cursor string) (ec ExternalContact, next_cursor string, err error) {
	param := map[string]string{
		"external_userid": external_userid,
		"cursor":          cursor,
	}
	body, _ := json.Marshal(param)

	var caller struct {
		baseCaller
		ExternalContact ExternalContact `json:"external_contact"`
		NextCursor      string          `json:"next_cursor"`
	}

	err = a.ExecuteWithToken("POST", "externalcontact/get", nil, bytes.NewReader(body), &caller)

	return caller.ExternalContact, caller.NextCursor, err
}
