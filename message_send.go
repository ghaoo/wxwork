package workwx

// 文本消息
type TextMessage struct {
	Message
	Content string `json:"content"` // 消息内容，最长不超过2048个字节
}
