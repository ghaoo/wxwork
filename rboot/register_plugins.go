package main

import "github.com/ghaoo/rboot"

func (pay *payload) registerPlugin() {
	rboot.RegisterPlugin(`wework_user`, pay.userPlugin())
	rboot.RegisterPlugin(`wework_media`, pay.mediaPlugin())
}
