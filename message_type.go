package workwx

// TextMessage 文本消息
type TextMessage struct {
	Content string `json:"content,omitempty"` // 消息内容，最长不超过2048个字节
}

// NewTextMessage 创建一条文本消息
func NewTextMessage(content string) *Message {
	return &Message{
		MsgType: MSG_TYPE_TEXT,
		Text:    &TextMessage{Content: content},
	}
}

// NewMarkdownMessage 创建一条 markdown 消息
func NewMarkdownMessage(content string) *Message {
	return &Message{
		MsgType:  MSG_TYPE_MARKDOWN,
		Markdown: &TextMessage{Content: content},
	}
}

// MediaMessage 素材消息
type MediaMessage struct {
	MediaId string `json:"media_id,omitempty"`
}

// NewMediaMessage 创建一条素材消息（image、voice、file）
func NewMediaMessage(mediaType, mediaId string) *Message {
	msg := &Message{MsgType: mediaType}
	mediaMsg := &MediaMessage{MediaId: mediaId}
	switch mediaType {
	case MSG_TYPE_IMAGE:
		msg.Image = mediaMsg
	case MSG_TYPE_VOICE:
		msg.Voice = mediaMsg
	case MSG_TYPE_FILE:
		msg.File = mediaMsg
	}

	return msg
}

// VideoMessage 视频消息
type VideoMessage struct {
	Title       string `json:"title,omitempty"`       // 视频标题
	Description string `json:"description,omitempty"` // 视频介绍
	MediaMessage
}

// NewVideoMessage 创建一条视频消息
func NewVideoMessage(title, desc, mediaId string) *Message {
	return &Message{
		MsgType: MSG_TYPE_VIDEO,
		Video: &VideoMessage{
			Title:        title,
			Description:  desc,
			MediaMessage: MediaMessage{MediaId: mediaId},
		},
	}
}

// TextCardMessage 文本卡片消息
type TextCardMessage struct {
	Title       string `json:"title,omitempty"`       // 消息标题
	Description string `json:"description,omitempty"` // 消息描述
	Url         string `json:"url,omitempty"`         // 消息跳转链接
	BtnTxt      string `json:"btntxt,omitempty"`      // 按钮文字, 默认为“详情”
}

// NewTextCardMessage 创建一条文本卡片消息
func NewTextCardMessage(title, desc, url, btntxt string) *Message {
	return &Message{
		MsgType: MSG_TYPE_TEXTCARD,
		TextCard: &TextCardMessage{
			Title:       title,
			Description: desc,
			Url:         url,
			BtnTxt:      btntxt,
		},
	}
}

// NewsMessage 图文消息
type NewsMessage struct {
	Articles []NewsArticle `json:"articles,omitempty"` // 图文消息内容，支持1-8条图文
}

// NewsArticle 图文消息内容
type NewsArticle struct {
	Title       string `json:"title,omitempty"`       // 图文标题
	Description string `json:"description,omitempty"` // 图文描述
	Url         string `json:"url,omitempty"`         // 跳转链接
	PicUrl      string `json:"picurl,omitempty"`      // 图片链接
}

// NewNewsMessage 创建一条图文消息，articles 最大容量为8
func NewNewsMessage(articles []NewsArticle) *Message {
	return &Message{
		MsgType: MSG_TYPE_NEWS,
		News:    &NewsMessage{Articles: articles},
	}
}

// MPNewsMessage 图文消息（mpnews）
type MPNewsMessage struct {
	Articles []MPNewsArticle `json:"articles,omitempty"` // 图文消息内容，支持1-8条图文
}

// MPNewsArticle 图文消息内容（mpnews）
type MPNewsArticle struct {
	Title            string `json:"title,omitempty"`              // 图文标题
	ThumbMediaId     string `json:"thumb_media_id,omitempty"`     // 缩略图素材ID
	Author           string `json:"author,omitempty"`             // 作者
	ContentSourceUrl string `json:"content_source_url,omitempty"` // 页面链接
	Content          string `json:"content,omitempty"`            // 消息内容
	Digest           string `json:"digest,omitempty"`             // 消息描述
}

// NewMPNewsMessage 创建一条图文消息（mpnews）
func NewMPNewsMessage(articles []MPNewsArticle) *Message {
	return &Message{
		MsgType: MSG_TYPE_MPNEWS,
		MPNews:  &MPNewsMessage{Articles: articles},
	}
}

// MiniprogramNoticeMessage 小程序消息
type MiniprogramNoticeMessage struct {
	Appid             string            `json:"appid,omitempty"`               // 小程序appid,必须是与当前小程序应用关联的小程序
	Page              string            `json:"page,omitempty"`                // 点击消息卡片后的小程序页面，仅限本小程序内的页面
	Title             string            `json:"title,omitempty"`               // 消息标题
	Description       string            `json:"description,omitempty"`         // 消息描述
	EmphasisFirstItem bool              `json:"emphasis_first_item,omitempty"` // 是否放大第一个content_item
	ContentItem       map[string]string `json:"content_item,omitempty"`        // 消息内容键值对，最多允许10个item
}

// NewMiniprogramNoticeMessage 创建一条小程序消息
func NewMiniprogramNoticeMessage(appid, page, title, desc string, efi bool, contentItem map[string]string) *Message {
	mini := &MiniprogramNoticeMessage{
		Appid:             appid,
		Page:              page,
		Title:             title,
		Description:       desc,
		EmphasisFirstItem: efi,
	}

	for key, value := range contentItem {
		mini.ContentItem["key"] = key
		mini.ContentItem["value"] = value
	}

	msg := &Message{
		MsgType:     MSG_TYPE_MINIPROGRAM_NOTICE,
		MiniProgram: mini,
	}

	return msg
}

// TaskCardMessage 任务卡片消息
type TaskCardMessage struct {
	Title       string        `json:"title,omitempty"`       // 消息标题
	Description string        `json:"description,omitempty"` // 消息描述
	Url         string        `json:"url,omitempty"`         // 跳转链接
	TaskId      string        `json:"task_id,omitempty"`     // 任务id
	Btn         []TaskCardBtn `json:"btn,omitempty"`         // 按钮列表，按钮个数为为1~2个
}

// TaskCardBtn 任务卡片按钮列表
type TaskCardBtn struct {
	Key         string `json:"key,omitempty"`          // 按钮key值
	Name        string `json:"name,omitempty"`         // 按钮名称
	ReplaceName string `json:"replace_name,omitempty"` // 点击按钮后显示的名称
	Color       string `json:"color,omitempty"`        // 按钮字体颜色
	IsBold      bool   `json:"is_bold,omitempty"`      // 按钮字体是否加粗
}

// NewTaskCardMessage 创建一条任务卡片消息
func NewTaskCardMessage(title, desc, url, taskId string, btn []TaskCardBtn) *Message {
	return &Message{
		MsgType: MSG_TYPE_TASKCARD,
		TaskCard: &TaskCardMessage{
			Title:       title,
			Description: desc,
			Url:         url,
			TaskId:      taskId,
			Btn:         btn,
		},
	}
}
