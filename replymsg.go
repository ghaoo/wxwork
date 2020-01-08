package wxwork

// ReplyMessage 被动回复消息格式
type ReplyMessage struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string

	// 文本消息
	Content string

	// 多媒体消息
	MediaId     string
	Title       string
	Description string

	ArticleCount int
	Url          string
	PicUrl       string
}
