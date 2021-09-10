package main

import (
	"fmt"
	"github.com/ghaoo/rboot"
)

// mediaUpload 文件上传
func (pay *payload) mediaUpload(file string) []*rboot.Message {
	media, err := pay.client.MediaUpload(file)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("上传临时文件失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("上传临时文件成功，media_id: %s", media.MediaId))
}

// uploadImg 上传图片
func (pay *payload) uploadImg(file string) []*rboot.Message {
	picUrl, err := pay.client.UploadImg(file)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("上传图片失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("[点击查看图片](%s)", picUrl))
}

// mediaSetup 素材管理插件解析函数
func (pay *payload) mediaSetup(bot *rboot.Robot, in *rboot.Message) []*rboot.Message {
	rule := in.Header.Get("rule")
	args := in.Header["args"]

	switch rule {
	case `media_upload`:
		return pay.mediaUpload(args[1])
	case `media_img`:
		return pay.uploadImg(args[1])
	}

	return nil
}

// mediaPlugin 素材管理插件
func (pay *payload) mediaPlugin() rboot.Plugin {
	return rboot.Plugin{
		Action: pay.mediaSetup,
		Ruleset: map[string]string{
			`media_upload`: `^!media upload (.+)`,
			`media_img`:    `^!media img (.+)`,
		},
		Usage: map[string]string{
			"!media upload [文件]": "上传临时素材!",
			"!media img [文件]":    "上传图片!",
		},
		Description: `企业微信媒体管理SDK示例和测试`,
	}
}
