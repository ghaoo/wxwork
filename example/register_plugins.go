package main

import "github.com/ghaoo/rboot"

func (pay *payload) registerPlugin() {
	rboot.RegisterPlugin(`wework_user`, pay.userPlugin())
}
