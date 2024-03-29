package wxwork

import (
	"bytes"
	"encoding/json"
)

// ExternalGetFollowUserList 获取配置了客户联系功能的成员列表
// 文档: https://developer.work.weixin.qq.com/document/path/92571
func (a *Agent) ExternalGetFollowUserList() ([]string, error) {
	var caller struct {
		baseCaller
		FollowUser []string `json:"follow_user"`
	}

	err := a.ExecuteWithToken("GET", "externalcontact/get_follow_user_list", nil, nil, &caller)

	return caller.FollowUser, err
}

// ContactWay 配置客户联系「联系我」方式
type ContactWay struct {
	ConfigID string `json:"config_id,omitempty"`

	// 联系方式类型,1-单人, 2-多人
	Type int `json:"type"`

	// 场景，1-在小程序中联系，2-通过二维码联系
	Scene int `json:"scene"`

	// 在小程序中联系时使用的控件样式，详见附表
	Style int `json:"style,omitempty"`

	// 联系方式的备注信息，用于助记，不超过30个字符
	Remark string `json:"remark,omitempty"`

	// 外部客户添加时是否无需验证，默认为true
	SkipVerify bool `json:"skip_verify,omitempty"`

	// 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符
	State string `json:"state,omitempty"`

	// 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
	User []string `json:"user,omitempty"`

	// 使用该联系方式的部门id列表，只在type为2时有效
	Party []int `json:"party,omitempty"`

	// 是否临时会话模式，true表示使用临时会话模式，默认为false
	IsTemp bool `json:"is_temp,omitempty"`

	// 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天，最多为14天
	ExpiresIn int `json:"expires_in,omitempty"`

	// 临时会话有效期，以秒为单位。该参数仅在is_temp为true时有效，默认为添加好友后24小时，最多为14天
	ChatExpiresIn int `json:"chat_expires_in,omitempty"`

	// 可进行临时会话的客户unionid，该参数仅在is_temp为true时有效，如不指定则不进行限制
	Unionid string `json:"unionid,omitempty"`

	// 结束语，会话结束时自动发送给客户，可参考“结束语定义”，仅在is_temp为true时有效
	Conclusions Conclusions `json:"conclusions,omitempty"`
}

// ConclusionsText 结束语文本消息
type ConclusionsText struct {
	Content string `json:"content,omitempty"`
}

// ConclusionsImage 结束语图片消息
type ConclusionsImage struct {
	MediaId string `json:"media_id,omitempty"`
	PicUrl  string `json:"pic_url,omitempty"`
}

// ConclusionsLink 结束语图文消息
type ConclusionsLink struct {
	Title  string `json:"title,omitempty"`
	PicUrl string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Url    string `json:"url,omitempty"`
}

// ConclusionsMiniprogram 结束语小程序消息
type ConclusionsMiniprogram struct {
	Title      string `json:"title,omitempty"`        // 消息标题
	PicMediaID string `json:"pic_media_id,omitempty"` // 小程序消息封面的mediaid，封面图建议尺寸为520*416
	Appid      string `json:"appid,omitempty"`        // 小程序appid,必须是与当前小程序应用关联的小程序
	Page       string `json:"page,omitempty"`         // 点击消息卡片后的小程序页面，仅限本小程序内的页面
}

// Conclusions 结束语定义
type Conclusions struct {
	Text        ConclusionsText        `json:"text,omitempty"`
	Image       ConclusionsImage       `json:"image,omitempty"`
	Link        ConclusionsLink        `json:"link,omitempty"`
	Miniprogram ConclusionsMiniprogram `json:"miniprogram,omitempty"`
}

// ExternalAddContactWay 配置客户联系「联系我」方式
// 文档: https://developer.work.weixin.qq.com/document/path/92572#%E9%85%8D%E7%BD%AE%E5%AE%A2%E6%88%B7%E8%81%94%E7%B3%BB%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (a *Agent) ExternalAddContactWay(cw ContactWay) (map[string]string, error) {
	body, _ := json.Marshal(cw)

	var caller struct {
		baseCaller
		ConfigID string `json:"config_id"`
		QrCode   string `json:"qr_code"`
	}

	err := a.ExecuteWithToken("POST", "externalcontact/add_contact_way", nil, bytes.NewReader(body), &caller)

	if err != nil {
		return nil, err
	}

	resp := map[string]string{
		"config_id": caller.ConfigID,
		"qr_code":   caller.QrCode,
	}

	return resp, nil
}

// ExternalGetContactWay 获取企业已配置的「联系我」方式
// 文档: https://developer.work.weixin.qq.com/document/path/92572#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (a *Agent) ExternalGetContactWay(config_id string) (ContactWay, error) {
	param := map[string]string{
		"config_id": config_id,
	}
	body, _ := json.Marshal(param)

	var caller struct {
		baseCaller
		ContactWay ContactWay `json:"contact_way"`
	}

	err := a.ExecuteWithToken("POST", "externalcontact/get_contact_way", nil, bytes.NewReader(body), &caller)

	return caller.ContactWay, err
}

// ExternalListContactWay 获取企业已配置的「联系我」列表
// 文档: https://developer.work.weixin.qq.com/document/path/92572#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E5%88%97%E8%A1%A8
// @response contact_way 联系方式的配置id
// @response next_cursor 分页参数，用于查询下一个分页的数据，为空时表示没有更多的分页
func (a *Agent) ExternalListContactWay(start_time, end_time int64, limit int, cursor string) (contact_way []map[string]string, next_cursor string, err error) {
	param := map[string]interface{}{
		"start_time": start_time,
		"end_time":   end_time,
		"limit":      limit,
		"cursor":     cursor,
	}
	body, _ := json.Marshal(param)

	var caller struct {
		baseCaller
		ContactWay []map[string]string `json:"contact_way"`
		NextCursor string              `json:"next_cursor"`
	}

	err = a.ExecuteWithToken("POST", "externalcontact/list_contact_way", nil, bytes.NewReader(body), &caller)

	return caller.ContactWay, caller.NextCursor, err
}

// ExternalUpdateContactWay 更新企业已配置的「联系我」方式
// 文档: https://developer.work.weixin.qq.com/document/path/92572#%E6%9B%B4%E6%96%B0%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (a *Agent) ExternalUpdateContactWay(cw ContactWay) error {
	body, _ := json.Marshal(cw)

	var caller baseCaller

	return a.ExecuteWithToken("POST", "externalcontact/update_contact_way", nil, bytes.NewReader(body), &caller)
}

// ExternalDelContactWay 删除企业已配置的「联系我」方式
// 文档: https://developer.work.weixin.qq.com/document/path/92572#%E5%88%A0%E9%99%A4%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (a *Agent) ExternalDelContactWay(config_id string) error {
	param := map[string]string{
		"config_id": config_id,
	}
	body, _ := json.Marshal(param)

	var caller baseCaller

	return a.ExecuteWithToken("POST", "externalcontact/del_contact_way", nil, bytes.NewReader(body), &caller)
}

// ExternalCloseTempChat 结束临时会话
// 文档: https://developer.work.weixin.qq.com/document/path/92572#%E7%BB%93%E6%9D%9F%E4%B8%B4%E6%97%B6%E4%BC%9A%E8%AF%9D
func (a *Agent) ExternalCloseTempChat(userid, external_userid string) error {
	param := map[string]string{
		"userid":          userid,
		"external_userid": external_userid,
	}

	body, _ := json.Marshal(param)

	var caller baseCaller

	return a.ExecuteWithToken("POST", "externalcontact/close_temp_chat", nil, bytes.NewReader(body), &caller)
}
