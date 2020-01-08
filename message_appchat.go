package wxwork

import (
	"bytes"
	"encoding/json"
	"net/url"
)

// AppChat 群聊会话
type AppChat struct {
	ChatID   string   `json:"chatid"`   // 群聊的唯一标志
	Name     string   `json:"name"`     // 群聊名称
	Owner    string   `json:"owner"`    // 群主ID
	UserList []string `json:"userlist"` // 群成员id列表
}

// AppChatMessage 自建应用群聊消息
type AppChatMessage struct {
	// 群聊会话ID
	ChatID string `json:"chatid"`

	// 消息类型
	MsgType string `json:"msgtype,omitempty"`

	// 是否是保密消息，0表示否，1表示是，默认0
	Safe int8 `json:"safe,omitempty"`

	Text     *TextMessage     `json:"text,omitempty"`     // 文本消息
	Markdown *TextMessage     `json:"markdown,omitempty"` // markdown 消息
	Image    *MediaMessage    `json:"image,omitempty"`    // 图片消息
	Voice    *MediaMessage    `json:"voice,omitempty"`    // 语音消息
	File     *MediaMessage    `json:"file,omitempty"`     // 文件消息
	Video    *VideoMessage    `json:"video,omitempty"`    // 视频消息
	TextCard *TextCardMessage `json:"textcard,omitempty"` // 文本卡片消息
	News     *NewsMessage     `json:"news,omitempty"`     // 图文消息
	MPNews   *MPNewsMessage   `json:"mpnews,omitempty"`   // 图文消息(mpnews)
	TaskCard *TaskCardMessage `json:"taskcard,omitempty"` // 任务卡片消息
}

// CreateAppChat 创建群聊会话
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90245
func (a *Agent) CreateAppChat(data map[string]interface{}) (string, error) {
	body, _ := json.Marshal(data)

	var caller struct {
		baseCaller
		ChatID string `json:"chatid"`
	}
	err := a.ExecuteWithToken("POST", "appchat/create", nil, bytes.NewReader(body), &caller)
	if err != nil {
		return "", err
	}

	return caller.ChatID, nil
}

// UpdateAppChat 修改群聊会话
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90246
func (a *Agent) UpdateAppChat(data map[string]interface{}) error {
	body, _ := json.Marshal(data)

	var caller baseCaller
	return a.ExecuteWithToken("POST", "appchat/update", nil, bytes.NewReader(body), &caller)
}

// GetAppChat 获取群聊会话
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90247
func (a *Agent) GetAppChat(chatid string) (*AppChat, error) {
	var caller struct {
		baseCaller
		ChatInfo *AppChat
	}

	query := url.Values{}
	query.Set("chatid", chatid)

	err := a.ExecuteWithToken("GET", "appchat/get", query, nil, &caller)

	return caller.ChatInfo, err
}

// AppChatSendMessage 应用推送消息
// 文档: https://work.weixin.qq.com/api/doc/90000/90135/90248
func (a *Agent) AppChatSendMessage(msg *AppChatMessage) error {
	body, _ := json.Marshal(msg)

	var caller baseCaller

	return a.ExecuteWithToken("POST", "appchat/send", nil, bytes.NewReader(body), &caller)

}
