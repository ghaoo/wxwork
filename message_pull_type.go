// 接收消息--消息类型
package wxwork

//
type TextPullMessage struct {
	Content string `xml:"Content,omitempty"`
}
