package wxwork

/******************************Message 消息格式*********************************/

// RecvTextMessage 文本消息
type RecvTextMessage struct {
	RecvMsg
	MsgID   int64  `xml:"MsgId"`   // 消息id，64位整型
	Content string `xml:"Content"` // 消息内容
}

// RecvImageMessage 图片消息
type RecvImageMessage struct {
	RecvMsg
	MsgID   int64  `xml:"MsgId"`   // 消息id，64位整型
	PicURL  string `xml:"PicUrl"`  // 图片链接
	MediaID string `xml:"MediaId"` // 图片媒体文件id
}

// RecvVoiceMessage 语音消息
type RecvVoiceMessage struct {
	RecvMsg
	MsgID   int64  `xml:"MsgId"`   // 消息id，64位整型
	MediaID string `xml:"MediaId"` // 语音媒体文件id
	Format  string `xml:"Format"`  // 语音格式
}

// RecvVideoMessage 视频消息
type RecvVideoMessage struct {
	RecvMsg
	MsgID        int64  `xml:"MsgId"`        // 消息id，64位整型
	MediaID      string `xml:"MediaId"`      // 视频媒体文件id
	ThumbMediaID string `xml:"ThumbMediaId"` // 缩略图的媒体id
}

// RecvLocationMessage 地理位置消息
type RecvLocationMessage struct {
	RecvMsg
	MsgID     int64   `xml:"MsgId"`      // 消息id，64位整型
	LocationX float64 `xml:"Location_X"` // 地理位置纬度
	LocationY float64 `xml:"Location_Y"` // 地理位置经度
	Scale     int     `xml:"Scale"`      // 地图缩放大小
	Label     string  `xml:"Label"`      // 地理位置信息
}

// RecvLinkMessage 链接消息
type RecvLinkMessage struct {
	RecvMsg
	MsgID       int64  `xml:"MsgId"`       // 消息id，64位整型
	Title       string `xml:"Title"`       // 标题
	Description string `xml:"Description"` // 链接跳转的url
	Url         string `xml:"Url"`         // 链接跳转的url
	PicUrl      string `xml:"PicUrl"`      // 封面缩略图的url
}

/******************************Event 事件格式*********************************/

const (
	EVENT_TYPE_SUBSCRIBE            = "subscribe"            // 关注
	EVENT_TYPE_UNSUBSCRIBE          = "unsubscribe"          // 取消关注
	EVENT_TYPE_ENTER_AGENT          = "enter_agent"          // 进入应用
	EVENT_TYPE_LOCATION             = "LOCATION"             // 上报地理位置
	EVENT_TYPE_BATCH_JOB_RESULT     = "batch_job_result"     // 异步任务完成事件推送
	EVENT_TYPE_CHANGE_CONTACT       = "change_contact"       // 通讯录变更事件
	EVENT_TYPE_MENU_CLICK           = "click"                // 点击菜单拉取消息的事件推送
	EVENT_TYPE_MENU_VIEW            = "view"                 // 点击菜单跳转链接的事件推送
	EVENT_TYPE_SCANCODE_PUSH        = "scancode_push"        // 扫码推事件的事件推送
	EVENT_TYPE_SCANCODE_WAITMSG     = "scancode_waitmsg"     // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EVENT_TYPE_PIC_SYSPHOTO         = "pic_sysphoto"         // 弹出系统拍照发图的事件推送
	EVENT_TYPE_PIC_PHOTO_OR_ALBUM   = "pic_photo_or_album"   // 弹出拍照或者相册发图的事件推送
	EVENT_TYPE_PIC_WEIXIN           = "pic_weixin"           // 弹出微信相册发图器的事件推送
	EVENT_TYPE_LOCATION_SELECT      = "location_select"      // 弹出地理位置选择器的事件推送
	EVENT_TYPE_OPEN_APPROVAL_CHANGE = "open_approval_change" // 审批状态通知事件
	EVENT_TYPE_TASKCARD_CLICK       = "taskcard_click"       // 任务卡片事件推送
)

// RecvBaseEvent 事件基础结构
// - 成员关注及取消关注事件、进入应用、菜单事件直接使用
type RecvEvent struct {
	RecvMsg
	Event    string `xml:"Event"`    // 事件类型
	EventKey string `xml:"EventKey"` // 事件KEY值
}

// RecvLocationEvent 上报地理位置事件
type RecvLocationEvent struct {
	RecvEvent
	Latitude  string `xml:"Latitude"`  // 地理位置纬度
	Longitude string `xml:"Longitude"` // 地理位置经度
	Precision string `xml:"Precision"` // 地理位置精度
}

// RecvBatchJobResultEvent 异步任务完成事件推送
type RecvBatchJobResultEvent struct {
	RecvEvent
	JobID   string `xml:"JobId"`   // 异步任务id
	JobType string `xml:"JobType"` // 操作类型
	ErrCode int    `xml:"ErrCode"`
	ErrMsg  string `xml:"ErrMsg"`
}

// RecvChangeContact 通讯录变更事件
type RecvChangeContact struct {
	RecvEvent
	ChangeType string `xml:"ChangeType"`
}
