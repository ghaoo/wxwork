package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghaoo/rboot"
	"github.com/ghaoo/wxwork"
	"os"
	"strconv"
	"strings"
	"time"
)

// createUser 创建企业成员
func (pay *payload) createUser(name, mobile, department string) []*rboot.Message {
	userId := RandomCreateBytes(6)

	dID, err := strconv.Atoi(department)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("部门ID解析失败，错误信息: %v", err))
	}

	user := &wxwork.User{
		UserID:     string(userId),
		Name:       name,
		Mobile:     mobile,
		Department: []int{dID},
	}

	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	err = client.CreateUser(user)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("创建成员失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("创建成员成功，成员ID: %s, 请前往企业微信后台查看", userId))
}

// getUser 获取成员信息
func (pay *payload) getUser(uid string) []*rboot.Message {
	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	user, err := client.GetUser(uid)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("获取成员失败，错误信息: %v", err))
	}
	gender := "未知"
	switch user.Gender {
	case "1":
		gender = "男"
	case "2":
		gender = "女"
	default:
	}

	status := "未知"
	switch user.Status {
	case 1:
		status = "已激活"
	case 2:
		status = "已禁用"
	case 4:
		status = "未激活"
	case 5:
		status = "退出企业"
	default:
	}

	info := fmt.Sprintf("ID: `%s`\n 姓名: `%s`\n 电话: `%s`\n性别: `%s`\n状态: `%s`", user.UserID, user.Name, user.Mobile, gender, status)

	return rboot.NewMessages(fmt.Sprintf("成员基础信息: \n%s", info))
}

// updateUser 更新成员
// 示例更新成员名称，更新其他参数类似
func (pay *payload) updateUser(uid, name string) []*rboot.Message {
	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	// 组织更新参数
	user := &wxwork.User{
		UserID: uid,
		Name:   name,
	}
	err := client.UpdateUser(user)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("更新成员失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("更新成员成功，成员ID: %s, 请前往企业微信后台查看", uid))
}

// deleteUser 删除成员
func (pay *payload) deleteUser(uid string) []*rboot.Message {
	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	err := client.DeleteUser(uid)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("删除成员失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("删除成员成功，成员ID: %s, 请前往企业微信后台查看", uid))
}

// batchDeleteUsers 批量删除成员
func (pay *payload) batchDeleteUsers(uids string) []*rboot.Message {
	uid := strings.Split(uids, ",")

	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	err := client.BatchDeleteUsers(uid...)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("批量删除成员失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("批量删除成员成功，成员ID: %s, 请前往企业微信后台查看", uid))
}

// simpleListUser 获取部门成员
func (pay *payload) simpleListUser(depId string) []*rboot.Message {
	//client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	did, _ := strconv.Atoi(depId)
	users, err := pay.client.SimpleListUser(did)

	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("获取部门成员失败，错误信息: %v", err))
	}

	data, _ := json.Marshal(users)
	fname := "department-" + depId + ".simple.json"
	fpath := storeResult(fname, data)

	msg := rboot.NewMessage(fmt.Sprintf("获取部门成员成功，文件位置: %s", fpath))
	msg.Header.Set("msgtype", "file")
	msg.Header.Set("file", fpath)

	return []*rboot.Message{msg}

}

// simpleListUser 获取部门成员详情
func (pay *payload) listUser(depId string) []*rboot.Message {
	//client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	did, _ := strconv.Atoi(depId)
	users, err := pay.client.SimpleListUser(did)

	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("获取部门成员详情失败，错误信息: %v", err))
	}

	data, _ := json.Marshal(users)
	fname := "department-" + depId + ".info.json"
	fpath := storeResult(fname, data)

	msg := rboot.NewMessage(fmt.Sprintf("获取部门成员详情成功，文件位置: %s", fpath))
	msg.Header.Set("msgtype", "file")
	msg.Header.Set("file", fpath)

	return []*rboot.Message{msg}
}

// useridToOpenid userid转openid
func (pay *payload) useridToOpenid(userid string) []*rboot.Message {
	openid, err := pay.client.UserIDConvertToOpenID(userid)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("转换失败，错误信息: %v", err))
	}

	return rboot.NewMessages(openid)
}

// openidToUserid openid转userid
func (pay *payload) openidToUserid(openid string) []*rboot.Message {
	userid, err := pay.client.OpenIDConvertToUserID(openid)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("转换失败，错误信息: %v", err))
	}

	return rboot.NewMessages(userid)
}

// batchInvite 邀请成员使用企业微信
func (pay *payload) batchInvite(userid string) []*rboot.Message {
	invaliduser, _, _, err := pay.client.BatchInvite([]string{userid}, nil, nil)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("邀请失败，错误信息: %v", err))
	}

	if len(invaliduser) > 0 {
		return rboot.NewMessages(fmt.Sprintf("邀请成功，但存在无效用户: %s", strings.Join(invaliduser, ",")))
	}

	return rboot.NewMessages("邀请成员成功")
}

// getJoinQrCode 获取加入企业二维码
// qrcode尺寸类型，1: 171 x 171; 2: 399 x 399; 3: 741 x 741; 4: 2052 x 2052
func (pay *payload) getJoinQrCode() []*rboot.Message {
	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	qrcode, err := client.GetJoinQrCode("3")
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("获取二维码失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("[点击查看二维码](%s)", qrcode))
}

// getActiveStat 获取企业活跃成员数
func (pay *payload) getActiveStat(dateStr string) []*rboot.Message {
	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("请输入正确的时间: %v", err))
	}

	stat, err := client.GetActiveStat(date)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("获取企业活跃成员数失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("企业活跃成员: %d", stat))
}

func (pay *payload) userSetup(bot *rboot.Robot, in *rboot.Message) []*rboot.Message {
	rule := in.Header.Get("rule")
	args := in.Header["args"]

	switch rule {
	case `user_create`:
		return pay.createUser(args[1], args[2], args[3])
	case `user_get`:
		return pay.getUser(args[1])
	case `user_update`:
		return pay.updateUser(args[1], args[2])
	case `user_delete`:
		return pay.deleteUser(args[1])
	case `user_batchdel`:
		return pay.batchDeleteUsers(args[1])
	case `user_list`:
		return pay.simpleListUser(args[1])
	case `user_lists`:
		return pay.listUser(args[1])
	case `user_invite`:
		return pay.batchInvite(args[1])
	case `user_get_qrcode`:
		return pay.getJoinQrCode()
	case `user_active_stat`:
		return pay.getActiveStat(args[1])
	case `user_to_openid`:
		return pay.useridToOpenid(args[1])
	case `user_openid_to_user`:
		return pay.openidToUserid(args[1])
	default:
	}

	return nil
}

var userRuleset = map[string]string{
	`user_create`:         `^!user create (.+) (\d+) (\d+)`,
	`user_get`:            `^!user get ([\w@\-\.]+)`,
	`user_update`:         `^!user update ([\w@\-\.]+) (.+)`,
	`user_delete`:         `^!user delete ([\w@\-\.]+)`,
	`user_batchdel`:       `^!user batchdel ([\w@\-\.,]+)`,
	`user_list`:           `^!user list (\d+)`,
	`user_lists`:          `^!user lists (\d+)`,
	`user_invite`:         `^!user invite ([\w@\-\.]+)`,
	`user_get_qrcode`:     `^!user qrcode`,
	`user_active_stat`:    `^!user stat (\d{4}\-\d{2}-\d{2})`,
	`user_to_openid`:      `^!userid to openid ([\w@\-\.]+)`,
	`user_openid_to_user`: `^!openid to userid ([\w@\-\.]+)`,
}

var userUsage = map[string]string{
	"!user create [姓名] [电话] [部门ID]": "新增成员，注意权限!",
	"!user get [成员ID]":              "读取成员",
	"!user update [成员ID] [姓名]":      "更新成员",
	"!user delete [成员ID]":           "删除成员",
	"!user batchdel [成员ID,成员ID...]": "批量删除成员，用`,`隔开",
	"!user list [部门ID]":             "获取部门成员，注意权限",
	"!user lists [部门ID]":            "获取部门成员详情，注意权限",
	"!user invite [成员ID]":           "邀请部门成员使用企业微信",
	"!user qrcode":                  "获取加入企业二维码，注意权限",
	"!user stat [起始时间2021-01-02]":   "获取企业活跃成员数，最长支持获取30天前数据",
	"!userid to openid [成员ID]":      "成员ID转换为Openid",
	"!openid to userid [Openid]":    "Openid转换为成员ID",
}

func (pay *payload) userPlugin() rboot.Plugin {
	return rboot.Plugin{
		Action:      pay.userSetup,
		Ruleset:     userRuleset,
		Usage:       userUsage,
		Description: `企业微信成员SDK示例和测试`,
	}
}
