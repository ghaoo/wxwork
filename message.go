package workwx

import (
	"bytes"
	"encoding/json"
	"strings"
)

const (
	MSG_TYPE_TEXT               = "text"               // 文本消息
	MSG_TYPE_IMAGE              = "image"              // 图片消息
	MSG_TYPE_VOICE              = "voice"              // 语音消息
	MSG_TYPE_VIDEO              = "video"              // 视频消息
	MSG_TYPE_FILE               = "file"               // 文件消息
	MSG_TYPE_TEXTCARD           = "textcard"           // 文本卡片消息
	MSG_TYPE_NEWS               = "news"               // 图文消息
	MSG_TYPE_MPNEWS             = "mpnews"             // 图文消息（mpnews）
	MSG_TYPE_MARKDOWN           = "markdown"           // markdown消息
	MSG_TYPE_MINIPROGRAM_NOTICE = "miniprogram_notice" // 小程序通知消息
	MSG_TYPE_TASKCARD           = "taskcard"           // 任务卡片消息
)

/**
 * 消息推送
 * - 文档地址: https://work.weixin.qq.com/api/doc/90000/90135/90235
 */
type Message struct {
	// 接收成员ID列表，“|”分隔,max:1000,全部成员:"@all"
	ToUser string `json:"touser"`

	// 接收部门ID列表，“|”分隔,max:100
	ToParty string `json:"toparty,omitempty"`

	// 接收标签ID列表，“|”分隔,max:100
	ToTag string `json:"totag,omitempty"`

	// 消息类型
	MsgType string `json:"msgtype"`

	// 企业应用ID
	AgentID int `json:"agentid"`

	// 是否是保密消息，0表示否，1表示是，默认0
	Safe bool `json:"safe,omitempty"`

	// 是否开启id转译，0表示否，1表示是，默认0
	EnableIdTrans bool `json:"enable_id_trans,omitempty"`

	// 是否开启重复消息检查，0表示否，1表示是，默认0
	EnableDuplicateCheck bool `json:"enable_duplicate_check,omitempty"`

	// 是否重复消息检查的时间间隔，默认1800s，最大不超过4小时
	DuplicateCheckInterval int64 `json:"duplicate_check_interval,omitempty"`

	Text        TextMessage              `json:"text,omitempty"`               // 文本消息
	Markdown    TextMessage              `json:"markdown,omitempty"`           // markdown 消息
	Image       MediaMessage             `json:"image,omitempty"`              // 图片消息
	Voice       MediaMessage             `json:"voice,omitempty"`              // 语音消息
	File        MediaMessage             `json:"file,omitempty"`               // 文件消息
	Video       VideoMessage             `json:"video,omitempty"`              // 视频消息
	TextCard    TextCardMessage          `json:"textcard,omitempty"`           // 文本卡片消息
	News        NewsMessage              `json:"news,omitempty"`               // 图文消息
	MPNews      MPNewsMessage            `json:"mpnews,omitempty"`             // 图文消息(mpnews)
	MiniProgram MiniprogramNoticeMessage `json:"miniprogram_notice,omitempty"` // 小程序消息
	TaskCard    TaskCardMessage          `json:"taskcard,omitempty"`           // 任务卡片消息
}

// SetUser 设置接收成员
func (msg *Message) SetUser(user ...string) {
	msg.ToUser = strings.Join(user, "|")
}

// SetParty 设置接收部门
func (msg *Message) SetParty(party ...string) {
	msg.ToParty = strings.Join(party, "|")
}

// SetTag 设置接收标签
func (msg *Message) SetTag(tag ...string) {
	msg.ToTag = strings.Join(tag, "|")
}

// RespMessage 定义了消息会话响应
type RespMessage struct {
	// 如果全部接收人无权限或不存在，则本次调用返回失败，errcode为81013。
	baseCaller
	InvalidUser  string `json:"invaliduser"`
	InvalidParty string `json:"invalidparty"`
	InvalidTag   string `json:"invalidtag"`
}

// SendMessage 用于消息推送-发送应用消息，返回接收失败用户、组织、标签列表
func (c *Client) SendMessage(msg *Message) (map[string][]string, error) {
	msg.AgentID = c.AgentID

	body, _ := json.Marshal(msg)

	var resp RespMessage

	var invalid = make(map[string][]string, 3)
	err := c.ExecuteWithToken("POST", "message/send", bytes.NewReader(body), &resp)
	if err != nil {
		return nil, err
	}
	invalid["user"] = strings.Split(resp.InvalidUser, "|")
	invalid["party"] = strings.Split(resp.InvalidParty, "|")
	invalid["tag"] = strings.Split(resp.InvalidTag, "|")

	return invalid, nil
}

// UpdateTaskcard 更新任务卡片消息状态,返回接收失败用户列表
func (c *Client) UpdateTaskcard(taskId, clickedKey string, userids []string) ([]string, error) {
	request := map[string]interface{}{
		"userids":     userids,
		"agentid":     c.AgentID,
		"task_id":     taskId,
		"clicked_key": clickedKey,
	}

	body, _ := json.Marshal(request)

	var resp struct {
		baseCaller
		Invaliduser []string `json:"invaliduser"`
	}

	err := c.ExecuteWithToken("POST", "message/update_taskcard", bytes.NewReader(body), &resp)

	return resp.Invaliduser, err

}
