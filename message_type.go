package workwx

// 文本消息
type TextMessage struct {
	Content string `json:"content"` // 消息内容，最长不超过2048个字节
}

func NewTextMessage(content string) *Message {
	return &Message{
		MsgType: "text",
		Text:    TextMessage{Content: content},
	}
}
