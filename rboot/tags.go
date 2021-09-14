package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghaoo/rboot"
	"strconv"
	"strings"
)

// tagCreate 创建标签
func (pay *payload) tagCreate(name string) []*rboot.Message {
	tagId, err := pay.client.CreateTag(name)

	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("创建标签失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("创建标签成功，标签ID: %d, 请前往企业微信后台查看", tagId))
}

// tagUpdate 更新标签名称
func (pay *payload) tagUpdate(id, name string) []*rboot.Message {
	tagId, _ := strconv.Atoi(id)
	err := pay.client.UpdateTag(name, tagId)

	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("更新标签名称失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("更新标签名称成功，标签名称已更名为: %s, 请前往企业微信后台查看", name))
}

// tagDelete 删除标签，必须为标签创建者才可删除
func (pay *payload) tagDelete(id string) []*rboot.Message {
	tagId, _ := strconv.Atoi(id)
	err := pay.client.DeleteTag(tagId)

	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("删除标签失败，错误信息: %v", err))
	}

	return rboot.NewMessages("删除标签成功，请前往企业微信后台查看")
}

// tagGet 获取标签成员
func (pay *payload) tagGet(id string) []*rboot.Message {
	tagId, _ := strconv.Atoi(id)

	tag, err := pay.client.GetTag(tagId)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("获取标签成员失败，错误信息: %v", err))
	}

	data, _ := json.Marshal(tag)
	fname := "tag-" + id + ".json"
	fpath := storeResult(fname, data)

	msg := rboot.NewMessage(fmt.Sprintf("获取部门成员详情成功，文件位置: %s", fpath))
	msg.Header.Set("msgtype", "file")
	msg.Header.Set("file", fpath)

	return []*rboot.Message{msg}
}

// tagAddUsers 增加标签成员
func (pay *payload) tagAddUsers(id, uid string) []*rboot.Message {
	tagId, _ := strconv.Atoi(id)
	users, _, err := pay.client.AddTagUsers(tagId, []string{uid}, nil)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("新增标签成员失败，错误信息: %v", err))
	}

	if len(users) > 0 {
		return rboot.NewMessages(fmt.Sprintf("非法的成员帐号: %s", users))
	}

	return rboot.NewMessages("新增标签成员成功，请前往企业微信后台查看")
}

// tagDelUsers 删除标签成员
func (pay *payload) tagDelUsers(id, uid string) []*rboot.Message {
	tagId, _ := strconv.Atoi(id)
	users, _, err := pay.client.DelTagUsers(tagId, []string{uid}, nil)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("删除标签成员失败，错误信息: %v", err))
	}

	if len(users) > 0 {
		return rboot.NewMessages(fmt.Sprintf("非法的成员帐号: %s", users))
	}

	return rboot.NewMessages("删除标签成员成功，请前往企业微信后台查看")
}

// tagList 获取标签列表
func (pay *payload) tagList() []*rboot.Message {
	tags, err := pay.client.ListTags()
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("获取标签列表失败，错误信息: %v", err))
	}

	msg := ""
	for _, tag := range tags {
		msg += "`" + tag.TagName + "`,"
	}

	msg = strings.TrimSuffix(msg, ",")

	return rboot.NewMessages(msg)
}

func (pay *payload) tagSetup(bot *rboot.Robot, in *rboot.Message) []*rboot.Message {
	rule := in.Header.Get("rule")
	args := in.Header["args"]

	switch rule {
	case `tag_create`:
		return pay.tagCreate(args[1])
	case `tag_update`:
		return pay.tagUpdate(args[1], args[2])
	case `tag_delete`:
		return pay.tagDelete(args[1])
	case `tag_get`:
		return pay.tagGet(args[1])
	case `tag_add_user`:
		return pay.tagAddUsers(args[1], args[2])
	case `tag_del_user`:
		return pay.tagDelUsers(args[1], args[2])
	case `tag_list`:
		return pay.tagList()
	}

	return nil
}

var tagRuleset = map[string]string{
	`tag_create`:   `^!tag create (.+)`,
	`tag_update`:   `^!tag update (\d+) (.+)`,
	`tag_delete`:   `^!tag delete (\d+)`,
	`tag_get`:      `^!tag get (\d+)`,
	`tag_add_user`: `^!tag adduser (\d+) ([\w@\-\.]+)`,
	`tag_del_user`: `^!tag deluser (\d+) ([\w@\-\.]+)`,
	`tag_list`:     `^!tag list`,
}

var tagUsage = map[string]string{
	"!tag create [标签名]":          "创建标签",
	"!tag update [标签ID] [标签名]":   "修改标签名称",
	"!tag delete [标签ID]":         "删除标签",
	"!tag get [标签ID]":            "获取标签成员",
	"!tag adduser [标签ID] [成员ID]": "新增标签成员",
	"!tag deluser [标签ID] [成员ID]": "删除标签成员",
	"!tag list":                  "获取标签列表",
}

func (pay *payload) tagPlugin() rboot.Plugin {
	return rboot.Plugin{
		Action:      pay.tagSetup,
		Ruleset:     tagRuleset,
		Usage:       tagUsage,
		Description: `企业微信标签SDK示例和测试`,
	}
}
