package main

import "github.com/ghaoo/rboot"

func (pay *payload) registerPlugin() {
	rboot.RegisterPlugin(`wework_user`, pay.userPlugin())
	rboot.RegisterPlugin(`wework_media`, pay.mediaPlugin())
	rboot.RegisterPlugin(`wework_department`, pay.departmentPlugin())
	rboot.RegisterPlugin(`wework_tag`, pay.tagPlugin())
}
