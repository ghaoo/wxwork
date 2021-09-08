package main

import (
	"fmt"
	"github.com/ghaoo/rboot"
)

func (pay *payload) setDebug(d bool) []*rboot.Message {
	pay.client.SetDebug(d)

	var status string
	if d {
		status = "开启"
	} else {
		status = "关闭"
	}

	return rboot.NewMessages(fmt.Sprintf("设置成功，DEBUG现在为 `%s` 状态", status))
}
