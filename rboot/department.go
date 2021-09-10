package main

import (
	"encoding/json"
	"fmt"
	"github.com/ghaoo/rboot"
	"github.com/ghaoo/wxwork"
	"os"
	"strconv"
)

// departmentCreate 创建部门
// pid 父部门id，32位整型
func (pay *payload) departmentCreate(name, parentID string) []*rboot.Message {
	pid, _ := strconv.Atoi(parentID)
	dept := wxwork.Department{
		Name:     name,
		ParentID: pid,
	}

	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	deptID, err := client.CreateDepartment(dept)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("创建部门失败，错误信息: %v", err))
	}

	return rboot.NewMessages(fmt.Sprintf("创建部门成功，部门ID: %d", deptID))
}

// departmentUpdate 更新部门信息
func (pay *payload) departmentUpdate(deptID, name string) []*rboot.Message {
	did, _ := strconv.Atoi(deptID)
	dept := wxwork.Department{
		ID:   did,
		Name: name,
	}

	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))

	err := client.UpdateDepartment(dept)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("更新部门失败，错误信息: %v", err))
	}

	return rboot.NewMessages("更新部门成功")
}

// departmentDelete 删除部门
func (pay *payload) departmentDelete(deptID string) []*rboot.Message {
	id, _ := strconv.Atoi(deptID)

	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))
	err := client.DeleteDepartment(id)

	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("删除部门失败，错误信息: %v", err))
	}

	return rboot.NewMessages("删除部门成功")
}

// departmentList 获取部门列表
func (pay *payload) departmentList(id string) []*rboot.Message {
	var did = make([]int, 0)
	i, _ := strconv.Atoi(id)
	if i > 0 {
		did = append(did, i)
	}

	client := pay.client.WithSecret(os.Getenv("WORKWX_CONTACT_SECRET"))
	depts, err := client.ListDepartment(did...)
	if err != nil {
		return rboot.NewMessages(fmt.Sprintf("获取部门列表失败，错误信息: %v", err))
	}

	data, _ := json.Marshal(depts)
	fname := "departments.json"
	fpath := storeResult(fname, data)

	msg := rboot.NewMessage(fmt.Sprintf("获取部门列表成功，文件位置: %s", fpath))
	msg.Header.Set("msgtype", "file")
	msg.Header.Set("file", fpath)

	return []*rboot.Message{msg}
}

// departmentSetup 部门管理插件解析函数
func (pay *payload) departmentSetup(bot *rboot.Robot, in *rboot.Message) []*rboot.Message {
	rule := in.Header.Get("rule")
	args := in.Header["args"]

	switch rule {
	case `dept_create`:
		return pay.departmentCreate(args[1], args[2])
	case `dept_update`:
		return pay.departmentUpdate(args[1], args[2])
	case `dept_delete`:
		return pay.departmentDelete(args[1])
	case `dept_list`:
		return pay.departmentList(args[1])
	}

	return nil
}

// departmentPlugin 部门管理插件
func (pay *payload) departmentPlugin() rboot.Plugin {
	return rboot.Plugin{
		Action: pay.departmentSetup,
		Ruleset: map[string]string{
			`dept_create`: `^!dept create (.+) (\d+)`,
			`dept_update`: `^!dept update (\d+) (.+)`,
			`dept_delete`: `^!dept delete (\d+)`,
			`dept_list`:   `^!dept list[ ]?(\d*)`,
		},
		Usage: map[string]string{
			"!dept create [名称] [父ID]":  "创建部门",
			"!dept update [部门ID] [名称]": "修改部门",
			"!dept delete [部门ID]":      "删除部门",
			"!dept list [部门ID]":        "部门列表，为空不输入默认全量组织架构",
		},
		Description: `企业微信部门管理SDK示例和测试`,
	}
}
