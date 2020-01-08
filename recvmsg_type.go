package wxwork

/******************************Event 事件格式*********************************/

const (
	EVENT_TYPE_SUBSCRIBE            string = "subscribe"            // 关注
	EVENT_TYPE_UNSUBSCRIBE                 = "unsubscribe"          // 取消关注
	EVENT_TYPE_ENTER_AGENT                 = "enter_agent"          // 进入应用
	EVENT_TYPE_LOCATION                    = "LOCATION"             // 上报地理位置
	EVENT_TYPE_BATCH_JOB_RESULT            = "batch_job_result"     // 异步任务完成事件推送
	EVENT_TYPE_CHANGE_CONTACT              = "change_contact"       // 通讯录变更事件
	EVENT_TYPE_MENU_CLICK                  = "click"                // 点击菜单拉取消息的事件推送
	EVENT_TYPE_MENU_VIEW                   = "view"                 // 点击菜单跳转链接的事件推送
	EVENT_TYPE_SCANCODE_PUSH               = "scancode_push"        // 扫码推事件的事件推送
	EVENT_TYPE_SCANCODE_WAITMSG            = "scancode_waitmsg"     // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EVENT_TYPE_PIC_SYSPHOTO                = "pic_sysphoto"         // 弹出系统拍照发图的事件推送
	EVENT_TYPE_PIC_PHOTO_OR_ALBUM          = "pic_photo_or_album"   // 弹出拍照或者相册发图的事件推送
	EVENT_TYPE_PIC_WEIXIN                  = "pic_weixin"           // 弹出微信相册发图器的事件推送
	EVENT_TYPE_LOCATION_SELECT             = "location_select"      // 弹出地理位置选择器的事件推送
	EVENT_TYPE_OPEN_APPROVAL_CHANGE        = "open_approval_change" // 审批状态通知事件
	EVENT_TYPE_TASKCARD_CLICK              = "taskcard_click"       // 任务卡片事件推送
)

// RecvEvent 事件基础结构
// - 成员关注及取消关注事件、进入应用、菜单事件直接使用
type RecvEvent struct {
	Event    string `xml:"Event"`    // 事件类型
	EventKey string `xml:"EventKey"` // 事件KEY值

	// 上报地理位置事件
	Latitude  string `xml:"Latitude"`  // 地理位置纬度
	Longitude string `xml:"Longitude"` // 地理位置经度
	Precision string `xml:"Precision"` // 地理位置精度

	// 异步任务完成事件推送
	JobID   string `xml:"JobId"`   // 异步任务id
	JobType string `xml:"JobType"` // 操作类型
	ErrCode int    `xml:"ErrCode"`
	ErrMsg  string `xml:"ErrMsg"`

	// 通讯录变更事件
	ChangeType string `xml:"ChangeType"`
}
