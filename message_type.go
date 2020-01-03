package workwx

// TextMessage 文本消息
type TextMessage struct {
	Content string `json:"content"` // 消息内容，最长不超过2048个字节
}

// NewTextMessage 实例化文本消息
func NewTextMessage(content string) *Message {
	return &Message{
		MsgType: "text",
		Text:    TextMessage{Content: content},
	}
}

// MediaMessage 多媒体消息
type MediaMessage struct {
	MediaId string `json:"media_id"`
}

// NewMediaMessage 创建一条素材消息（image、voice、file）
func NewMediaMessage(mediaType, mediaId string) *Message {
	msg := &Message{MsgType: mediaType}
	mediaMsg := MediaMessage{MediaId: mediaId}
	switch mediaType {
	case "image":
		msg.Image = mediaMsg
	case "voice":
		msg.Voice = mediaMsg
	case "file":
		msg.File = mediaMsg
	}

	return msg
}

// VideoMessage 定义了消息推送中的视频消息
type VideoMessage struct {
	Title       string `json:"title,omitempty"`       // 视频标题
	Description string `json:"description,omitempty"` // 视频介绍
	MediaMessage
}
