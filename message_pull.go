// 接收消息
package wxwork

type MessagePull struct {
	Msgid        uint64 `xml:"MsgId"`        // 消息id，64位整型
	Agentid      int    `xml:"AgentId"`      // 企业应用的id，整型
	MsgType      string `xml:"MsgType"`      // 消息类型
	ToUsername   string `xml:"ToUserName"`   // 企业微信CorpID
	FromUsername string `xml:"FromUserName"` // 成员UserID
	CreateTime   int64  `xml:"CreateTime"`   // 消息创建时间（整型）
}
