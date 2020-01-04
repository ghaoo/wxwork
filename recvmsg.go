// 接收消息
package wxwork

import (
	"fmt"
	"net/http"
)

// RecvMsg 消息接收基础结构
type RecvMsg struct {
	MsgType      string `xml:"MsgType"`      // 消息类型
	ToUsername   string `xml:"ToUserName"`   // 企业微信CorpID
	FromUsername string `xml:"FromUserName"` // 成员UserID
	CreateTime   int64  `xml:"CreateTime"`   // 消息创建时间（整型）
	AgentID      int    `xml:"AgentId"`      // 企业应用的id，整型
}

func (a *Agent) ParseRecvMsg(signature, timestamp, nonce string, body []byte) (interface{}, error) {

	return nil, nil

}

func (a *Agent) CallbackVerify(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
}
