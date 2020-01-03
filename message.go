package workwx

import "strings"

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
}

func (msg *Message) SetUsers(user ...string) {
	msg.ToUser = strings.Join(user, "|")
}

type Msg interface {
	Send() error
}
