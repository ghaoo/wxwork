// 接收消息
package wxwork

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
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

// ParseRecvMsg 解析接收到的消息
func (a *Agent) ParseRecvMsg(signature, timestamp, nonce string, data []byte) (interface{}, error) {
	msg, cryptErr := a.callback.crypt.DecryptMsg(signature, timestamp, nonce, data)
	if nil != cryptErr {
		return nil, fmt.Errorf("DecryptMsg fail: %v", cryptErr)
	}

	var probe struct {
		MsgType string `xml:"MsgType"`
		Event   string `xml:"Event"`
	}

	if err := xml.Unmarshal(msg, probe); err != nil {
		return nil, err
	}

	var msgData interface{}
	switch probe.MsgType {
	case MSG_TYPE_TEXT:
		msgData = &RecvTextMessage{}
	case MSG_TYPE_IMAGE:
		msgData = &RecvImageMessage{}
	case MSG_TYPE_VIDEO:
		msgData = &RecvVideoMessage{}
	case MSG_TYPE_VOICE:
		msgData = &RecvVideoMessage{}
	case MSG_TYPE_LOCATION:
		msgData = &RecvLocationMessage{}
	case MSG_TYPE_LINK:
		msgData = &RecvLinkMessage{}
	case MSG_TYPE_EVENT:
		switch probe.Event {
		case EVENT_TYPE_SUBSCRIBE:
			fmt.Println(EVENT_TYPE_SUBSCRIBE)

		case EVENT_TYPE_UNSUBSCRIBE:
			fmt.Println(EVENT_TYPE_UNSUBSCRIBE)

		case EVENT_TYPE_ENTER_AGENT:
			fmt.Println(EVENT_TYPE_ENTER_AGENT)

		case EVENT_TYPE_LOCATION:
			fmt.Println(EVENT_TYPE_LOCATION)

		case EVENT_TYPE_BATCH_JOB_RESULT:
			fmt.Println(EVENT_TYPE_BATCH_JOB_RESULT)

		case EVENT_TYPE_CHANGE_CONTACT:
			fmt.Println(EVENT_TYPE_CHANGE_CONTACT)

		case EVENT_TYPE_MENU_CLICK:
			fmt.Println(EVENT_TYPE_MENU_CLICK)

		case EVENT_TYPE_MENU_VIEW:
			fmt.Println(EVENT_TYPE_MENU_VIEW)

		case EVENT_TYPE_SCANCODE_PUSH:
			fmt.Println(EVENT_TYPE_SCANCODE_PUSH)

		case EVENT_TYPE_SCANCODE_WAITMSG:
			fmt.Println(EVENT_TYPE_SCANCODE_WAITMSG)

		case EVENT_TYPE_PIC_PHOTO_OR_ALBUM:
			fmt.Println(EVENT_TYPE_PIC_PHOTO_OR_ALBUM)

		case EVENT_TYPE_PIC_WEIXIN:
			fmt.Println(EVENT_TYPE_PIC_WEIXIN)

		case EVENT_TYPE_LOCATION_SELECT:
			fmt.Println(EVENT_TYPE_LOCATION_SELECT)

		case EVENT_TYPE_OPEN_APPROVAL_CHANGE:
			fmt.Println(EVENT_TYPE_OPEN_APPROVAL_CHANGE)

		case EVENT_TYPE_TASKCARD_CLICK:
			fmt.Println(EVENT_TYPE_TASKCARD_CLICK)
		default:
			return nil, fmt.Errorf("unknown event type: %s", probe.Event)
		}
	default:
		return nil, fmt.Errorf("unknown message type: %s", probe.MsgType)
	}

	if err := xml.Unmarshal(data, msgData); err != nil {
		return nil, err
	}

	return msgData, nil
}

func (a *Agent) ParseRecvHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	signature := query.Get("msg_signature")
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("parse recive message err:", err)
	}

}

// CallbackVerify 回调配置验证URL有效性
func (a *Agent) CallbackVerify(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	verifySignature := query.Get("msg_signature")
	verifyTimestamp := query.Get("timestamp")
	verifyNonce := query.Get("nonce")
	verifyEchoStr := query.Get("echostr")

	echoStr, cryptErr := a.callback.crypt.VerifyURL(verifySignature, verifyTimestamp, verifyNonce, verifyEchoStr)

	if nil != cryptErr {
		log.Println("verifyUrl fail", cryptErr)
	}

	w.Write(echoStr)
}
