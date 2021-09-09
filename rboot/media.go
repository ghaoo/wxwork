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
